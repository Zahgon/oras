//go:build !windows && !darwin

/*
Copyright The ORAS Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package testutils

import (
	"os"

	containerd "github.com/containerd/console"
)

// NewPty creates a new pty pair for testing, caller is responsible for closing
// the returned device file if err is not nil.
func NewPty() (containerd.Console, *os.File, error) {
	_ = "STUB: not implemented"
	return *new(containerd.Console), nil, nil
}

// MatchPty checks that the output matches the expected strings in specified
// order.
func MatchPty(pty containerd.Console, device *os.File, expected ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// OrderedMatch matches the got with the expected strings in order.
func OrderedMatch(got string, want ...string) error { _ = "STUB: not implemented"; return nil }
