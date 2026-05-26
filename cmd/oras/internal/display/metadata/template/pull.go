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

package template

import (
	"io"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/display/metadata/model"
	"oras.land/oras/cmd/oras/internal/option"
)

// PullHandler handles text metadata output for pull events.
type PullHandler struct {
	template string
	path     string
	out      io.Writer
	pulled   model.Pulled
	root     ocispec.Descriptor
}

// NewPullHandler returns a new handler for pull events.
func NewPullHandler(out io.Writer, path string, template string) metadata.PullHandler {
	_ = "STUB: not implemented"
	return *new(metadata.PullHandler)
}

// OnPulled implements metadata.PullHandler.
func (ph *PullHandler) OnPulled(_ *option.Target, desc ocispec.Descriptor) {
	_ = "STUB: not implemented"

	// Render implements metadata.PullHandler.
	return
}

func (ph *PullHandler) Render() error { _ = "STUB: not implemented"; return nil }

// OnFilePulled implements metadata.PullHandler.
func (ph *PullHandler) OnFilePulled(name string, outputDir string, desc ocispec.Descriptor, descPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// OnLayerSkipped implements metadata.PullHandler.
func (ph *PullHandler) OnLayerSkipped(ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}
