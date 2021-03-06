// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package handlerfactory allows creating HTTP handlers for services.
package handlerfactory

import (
	"net/http"
	"regexp"

	"github.com/golang/protobuf/proto" /* copybara-comment */
	"github.com/gorilla/mux" /* copybara-comment */
	"google.golang.org/grpc/codes" /* copybara-comment */
	"google.golang.org/grpc/status" /* copybara-comment */
	"github.com/GoogleCloudPlatform/healthcare-federated-access-services/lib/httputils" /* copybara-comment: httputils */
	"github.com/GoogleCloudPlatform/healthcare-federated-access-services/lib/storage" /* copybara-comment: storage */
)

// extractVars extracts variables from a request.
// defined here to faciliated testing.
// TODO: do not rely on registeration of routes at global mux for parsing names,
// pass it explicitly.
var extractVars = mux.Vars

// Options contains the information about a handler service.
// Essentially the service interface + some options to the HTTP wrapper for it.
type Options struct {
	TypeName            string
	NameField           string
	PathPrefix          string
	HasNamedIdentifiers bool
	NameChecker         map[string]*regexp.Regexp
	Service             Service
}

// Service is the role interface for a service that will be wrapped.
type Service interface {
	Setup(r *http.Request, tx storage.Tx) (int, error)
	// TODO: Have LookupItem() return an error instead, so different errors can be handled
	// properly, e.g. permission denied error vs. lookup error.
	LookupItem(r *http.Request, name string, vars map[string]string) bool
	NormalizeInput(r *http.Request, name string, vars map[string]string) error
	Get(r *http.Request, name string) (proto.Message, error)
	Post(r *http.Request, name string) (proto.Message, error)
	Put(r *http.Request, name string) (proto.Message, error)
	Patch(r *http.Request, name string) (proto.Message, error)
	Remove(r *http.Request, name string) (proto.Message, error)
	CheckIntegrity(r *http.Request) *status.Status
	Save(r *http.Request, tx storage.Tx, name string, vars map[string]string, desc, typeName string) error
}

// MakeHandler created a HTTP handler wrapper around a given service.
func MakeHandler(s storage.Store, opts *Options) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := Process(s, opts, r)
		if err != nil {
			httputils.WriteError(w, err)
			return
		}
		if resp != nil {
			httputils.WriteResp(w, resp)
		}
	}
}

// Process computes the response for a request.
func Process(s storage.Store, opts *Options, r *http.Request) (_ proto.Message, ferr error) {
	hi := opts.Service
	var op func(*http.Request, string) (proto.Message, error)
	switch r.Method {
	case http.MethodGet:
		op = hi.Get
	case http.MethodPost:
		op = hi.Post
	case http.MethodPut:
		op = hi.Put
	case http.MethodPatch:
		op = hi.Patch
	case http.MethodDelete:
		op = hi.Remove
	default:
		return nil, status.Errorf(codes.InvalidArgument, "request method not supported: %q", r.Method)
	}

	// TODO: move inside each service and don't pass NameChecker here.
	name, vars, err := ValidateResourceName(r, opts.NameField, opts.NameChecker)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	typ := opts.TypeName
	desc := r.Method + " " + typ

	tx, err := s.Tx(r.Method != http.MethodGet)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "service dependencies not available; try again later")
	}
	defer func() {
		err := tx.Finish()
		if ferr == nil {
			ferr = err
		}
	}()

	// Get rid of Setup and move creation of transaction inside service methods.
	//
	if _, err = hi.Setup(r, tx); err != nil {
		return nil, err
	}

	// TODO: Replace NormalizeInput with a ParseReq that returns a request proto message.
	// TODO: Explicitly pass the message to the service methods.
	if err := hi.NormalizeInput(r, name, vars); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	// TODO: get rid of LookupItem and move this inside the service methods.
	exists := hi.LookupItem(r, name, vars)
	switch r.Method {
	case http.MethodPost:
		if exists {
			return nil, status.Errorf(codes.AlreadyExists, "%s already exists: %q", typ, name)
		}
	case http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodDelete:
		if !exists {
			if opts.HasNamedIdentifiers {
				return nil, status.Errorf(codes.NotFound, "%s not found: %q", typ, name)
			}
			return nil, status.Errorf(codes.NotFound, "%s not found", typ)
		}
	}

	if r.Method == http.MethodGet {
		resp, err := op(r, name)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		return resp, nil
	}

	resp, err := RunRMWTx(r, tx, op, hi.CheckIntegrity, hi.Save, name, vars, typ, desc)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ValidateResourceName checks if the resource name is valid.
// Returns the resource name and vars in it.
func ValidateResourceName(r *http.Request, field string, nameRE map[string]*regexp.Regexp) (string, map[string]string, error) {
	nameVar := "name"
	if len(field) > 0 {
		nameVar = field
	}
	vars := extractVars(r)
	name := vars[nameVar]

	for k, v := range vars {
		if err := httputils.CheckName(k, v, nameRE); err != nil {
			return "", nil, err
		}
	}
	return name, vars, nil
}

// RunRMWTx performs a RMW operation.
// Saves the transaction after performing integraty check.
// Rolls back the transaction on any failure.
// TODO: move outside this package. Service handlers should call it.
func RunRMWTx(
	r *http.Request,
	tx storage.Tx,
	op func(*http.Request, string) (proto.Message, error),
	check func(*http.Request) *status.Status,
	save func(*http.Request, storage.Tx, string, map[string]string, string, string) error,
	name string,
	vars map[string]string,
	typ string,
	desc string,
) (proto.Message, error) {
	resp, err := op(r, name)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if st := check(r); st != nil {
		tx.Rollback()
		return nil, st.Err()
	}
	if err := save(r, tx, name, vars, desc, typ); err != nil {
		tx.Rollback()
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return resp, nil
}
