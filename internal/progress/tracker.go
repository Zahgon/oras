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
	"io"
)

// Tracker updates the status of a descriptor.
type Tracker interface {
	io.Closer

	// Update updates the status of the descriptor.
	Update(status Status) error

	// Fail marks the descriptor as failed.
	// Fail should return nil on successful failure marking.
	Fail(err error) error
}

// TrackerFunc is an adapter to allow the use of ordinary functions as Trackers.
// If f is a function with the appropriate signature, TrackerFunc(f) is a
// [Tracker] that calls f.
type TrackerFunc func(status Status, err error) error

// Close closes the tracker.
func (f TrackerFunc) Close() error {
	_ = "STUB: not implemented"

	// Update updates the status of the descriptor.
	return nil
}

func (f TrackerFunc) Update(status Status) error { _ = "STUB: not implemented"; return nil }

// Fail marks the descriptor as failed.
func (f TrackerFunc) Fail(err error) error { _ = "STUB: not implemented"; return nil }

// Start starts tracking the transmission.
func Start(t Tracker) error { _ = "STUB: not implemented"; return nil }

// Done marks the transmission as complete.
// Done should be called after the transmission is complete.
// Note: Reading all content from the reader does not imply the transmission is
// complete.
func Done(t Tracker) error { _ = "STUB: not implemented"; return nil }

// TrackReader bind a reader with a tracker.
func TrackReader(t Tracker, r io.Reader) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

// readTracker tracks the transmission based on the read operation.
type readTracker struct {
	base    io.Reader
	tracker Tracker
	offset  int64
}

// Read reads from the base reader and updates the status.
// On partial read, the tracker treats it as two reads: a successful read with
// status update and a failed read with failure report.
func (rt *readTracker) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// readTrackerWriteTo is readTracker with WriteTo support.
type readTrackerWriteTo struct {
	readTracker
}

// WriteTo writes to the base writer and updates the status.
// On partial write, the tracker treats it as two writes: a successful write
// with status update and a failed write with failure report.
func (rt *readTrackerWriteTo) WriteTo(w io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// writeTracker tracks the transmission based on the write operation.
type writeTracker struct {
	base       io.Writer
	tracker    Tracker
	offset     int64
	trackerErr error
}

// Write writes to the base writer and updates the status.
// On partial write, the tracker treats it as two writes: a successful write
// with status update and a failed write with failure report.
func (wt *writeTracker) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
