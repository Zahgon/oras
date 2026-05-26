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

package json

import (
	"io"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/display/metadata/model"
	"oras.land/oras/cmd/oras/internal/option"
)

// PushHandler handles JSON metadata output for push events.
type PushHandler struct {
	path   string
	out    io.Writer
	tagged model.Tagged
	root   ocispec.Descriptor
}

// NewPushHandler creates a new handler for push events.
func NewPushHandler(out io.Writer) metadata.PushHandler {
	_ = "STUB: not implemented"
	return *new(metadata.PushHandler)
}

// OnTagged implements metadata.TaggedHandler.
func (ph *PushHandler) OnTagged(_ ocispec.Descriptor, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnCopied is called after files are copied.
func (ph *PushHandler) OnCopied(opts *option.Target, root ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// Render implements PushHandler.
func (ph *PushHandler) Render() error { _ = "STUB: not implemented"; return nil }
