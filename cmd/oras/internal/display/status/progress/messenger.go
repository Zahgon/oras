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

package progress

import "oras.land/oras/internal/progress"

// messenger is progress message channel.
type messenger struct {
	update  chan statusUpdate
	closed  bool
	prompts map[progress.State]string
}

// Update sends the status to the message channel.
func (m *messenger) Update(status progress.Status) error { _ = "STUB: not implemented"; return nil }

// drop message if channel is full

// Fail sends the error to the message channel.
func (m *messenger) Fail(err error) error { _ = "STUB: not implemented"; return nil }

// Close marks the progress as completed and closes the message channel.
func (m *messenger) Close() error { _ = "STUB: not implemented"; return nil }
