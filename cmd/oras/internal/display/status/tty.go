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

package status

import (
	"context"
	"os"
	"sync"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content"
	"oras.land/oras/cmd/oras/internal/display/status/track"
)

// TTYPushHandler handles TTY status output for push command.
type TTYPushHandler struct {
	tty       *os.File
	tracked   track.GraphTarget
	committed *sync.Map
	fetcher   content.Fetcher
}

// NewTTYPushHandler returns a new handler for push status events.
func NewTTYPushHandler(tty *os.File, fetcher content.Fetcher) PushHandler {
	_ = "STUB: not implemented"
	return *new(PushHandler)
}

// OnFileLoading is called before loading a file.
func (ph *TTYPushHandler) OnFileLoading(_ string) error {
	_ = "STUB: not implemented"

	// OnEmptyArtifact is called when no file is loaded for an artifact push.
	return nil
}

func (ph *TTYPushHandler) OnEmptyArtifact() error {
	_ = "STUB: not implemented"

	// TrackTarget returns a tracked target.
	return nil
}

func (ph *TTYPushHandler) TrackTarget(gt oras.GraphTarget) (oras.GraphTarget, StopTrackTargetFunc, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), *new(StopTrackTargetFunc), nil
}

// OnCopySkipped is called when an object already exists.
func (ph *TTYPushHandler) OnCopySkipped(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCopy implements PreCopy of CopyHandler.
func (ph *TTYPushHandler) PreCopy(_ context.Context, _ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// PostCopy implements PostCopy of CopyHandler.
	return nil
}

func (ph *TTYPushHandler) PostCopy(ctx context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// NewTTYAttachHandler returns a new handler for attach status events.
func NewTTYAttachHandler(tty *os.File, fetcher content.Fetcher) AttachHandler {
	_ = "STUB: not implemented"
	return *new(AttachHandler)
}

// TTYPullHandler handles TTY status output for pull events.
type TTYPullHandler struct {
	tty     *os.File
	tracked track.GraphTarget
}

// NewTTYPullHandler returns a new handler for Pull status events.
func NewTTYPullHandler(tty *os.File) PullHandler {
	_ = "STUB: not implemented"
	return *new(PullHandler)
}

// OnNodeDownloading implements PullHandler.
func (ph *TTYPullHandler) OnNodeDownloading(_ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnNodeDownloaded implements PullHandler.
	return nil
}

func (ph *TTYPullHandler) OnNodeDownloaded(_ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnNodeProcessing implements PullHandler.
	return nil
}

func (ph *TTYPullHandler) OnNodeProcessing(_ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnNodeRestored implements PullHandler.
	return nil
}

func (ph *TTYPullHandler) OnNodeRestored(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnNodeSkipped implements PullHandler.
func (ph *TTYPullHandler) OnNodeSkipped(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// TrackTarget returns a tracked target.
func (ph *TTYPullHandler) TrackTarget(gt oras.GraphTarget) (oras.GraphTarget, StopTrackTargetFunc, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), *new(StopTrackTargetFunc), nil
}

// TTYCopyHandler handles tty status output for copy events.
type TTYCopyHandler struct {
	tty       *os.File
	committed sync.Map
	tracked   track.GraphTarget
}

// NewTTYCopyHandler returns a new handler for copy command.
func NewTTYCopyHandler(tty *os.File) CopyHandler {
	_ = "STUB: not implemented"
	return *new(CopyHandler)
}

// StartTracking returns a tracked target from a graph target.
func (ch *TTYCopyHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), nil
}

// StopTracking ends the copy tracking for the target.
func (ch *TTYCopyHandler) StopTracking() error { _ = "STUB: not implemented"; return nil }

// OnCopySkipped is called when an object already exists.
func (ch *TTYCopyHandler) OnCopySkipped(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCopy implements PreCopy of CopyHandler.
func (ch *TTYCopyHandler) PreCopy(context.Context, ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// PostCopy implements PostCopy of CopyHandler.
	return nil
}

func (ch *TTYCopyHandler) PostCopy(ctx context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnMounted implements OnMounted of CopyHandler.
func (ch *TTYCopyHandler) OnMounted(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// TTYBackupHandler handles tty status output for backup events.
type TTYBackupHandler struct {
	tty       *os.File
	committed *sync.Map
	tracked   track.GraphTarget
	fetcher   content.Fetcher
}

// NewTTYBackupHandler returns a new handler for backup command.
func NewTTYBackupHandler(tty *os.File, fetcher content.Fetcher) BackupHandler {
	_ = "STUB: not implemented"
	return *new(BackupHandler)
}

// StartTracking returns a tracked target from a graph target.
func (bh *TTYBackupHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), nil
}

// StopTracking ends the backup tracking for the target.
func (bh *TTYBackupHandler) StopTracking() error { _ = "STUB: not implemented"; return nil }

// OnCopySkipped implements OnCopySkipped of BackupHandler.
func (bh *TTYBackupHandler) OnCopySkipped(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCopy implements PreCopy of BackupHandler.
func (bh *TTYBackupHandler) PreCopy(_ context.Context, _ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// PostCopy implements PostCopy of BackupHandler.
	return nil
}

func (bh *TTYBackupHandler) PostCopy(ctx context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// TTYRestoreHandler handles tty status output for restore events.
type TTYRestoreHandler struct {
	tty       *os.File
	committed *sync.Map
	tracked   track.GraphTarget
	fetcher   content.Fetcher
}

// NewTTYRestoreHandler returns a new handler for restore command.
func NewTTYRestoreHandler(tty *os.File, fetcher content.Fetcher) RestoreHandler {
	_ = "STUB: not implemented"
	return *new(RestoreHandler)
}

// StartTracking returns a tracked target from a graph target.
func (rh *TTYRestoreHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), nil
}

// StopTracking ends the restore tracking for the target.
func (rh *TTYRestoreHandler) StopTracking() error { _ = "STUB: not implemented"; return nil }

// OnCopySkipped implements OnCopySkipped of RestoreHandler.
func (rh *TTYRestoreHandler) OnCopySkipped(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCopy implements PreCopy of RestoreHandler.
func (rh *TTYRestoreHandler) PreCopy(_ context.Context, _ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// PostCopy implements PostCopy of RestoreHandler.
	return nil
}

func (rh *TTYRestoreHandler) PostCopy(ctx context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// TTYBlobPushHandler handles tty status output for blob push events.
type TTYBlobPushHandler struct {
	desc    ocispec.Descriptor
	tty     *os.File
	tracked track.GraphTarget
}

// NewTTYBlobPushHandler returns a new handler for blob push command.
func NewTTYBlobPushHandler(tty *os.File, desc ocispec.Descriptor) BlobPushHandler {
	_ = "STUB: not implemented"
	return *new(BlobPushHandler)
}

// StartTracking returns a tracked target from a graph target.
func (bph *TTYBlobPushHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), nil
}

// StopTracking ends the blob push tracking for the target.
func (bph *TTYBlobPushHandler) StopTracking() error { _ = "STUB: not implemented"; return nil }

// OnBlobExists implements BlobPushHandler.
func (bph *TTYBlobPushHandler) OnBlobExists() error { _ = "STUB: not implemented"; return nil }

// OnBlobUploading implements BlobPushHandler.
func (bph *TTYBlobPushHandler) OnBlobUploading() error {
	_ = "STUB: not implemented"

	// OnBlobUploaded implements BlobPushHandler.
	return nil
}

func (bph *TTYBlobPushHandler) OnBlobUploaded() error { _ = "STUB: not implemented"; return nil }
