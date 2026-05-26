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

package match

import (
	"github.com/onsi/gomega/gbytes"
)

// status represents the expected value of first field in the status log.
type status = string

// StateKey represents the expected value of second and third fields in status log.
type StateKey struct {
	Digest string
	Name   string
}

type state uintptr

var lastState uintptr

func newState() state { _ = "STUB: not implemented"; return *new(state) }

type edge struct {
	from state
	to   state
}

// stateMachine with edges named after known status.
type stateMachine struct {
	edges map[status][]edge
	start state
	end   state
}

func newStateMachine(cmd string) *stateMachine { _ = "STUB: not implemented"; return nil }

// prepare edges

// for `manifest push` and `blob push`
// TODO: refactor the matcher to match full command like `manifest push`, `manifest delete`, etc.
// Tracking issue: https://github.com/oras-project/oras/issues/1571

func findState(from state, edges []edge) *edge { _ = "STUB: not implemented"; return nil }

func (opts *stateMachine) addPath(statuses ...string) { _ = "STUB: not implemented"; return }

// new edge

type statusMatcher struct {
	states       map[StateKey]state
	endResult    map[status][]StateKey
	successCount int
	verbose      bool

	*stateMachine
}

// NewStatusMatcher generates a instance for matchable status logs.
func NewStatusMatcher(keys []StateKey, cmd string, verbose bool, expectSuccessCount int) *statusMatcher {
	_ = "STUB: not implemented"
	return nil
}

// switchState moves a node forward in the state machine graph.
func (s *statusMatcher) switchState(st status, key StateKey) {
	_ = "STUB: not implemented"
	// load state
	return
}

// find next

// switch

// collect last state for matching

// Match checks text status output.
func (s *statusMatcher) Match(got *gbytes.Buffer) { _ = "STUB: not implemented"; return }

// get state key

// media type is hidden, add it

// ignore other logs
