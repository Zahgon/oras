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
	"io"

	"oras.land/oras/cmd/oras/internal/display/metadata"
)

// repoTagsHandler handles text output for repo tags command.
type repoTagsHandler struct {
	out io.Writer
}

// NewRepoTagsHandler creates a new text handler for repo tags command.
func NewRepoTagsHandler(out io.Writer) metadata.RepoTagsHandler {
	_ = "STUB: not implemented"
	return *new(metadata.RepoTagsHandler)
}

// OnTagListed implements metadata.TagsHandler.
func (h *repoTagsHandler) OnTagListed(tag string) error { _ = "STUB: not implemented"; return nil }

// Render implements metadata.TagsHandler.
func (h *repoTagsHandler) Render() error { _ = "STUB: not implemented"; return nil }
