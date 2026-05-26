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

package display

import (
	"io"
	"os"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	fetcher "oras.land/oras-go/v2/content"

	"oras.land/oras/cmd/oras/internal/display/content"
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/display/status"
	"oras.land/oras/cmd/oras/internal/option"
	"oras.land/oras/cmd/oras/internal/output"
)

// NewPushHandler returns status and metadata handlers for push command.
func NewPushHandler(printer *output.Printer, format option.Format, tty *os.File, fetcher fetcher.Fetcher) (status.PushHandler, metadata.PushHandler, error) {
	_ = "STUB: not implemented"
	return *new(status.PushHandler), *new(metadata.PushHandler), nil
}

// NewAttachHandler returns status and metadata handlers for attach command.
func NewAttachHandler(printer *output.Printer, format option.Format, tty *os.File, fetcher fetcher.Fetcher) (status.AttachHandler, metadata.AttachHandler, error) {
	_ = "STUB: not implemented"
	return *new(status.AttachHandler), *new(metadata.AttachHandler), nil
}

// NewPullHandler returns status and metadata handlers for pull command.
func NewPullHandler(printer *output.Printer, format option.Format, path string, tty *os.File) (status.PullHandler, metadata.PullHandler, error) {
	_ = "STUB: not implemented"
	return *new(status.PullHandler), *new(metadata.PullHandler), nil
}

// NewDiscoverHandler returns status and metadata handlers for discover command.
func NewDiscoverHandler(out io.Writer, format option.Format, path string, rawReference string, desc ocispec.Descriptor, verbose bool, tty *os.File) (metadata.DiscoverHandler, error) {
	_ = "STUB: not implemented"
	return *new(metadata.DiscoverHandler), nil
}

// NewManifestFetchHandler returns a manifest fetch handler.
func NewManifestFetchHandler(out io.Writer, format option.Format, outputDescriptor, pretty bool, outputPath string) (metadata.ManifestFetchHandler, content.ManifestFetchHandler, error) {
	_ = "STUB: not implemented"
	return *new(metadata.ManifestFetchHandler), *new(content.ManifestFetchHandler), nil
}

// raw

// json

// go template

// NewTagHandler returns a tag handler.
func NewTagHandler(printer *output.Printer, target option.Target) metadata.TagHandler {
	_ = "STUB: not implemented"
	return *new(metadata.TagHandler)
}

// NewManifestPushHandler returns a manifest push handler.
func NewManifestPushHandler(printer *output.Printer, outputDescriptor bool, _ bool, desc ocispec.Descriptor, target *option.Target) (status.ManifestPushHandler, metadata.ManifestPushHandler) {
	_ = "STUB: not implemented"
	return *new(status.ManifestPushHandler), *new(metadata.ManifestPushHandler)
}

// NewManifestDeleteHandler returns a manifest delete handler.
func NewManifestDeleteHandler(printer *output.Printer, target *option.Target) metadata.ManifestDeleteHandler {
	_ = "STUB: not implemented"
	return *new(metadata.ManifestDeleteHandler)
}

// NewManifestIndexCreateHandler returns status, metadata and content handlers for index create command.
func NewManifestIndexCreateHandler(outputPath string, printer *output.Printer, pretty bool) (status.ManifestIndexCreateHandler, metadata.ManifestIndexCreateHandler, content.ManifestIndexCreateHandler) {
	_ = "STUB: not implemented"
	return *new(status.ManifestIndexCreateHandler), *new(metadata.ManifestIndexCreateHandler), *new(content.ManifestIndexCreateHandler)
}

// NewManifestIndexUpdateHandler returns status, metadata and content handlers for index update command.
func NewManifestIndexUpdateHandler(outputPath string, printer *output.Printer, pretty bool) (
	status.ManifestIndexUpdateHandler,
	metadata.ManifestIndexUpdateHandler,
	content.ManifestIndexUpdateHandler) {
	_ = "STUB: not implemented"
	return *new(status.ManifestIndexUpdateHandler), *new(metadata.ManifestIndexUpdateHandler), *new(content.ManifestIndexUpdateHandler)
}

// NewCopyHandler returns copy handlers.
func NewCopyHandler(printer *output.Printer, tty *os.File, fetcher fetcher.Fetcher) (status.CopyHandler, metadata.CopyHandler) {
	_ = "STUB: not implemented"
	return *new(status.CopyHandler), *new(metadata.CopyHandler)
}

// NewBackupHandler returns backup handlers.
func NewBackupHandler(printer *output.Printer, tty *os.File, repo string, fetcher fetcher.Fetcher) (status.BackupHandler, metadata.BackupHandler) {
	_ = "STUB: not implemented"
	return *new(status.BackupHandler), *new(metadata.BackupHandler)
}

// NewRestoreHandler returns restore handlers.
func NewRestoreHandler(printer *output.Printer, tty *os.File, fetcher fetcher.Fetcher, dryRun bool) (status.RestoreHandler, metadata.RestoreHandler) {
	_ = "STUB: not implemented"
	return *new(status.RestoreHandler), *new(metadata.RestoreHandler)
}

// NewBlobPushHandler returns blob push handlers.
func NewBlobPushHandler(printer *output.Printer, outputDescriptor bool, _ bool, desc ocispec.Descriptor, tty *os.File) (status.BlobPushHandler, metadata.BlobPushHandler) {
	_ = "STUB: not implemented"
	return *new(status.BlobPushHandler), *new(metadata.BlobPushHandler)
}

// NewResolveHandler returns a resolve metadata handler.
func NewResolveHandler(printer *output.Printer, fullRef bool, path string) metadata.ResolveHandler {
	_ = "STUB: not implemented"
	return *new(metadata.ResolveHandler)
}

// NewBlobDeleteHandler returns blob delete handlers.
func NewBlobDeleteHandler(printer *output.Printer, target *option.Target) metadata.BlobDeleteHandler {
	_ = "STUB: not implemented"
	return *new(metadata.BlobDeleteHandler)
}

// NewRepoTagsHandler returns a repo tags handler.
func NewRepoTagsHandler(out io.Writer, format option.Format) (metadata.RepoTagsHandler, error) {
	_ = "STUB: not implemented"
	return *new(metadata.RepoTagsHandler), nil
}

// NewRepoListHandler returns a repo ls handler.
func NewRepoListHandler(out io.Writer, format option.Format, registry, namespace string) (metadata.RepoListHandler, error) {
	_ = "STUB: not implemented"
	return *new(metadata.RepoListHandler), nil
}
