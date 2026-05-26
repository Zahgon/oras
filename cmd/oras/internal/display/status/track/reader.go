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

package track

import (
	"io"
	"os"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/internal/progress"
)

// Reader is a tracked io.Reader.
type Reader struct {
	io.Reader
	tracker progress.Tracker
	manager progress.Manager
}

// NewReader returns a new reader with tracked progress.
func NewReader(r io.Reader, descriptor ocispec.Descriptor, actionPrompt string, donePrompt string, tty *os.File) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newReader(r io.Reader, descriptor ocispec.Descriptor, manager progress.Manager) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tracker returns the progress tracker.
func (r *Reader) Tracker() progress.Tracker {
	_ = "STUB: not implemented"

	// StopTracker stops the messenger channel.
	return *new(progress.Tracker)
}

func (r *Reader) StopTracker() { _ = "STUB: not implemented"; return }

// StopManager stops the messenger channel and related manager.
func (r *Reader) StopManager() { _ = "STUB: not implemented"; return }
