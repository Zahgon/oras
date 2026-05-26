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
	"sync/atomic"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/option"
	"oras.land/oras/cmd/oras/internal/output"
)

// PullHandler handles text metadata output for pull events.
type PullHandler struct {
	printer      *output.Printer
	layerSkipped atomic.Bool
	target       *option.Target
	root         ocispec.Descriptor
}

// NewPullHandler returns a new handler for Pull events.
func NewPullHandler(printer *output.Printer) metadata.PullHandler {
	_ = "STUB: not implemented"
	return *new(metadata.PullHandler)
}

func (ph *PullHandler) OnFilePulled(_ string, _ string, _ ocispec.Descriptor, _ string) error {
	_ = "STUB: not implemented"

	// OnLayerSkipped implements metadata.PullHandler.
	return nil
}

func (ph *PullHandler) OnLayerSkipped(ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// OnPulled implements metadata.PullHandler.
func (ph *PullHandler) OnPulled(target *option.Target, desc ocispec.Descriptor) {
	_ = "STUB: not implemented"
	return
}

// Render implements metadata.PullHandler.
func (ph *PullHandler) Render() error { _ = "STUB: not implemented"; return nil }
