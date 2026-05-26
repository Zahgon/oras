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

package text

import (
	"time"

	"oras.land/oras/cmd/oras/internal/output"
)

// RestoreHandler handles text metadata output for restore command.
type RestoreHandler struct {
	printer *output.Printer
	dryRun  bool
}

// NewRestoreHandler creates a new RestoreHandler.
func NewRestoreHandler(printer *output.Printer, dryRun bool) *RestoreHandler {
	_ = "STUB: not implemented"
	return nil
}

// OnTarLoaded implements metadata.RestoreHandler.
func (rh *RestoreHandler) OnTarLoaded(path string, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTagsFound implements metadata.RestoreHandler.
func (rh *RestoreHandler) OnTagsFound(tags []string) error { _ = "STUB: not implemented"; return nil }

// print small number of tags in one line

// print large number of tags line by line

// OnArtifactPushed implements metadata.RestoreHandler.
func (rh *RestoreHandler) OnArtifactPushed(tag string, referrerCount int) error {
	_ = "STUB: not implemented"
	return nil
}

// OnRestoreCompleted implements metadata.RestoreHandler.
func (rh *RestoreHandler) OnRestoreCompleted(tagsCount int, repo string, duration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Render implements metadata.RestoreHandler.
func (rh *RestoreHandler) Render() error { _ = "STUB: not implemented"; return nil }
