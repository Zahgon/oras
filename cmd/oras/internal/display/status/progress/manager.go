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

import (
	"errors"
	"os"
	"sync"
	"time"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/cmd/oras/internal/display/status/console"
	"oras.land/oras/internal/progress"
)

const (
	// bufferSize is the size of the status channel buffer.
	bufferSize       = 1
	framePerSecond   = 5
	bufFlushDuration = time.Second / framePerSecond
)

var errManagerStopped = errors.New("progress output manager has already been stopped")

type manager struct {
	status       []*status
	lock         sync.RWMutex // locks status and console, write lock is used for adding new status so that it has a higher priority
	console      console.Console
	updating     sync.WaitGroup
	renderDone   chan struct{}
	renderClosed chan struct{}
	prompts      map[progress.State]string
}

// NewManager initialized a new progress manager.
func NewManager(tty *os.File, prompts map[progress.State]string) (progress.Manager, error) {
	_ = "STUB: not implemented"
	return *new(progress.Manager), nil
}

func newManager(c console.Console, prompts map[progress.State]string) progress.Manager {
	_ = "STUB: not implemented"
	return *new(progress.Manager)
}

func (m *manager) start() { _ = "STUB: not implemented"; return }

func (m *manager) render() { _ = "STUB: not implemented"; return }

// render with culling: only the latter statuses are rendered.

// Track appends a new status with 2-line space for rendering.
func (m *manager) Track(desc ocispec.Descriptor) (progress.Tracker, error) {
	_ = "STUB: not implemented"
	return *new(progress.Tracker), nil
}

func (m *manager) newTracker(s *status) progress.Tracker {
	_ = "STUB: not implemented"
	return *new(progress.Tracker)
}

// Close stops all status and waits for updating and rendering.
func (m *manager) Close() error { _ = "STUB: not implemented"; return nil }

// 1. wait for update to stop

// 2. stop periodic rendering

// 3. wait for the render stop

func (m *manager) closed() bool { _ = "STUB: not implemented"; return false }
