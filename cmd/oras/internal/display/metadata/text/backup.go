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

	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/output"
)

// BackupHandler handles text metadata output for backup events.
type BackupHandler struct {
	printer *output.Printer
	repo    string
}

// NewBackupHandler returns a new handler for backup events.
func NewBackupHandler(repo string, printer *output.Printer) metadata.BackupHandler {
	_ = "STUB: not implemented"
	return *new(metadata.BackupHandler)
}

// OnBackupCompleted implements metadata.BackupHandler.
func (bh *BackupHandler) OnBackupCompleted(tagsCount int, path string, duration time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTarExported implements metadata.BackupHandler.
func (bh *BackupHandler) OnTarExported(path string, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

// OnTarExporting implements metadata.BackupHandler.
func (bh *BackupHandler) OnTarExporting(path string) error { _ = "STUB: not implemented"; return nil }

// OnArtifactPulled implements metadata.BackupHandler.
func (bh *BackupHandler) OnArtifactPulled(tag string, referrerCount int) error {
	_ = "STUB: not implemented"
	// represent duration in a human-readable format
	return nil
}

// OnTagsFound implements metadata.BackupHandler.
func (bh *BackupHandler) OnTagsFound(tags []string) error { _ = "STUB: not implemented"; return nil }

// print small number of tags in one line

// print large number of tags line by line

// Render implements metadata.BackupHandler.
func (bh *BackupHandler) Render() error { _ = "STUB: not implemented"; return nil }
