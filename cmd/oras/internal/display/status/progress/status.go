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
	"sync"
	"time"

	"github.com/morikuni/aec"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/cmd/oras/internal/display/status/progress/humanize"
)

const (
	barLength    = 20
	speedLength  = 7    // speed_size(4) + space(1) + speed_unit(2)
	zeroDuration = "0s" // default zero value of time.Duration.String()
)

var (
	spinnerColor  = aec.LightYellowF
	doneMarkColor = aec.LightGreenF
	progressColor = aec.LightBlueB
	failureColor  = aec.LightRedF
)

// status is the model to present the progress of an operation.
type status struct {
	lock sync.RWMutex
	done bool  // true if the operation is succeeded
	err  error // non-nil if the operation fails

	mark      spinner
	text      string
	startTime time.Time
	endTime   time.Time

	descriptor ocispec.Descriptor
	offset     int64
	total      humanize.Bytes
	speed      *speedWindow
}

// newStatus generates a base empty status.
func newStatus(desc ocispec.Descriptor) *status { _ = "STUB: not implemented"; return nil }

// Render returns human-readable TTY strings of the status.
// Format:
//
//	[left--------------------------------------------][margin][right---------------------------------]
//	mark(1) bar(22) speed(8) action(<=11) name(<=126)        size_per_size(<=13) percent(8) time(>=6)
//	 └─ digest(72)
func (s *status) Render(width int) [2]string { _ = "STUB: not implemented"; return nil }

// obtain object name

// calculate the progress percentage

// 100%, show exact size

// not started, show 0%

// 0 byte, show 100%

// 0% ~ 99%, show 2-digit precision

// render the left side of the primary line

// manually calculate the string length due to the color escape sequence

// bar + wrapper(2) + space(1) + speed + "/s"(2) + wrapper(2) = len(bar) + len(speed) + 7

// mark(1) + space(1) + prompt + space(1) + name = len(prompt) + len(name) + 3

// render the right side of the primary line

// render view

// hide partial name with one space left

// calculateSpeed calculates the speed of the progress and update last status.
// caller must hold the lock.
func (s *status) calculateSpeed() humanize.Bytes {
	_ = "STUB: not implemented"

	// not started
	return *new(humanize.Bytes)
}

// durationString returns a viewable TTY string of the status with duration.
func (s *status) durationString() string { _ = "STUB: not implemented"; return "" }

// statusUpdate is a function to update the status.
type statusUpdate func(*status)

// updateStatusMessage returns a statusUpdate to update the status message.
// Optionally, it can update the offset of the status.
func updateStatusMessage(text string, offset int64) statusUpdate {
	_ = "STUB: not implemented"
	return *new(statusUpdate)
}

// updateStatusStartTime returns a statusUpdate to update the status start time.
func updateStatusStartTime() statusUpdate { _ = "STUB: not implemented"; return *new(statusUpdate) }

// updateStatusEndTime returns a statusUpdate to update the status end time.
func updateStatusEndTime() statusUpdate { _ = "STUB: not implemented"; return *new(statusUpdate) }

// updateStatusError returns a statusUpdate to update the status error.
func updateStatusError(err error) statusUpdate {
	_ = "STUB: not implemented"
	return *new(statusUpdate)
}
