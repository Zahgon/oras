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
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/option"
	"oras.land/oras/cmd/oras/internal/output"
)

// ManifestDeleteHandler handles text metadata output for manifest delete events.
type ManifestDeleteHandler struct {
	printer *output.Printer
	target  *option.Target
}

// NewManifestDeleteHandler returns a new handler for manifest delete events.
func NewManifestDeleteHandler(printer *output.Printer, target *option.Target) metadata.ManifestDeleteHandler {
	_ = "STUB: not implemented"
	return *new(metadata.ManifestDeleteHandler)
}

// OnManifestMissing implements ManifestDeleteHandler.
func (h *ManifestDeleteHandler) OnManifestMissing() error { _ = "STUB: not implemented"; return nil }

// OnManifestDeleted implements ManifestDeleteHandler.
func (h *ManifestDeleteHandler) OnManifestDeleted() error { _ = "STUB: not implemented"; return nil }
