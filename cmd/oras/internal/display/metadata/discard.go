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

package metadata

import (
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/cmd/oras/internal/option"
)

type Discard struct{}

// NewDiscardHandler creates a new handler that discards output for all events.
func NewDiscardHandler() Discard {
	_ = "STUB: not implemented"

	// OnFetched implements ManifestFetchHandler.
	return *new(Discard)
}

func (Discard) OnFetched(string, ocispec.Descriptor, []byte) error {
	_ = "STUB: not implemented"

	// OnManifestPushed implements ManifestPushHandler.
	return nil
}

func (Discard) OnManifestPushed(ocispec.Descriptor) error {
	_ = "STUB: not implemented"

	// Render implements ManifestPushHandler.
	return nil
}

func (Discard) Render() error {
	_ = "STUB: not implemented"

	// OnTagged implements ManifestIndexCreateHandler.
	return nil
}

func (Discard) OnTagged(ocispec.Descriptor, string) error {
	_ = "STUB: not implemented"

	// OnIndexCreated implements ManifestIndexCreateHandler.
	return nil
}

func (Discard) OnIndexCreated(ocispec.Descriptor) {
	_ = "STUB: not implemented"

	// OnBlobPushed implements BlobPushHandler
	return
}

func (Discard) OnBlobPushed(*option.Target) error { _ = "STUB: not implemented"; return nil }
