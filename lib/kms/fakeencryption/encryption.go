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

// Package fakeencryption is using for testing
package fakeencryption

import "context"

// FakeEncryption does not encrypt data.
type FakeEncryption struct{}

// New creates a new FakeEncryption.
func New() *FakeEncryption {
	return &FakeEncryption{}
}

func (s *FakeEncryption) Encrypt(ctx context.Context, data []byte, additionalAuthData string) ([]byte, error) {
	return data, nil
}

func (s *FakeEncryption) Decrypt(ctx context.Context, encrypted []byte, additionalAuthData string) ([]byte, error) {
	return encrypted, nil
}
