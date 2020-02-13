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

package process

import (
	"context"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/healthcare-federated-access-services/lib/clouds" /* copybara-comment: clouds */
	"github.com/GoogleCloudPlatform/healthcare-federated-access-services/lib/common" /* copybara-comment: common */
	"github.com/GoogleCloudPlatform/healthcare-federated-access-services/lib/storage" /* copybara-comment: storage */
	pb "github.com/GoogleCloudPlatform/healthcare-federated-access-services/proto/process/v1" /* copybara-comment: go_proto */
)

const (
	keyScheduleFrequency = time.Hour
)

// KeyGC is a service account key garbage collector background process.
type KeyGC struct {
	name    string
	am      clouds.AccountManager
	process *Process
	wait    func(ctx context.Context, duration time.Duration) bool
}

// NewKeyGC creates a new key garbage collector.
func NewKeyGC(name string, warehouse clouds.AccountManager, store storage.Store, maxRequestedTTL time.Duration, keysPerAccount int) *KeyGC {
	k := &KeyGC{
		name: name,
		am:   warehouse,
	}
	defaultParams := &pb.Process_Params{
		IntParams: map[string]int64{
			"maxRequestedTtl": int64(maxRequestedTTL.Seconds()),
			"keysPerAccount":  int64(keysPerAccount),
			"keyTtl":          int64(common.KeyTTL(maxRequestedTTL, keysPerAccount).Seconds()),
		},
	}
	k.process = NewProcess(name, k, store, keyScheduleFrequency, defaultParams)
	return k
}

// RegisterProject adds a project to the state for workers to process.
func (k *KeyGC) RegisterProject(project string, params *pb.Process_Params) (*pb.Process_Project, error) {
	return k.process.RegisterProject(project, params)
}

// UnregisterProject (eventually) removes a project from the active state, and allows cleanup work to be performed.
func (k *KeyGC) UnregisterProject(projectName string) error {
	return k.process.UnregisterProject(projectName)
}

// UpdateSettings alters resource management settings.
func (k *KeyGC) UpdateSettings(maxRequestedTTL time.Duration, keysPerAccount int) error {
	params := &pb.Process_Params{
		IntParams: map[string]int64{
			"maxRequestedTtl": int64(maxRequestedTTL.Seconds()),
			"keysPerAccount":  int64(keysPerAccount),
			"keyTtl":          int64(common.KeyTTL(maxRequestedTTL, keysPerAccount).Seconds()),
		},
	}
	return k.process.UpdateSettings(params)
}

// WaitCondition registers a callback that is called and checks conditions before every wait cycle.
func (k *KeyGC) WaitCondition(fn func(ctx context.Context, duration time.Duration) bool) {
	k.wait = fn
}

// Run schedules a background process. Typically this will be on its own go routine.
func (k *KeyGC) Run(ctx context.Context) {
	k.process.Run(ctx)
}

// ProcessActiveProject has a worker perform the work needed to process an active project.
func (k *KeyGC) ProcessActiveProject(ctx context.Context, state *pb.Process, projectName string, project *pb.Process_Project, process *Process) error {
	return k.am.GetServiceAccounts(ctx, projectName, func(sa *clouds.Account) bool {
		if !isGarbageCollectAccount(sa) {
			return true
		}
		run := Continue
		process.AddProjectStats(1, "accounts", projectName, state)
		keyTTL := project.Params.IntParams["keyTtl"]
		keysPerAccount := project.Params.IntParams["keysPerAccount"]
		got, rm, err := k.am.ManageAccountKeys(ctx, projectName, sa.ID, 0, time.Duration(keyTTL)*time.Second, int(keysPerAccount))
		if err != nil {
			run = process.AddProjectError(err, projectName, state)
		}
		process.AddProjectStats(float64(got), "keysKept", projectName, state)
		process.AddProjectStats(float64(rm), "keysRemoved", projectName, state)
		return run == Continue
	})
}

// CleanupProject has a worker perform the work needed to clean up a project that was active previously.
func (k *KeyGC) CleanupProject(ctx context.Context, state *pb.Process, projectName string, process *Process) error {
	return k.am.GetServiceAccounts(ctx, projectName, func(sa *clouds.Account) bool {
		if isGarbageCollectAccount(sa) {
			if err := k.am.RemoveServiceAccount(ctx, projectName, sa.ID); err == nil {
				process.AddProjectStats(1, "accountsRemoved", projectName, state)
			} else {
				process.AddProjectStats(1, "accountsDirty", projectName, state)
				process.AddProjectError(err, projectName, state)
				return false
			}
		}
		return true
	})
}

// Wait indicates that the worker should wait for the next active cycle to begin.
func (k *KeyGC) Wait(ctx context.Context, duration time.Duration) bool {
	if k.wait != nil && !k.wait(ctx, duration) {
		return false
	}
	time.Sleep(duration)
	return true
}

func isGarbageCollectAccount(sa *clouds.Account) bool {
	return strings.Contains(sa.DisplayName, "@") || strings.Contains(sa.DisplayName, "|")
}