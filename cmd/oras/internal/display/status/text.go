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
	"sync"

	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content"
	"oras.land/oras/cmd/oras/internal/output"
)

// TextPushHandler handles text status output for push events.
type TextPushHandler struct {
	printer   *output.Printer
	committed *sync.Map
	fetcher   content.Fetcher
}

// NewTextPushHandler returns a new handler for push command.
func NewTextPushHandler(printer *output.Printer, fetcher content.Fetcher) PushHandler {
	_ = "STUB: not implemented"
	return *new(PushHandler)
}

// OnFileLoading is called when a file is being prepared for upload.
func (ph *TextPushHandler) OnFileLoading(name string) error { _ = "STUB: not implemented"; return nil }

// OnEmptyArtifact is called when an empty artifact is being uploaded.
func (ph *TextPushHandler) OnEmptyArtifact() error { _ = "STUB: not implemented"; return nil }

// TrackTarget returns a tracked target.
func (ph *TextPushHandler) TrackTarget(gt oras.GraphTarget) (oras.GraphTarget, StopTrackTargetFunc, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), *new(StopTrackTargetFunc), nil
}

// OnCopySkipped is called when an object already exists.
func (ph *TextPushHandler) OnCopySkipped(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCopy implements PreCopy of CopyHandler.
func (ph *TextPushHandler) PreCopy(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PostCopy implements PostCopy of CopyHandler.
func (ph *TextPushHandler) PostCopy(ctx context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// NewTextAttachHandler returns a new handler for attach command.
func NewTextAttachHandler(printer *output.Printer, fetcher content.Fetcher) AttachHandler {
	_ = "STUB: not implemented"
	return *new(AttachHandler)
}

// TextPullHandler handles text status output for pull events.
type TextPullHandler struct {
	printer *output.Printer
}

// TrackTarget implements PullHandler.
func (ph *TextPullHandler) TrackTarget(gt oras.GraphTarget) (oras.GraphTarget, StopTrackTargetFunc, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), *new(StopTrackTargetFunc), nil
}

// OnNodeDownloading implements PullHandler.
func (ph *TextPullHandler) OnNodeDownloading(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnNodeDownloaded implements PullHandler.
func (ph *TextPullHandler) OnNodeDownloaded(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnNodeRestored implements PullHandler.
func (ph *TextPullHandler) OnNodeRestored(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnNodeProcessing implements PullHandler.
func (ph *TextPullHandler) OnNodeProcessing(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnNodeSkipped implements PullHandler.
func (ph *TextPullHandler) OnNodeSkipped(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// NewTextPullHandler returns a new handler for pull command.
func NewTextPullHandler(printer *output.Printer) PullHandler {
	_ = "STUB: not implemented"
	return *new(PullHandler)
}

// TextCopyHandler handles text status output for push events.
type TextCopyHandler struct {
	printer   *output.Printer
	committed *sync.Map
	fetcher   content.Fetcher
}

// NewTextCopyHandler returns a new handler for push command.
func NewTextCopyHandler(printer *output.Printer, fetcher content.Fetcher) CopyHandler {
	_ = "STUB: not implemented"
	return *new(CopyHandler)
}

// StartTracking starts a tracked target from a graph target.
func (ch *TextCopyHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"

	// StopTracking ends the copy tracking for the target.
	return *new(oras.GraphTarget), nil
}

func (ch *TextCopyHandler) StopTracking() error {
	_ = "STUB: not implemented"

	// OnCopySkipped is called when an object already exists.
	return nil
}

func (ch *TextCopyHandler) OnCopySkipped(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCopy implements PreCopy of CopyHandler.
func (ch *TextCopyHandler) PreCopy(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PostCopy implements PostCopy of CopyHandler.
func (ch *TextCopyHandler) PostCopy(ctx context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnMounted implements OnMounted of CopyHandler.
func (ch *TextCopyHandler) OnMounted(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// TextBackupHandler handles text status output for backup events.
type TextBackupHandler struct {
	printer   *output.Printer
	committed *sync.Map
	fetcher   content.Fetcher
}

// NewTextBackupHandler returns a new text handler for backup command.
func NewTextBackupHandler(printer *output.Printer, fetcher content.Fetcher) BackupHandler {
	_ = "STUB: not implemented"
	return *new(BackupHandler)
}

// OnCopySkipped implements OnCopySkipped of BackupHandler.
func (tbh *TextBackupHandler) OnCopySkipped(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCopy implements PreCopy of BackupHandler.
func (tbh *TextBackupHandler) PreCopy(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PostCopy implements PostCopy of BackupHandler.
func (tbh *TextBackupHandler) PostCopy(ctx context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// StartTracking implements StartTracking of BackupHandler.
func (tbh *TextBackupHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"

	// StopTracking implements StopTracking of BackupHandler.
	return *new(oras.GraphTarget), nil
}

func (tbh *TextBackupHandler) StopTracking() error {
	_ = "STUB: not implemented"

	// TextRestoreHandler handles text status output for restore events.
	return nil
}

type TextRestoreHandler struct {
	printer   *output.Printer
	committed *sync.Map
	fetcher   content.Fetcher
}

// NewTextRestoreHandler returns a new text handler for restore command.
func NewTextRestoreHandler(printer *output.Printer, fetcher content.Fetcher) RestoreHandler {
	_ = "STUB: not implemented"
	return *new(RestoreHandler)
}

// OnCopySkipped implements OnCopySkipped of RestoreHandler.
func (trh *TextRestoreHandler) OnCopySkipped(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PreCopy implements PreCopy of RestoreHandler.
func (trh *TextRestoreHandler) PreCopy(_ context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// PostCopy implements PostCopy of RestoreHandler.
func (trh *TextRestoreHandler) PostCopy(ctx context.Context, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// StartTracking implements StartTracking of RestoreHandler.
func (trh *TextRestoreHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"

	// StopTracking implements StopTracking of RestoreHandler.
	return *new(oras.GraphTarget), nil
}

func (trh *TextRestoreHandler) StopTracking() error {
	_ = "STUB: not implemented"

	// TextManifestPushHandler handles text status output for manifest push events.
	return nil
}

type TextManifestPushHandler struct {
	desc    ocispec.Descriptor
	printer *output.Printer
}

// NewTextManifestPushHandler returns a new handler for manifest push command.
func NewTextManifestPushHandler(printer *output.Printer, desc ocispec.Descriptor) ManifestPushHandler {
	_ = "STUB: not implemented"
	return *new(ManifestPushHandler)
}

func (mph *TextManifestPushHandler) OnManifestPushSkipped() error {
	_ = "STUB: not implemented"
	return nil
}

func (mph *TextManifestPushHandler) OnManifestPushing() error {
	_ = "STUB: not implemented"
	return nil
}

func (mph *TextManifestPushHandler) OnManifestPushed() error { _ = "STUB: not implemented"; return nil }

// TextManifestIndexCreateHandler handles text status output for manifest index create events.
type TextManifestIndexCreateHandler struct {
	printer *output.Printer
}

// NewTextManifestIndexCreateHandler returns a new handler for manifest index create command.
func NewTextManifestIndexCreateHandler(printer *output.Printer) ManifestIndexCreateHandler {
	_ = "STUB: not implemented"
	return *new(ManifestIndexCreateHandler)
}

// OnFetching implements ManifestIndexCreateHandler.
func (mich *TextManifestIndexCreateHandler) OnFetching(source string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnFetched implements ManifestIndexCreateHandler.
func (mich *TextManifestIndexCreateHandler) OnFetched(ref string, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnIndexPacked implements ManifestIndexCreateHandler.
func (mich *TextManifestIndexCreateHandler) OnIndexPacked(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnIndexPushed implements ManifestIndexCreateHandler.
func (mich *TextManifestIndexCreateHandler) OnIndexPushed(path string) error {
	_ = "STUB: not implemented"
	return nil
}

// TextManifestIndexUpdateHandler handles text status output for manifest index update events.
type TextManifestIndexUpdateHandler struct {
	printer *output.Printer
}

// NewTextManifestIndexUpdateHandler returns a new handler for manifest index create command.
func NewTextManifestIndexUpdateHandler(printer *output.Printer) ManifestIndexUpdateHandler {
	_ = "STUB: not implemented"
	return *new(ManifestIndexUpdateHandler)
}

// OnFetching implements ManifestIndexUpdateHandler.
func (miuh *TextManifestIndexUpdateHandler) OnFetching(ref string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnFetched implements ManifestIndexUpdateHandler.
func (miuh *TextManifestIndexUpdateHandler) OnFetched(ref string, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnManifestRemoved implements ManifestIndexUpdateHandler.
func (miuh *TextManifestIndexUpdateHandler) OnManifestRemoved(digest digest.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

// OnManifestAdded implements ManifestIndexUpdateHandler.
func (miuh *TextManifestIndexUpdateHandler) OnManifestAdded(ref string, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnIndexMerged implements ManifestIndexUpdateHandler.
func (miuh *TextManifestIndexUpdateHandler) OnIndexMerged(ref string, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnIndexPacked implements ManifestIndexUpdateHandler.
func (miuh *TextManifestIndexUpdateHandler) OnIndexPacked(desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnIndexPushed implements ManifestIndexUpdateHandler.
func (miuh *TextManifestIndexUpdateHandler) OnIndexPushed(indexRef string) error {
	_ = "STUB: not implemented"
	return nil
}

// TextBlobPushHandler handles text status output for blob push events.
type TextBlobPushHandler struct {
	desc    ocispec.Descriptor
	printer *output.Printer
}

// NewTextBlobPushHandler returns a new handler for blob push command.
func NewTextBlobPushHandler(printer *output.Printer, desc ocispec.Descriptor) BlobPushHandler {
	_ = "STUB: not implemented"
	return *new(BlobPushHandler)
}

// OnBlobExists implements BlobPushHandler.
func (bph *TextBlobPushHandler) OnBlobExists() error { _ = "STUB: not implemented"; return nil }

// OnBlobUploading implements BlobPushHandler.
func (bph *TextBlobPushHandler) OnBlobUploading() error { _ = "STUB: not implemented"; return nil }

// OnBlobUploaded implements BlobPushHandler.
func (bph *TextBlobPushHandler) OnBlobUploaded() error { _ = "STUB: not implemented"; return nil }

// StartTracking implements BlobPushHandler.
func (bph *TextBlobPushHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"

	// StopTracking implements BlobPushHandler.
	return *new(oras.GraphTarget), nil
}

func (bph *TextBlobPushHandler) StopTracking() error { _ = "STUB: not implemented"; return nil }
