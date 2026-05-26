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

	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
)

func discardStopTrack() error {
	_ = "STUB: not implemented"

	// DiscardHandler is a no-op handler that discards all status updates.
	return nil
}

type DiscardHandler struct{}

// NewDiscardHandler returns a new no-op handler.
func NewDiscardHandler() DiscardHandler {
	_ = "STUB: not implemented"
	return *

	// OnFileLoading is called before a file is being loaded.
	new(DiscardHandler)
}

func (DiscardHandler) OnFileLoading(string) error {
	_ = "STUB: not implemented"

	// OnEmptyArtifact is called when no file is loaded for an artifact push.
	return nil
}

func (DiscardHandler) OnEmptyArtifact() error {
	_ = "STUB: not implemented"

	// TrackTarget returns a target with status tracking.
	return nil
}

func (DiscardHandler) TrackTarget(gt oras.GraphTarget) (oras.GraphTarget, StopTrackTargetFunc, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), *new(StopTrackTargetFunc), nil
}

// OnCopySkipped is called when an object already exists.
func (DiscardHandler) OnCopySkipped(_ context.Context, _ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// PreCopy implements PreCopy of CopyHandler.
	return nil
}

func (DiscardHandler) PreCopy(_ context.Context, _ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// PostCopy implements PostCopy of CopyHandler.
	return nil
}

func (DiscardHandler) PostCopy(_ context.Context, _ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnNodeDownloading implements PullHandler.
	return nil
}

func (DiscardHandler) OnNodeDownloading(_ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnNodeDownloaded implements PullHandler.
	return nil
}

func (DiscardHandler) OnNodeDownloaded(_ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnNodeRestored implements PullHandler.
	return nil
}

func (DiscardHandler) OnNodeRestored(_ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnNodeProcessing implements PullHandler.
	return nil
}

func (DiscardHandler) OnNodeProcessing(_ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnNodeProcessing implements PullHandler.
	return nil
}

func (DiscardHandler) OnNodeSkipped(_ ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnFetching implements referenceFetchHandler.
	return nil
}

func (DiscardHandler) OnFetching(string) error {
	_ = "STUB: not implemented"

	// OnFetched implements referenceFetchHandler.
	return nil
}

func (DiscardHandler) OnFetched(string, ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnManifestPushSkipped implements ManifestPushHandler.
	return nil
}

func (DiscardHandler) OnManifestPushSkipped() error {
	_ = "STUB: not implemented"

	// OnManifestPushing implements ManifestPushHandler.
	return nil
}

func (DiscardHandler) OnManifestPushing() error {
	_ = "STUB: not implemented"

	// OnManifestPushed implements ManifestPushHandler.
	return nil
}

func (DiscardHandler) OnManifestPushed() error {
	_ = "STUB: not implemented"

	// OnManifestRemoved implements ManifestIndexUpdateHandler.
	return nil
}

func (DiscardHandler) OnManifestRemoved(digest.Digest) error {
	_ = "STUB: not implemented"

	// OnManifestAdded implements ManifestIndexUpdateHandler.
	return nil
}

func (DiscardHandler) OnManifestAdded(string, ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnIndexMerged implements ManifestIndexUpdateHandler.
	return nil
}

func (DiscardHandler) OnIndexMerged(string, ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnIndexPacked implements ManifestIndexCreateHandler.
	return nil
}

func (DiscardHandler) OnIndexPacked(ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// OnIndexPushed implements ManifestIndexCreateHandler.
	return nil
}

func (DiscardHandler) OnIndexPushed(string) error {
	_ = "STUB: not implemented"

	// OnBlobExists implements BlobPushHandler.
	return nil
}

func (DiscardHandler) OnBlobExists() error {
	_ = "STUB: not implemented"

	// OnBlobUploading implements BlobPushHandler.
	return nil
}

func (DiscardHandler) OnBlobUploading() error {
	_ = "STUB: not implemented"

	// OnBlobUploaded implements BlobPushHandler.
	return nil
}

func (DiscardHandler) OnBlobUploaded() error {
	_ = "STUB: not implemented"

	// StartTracking implements BlobPushHandler.
	return nil
}

func (DiscardHandler) StartTracking(gt oras.GraphTarget) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"

	// StopTracking implements BlobPushHandler.
	return *new(oras.GraphTarget), nil
}

func (DiscardHandler) StopTracking() error { _ = "STUB: not implemented"; return nil }
