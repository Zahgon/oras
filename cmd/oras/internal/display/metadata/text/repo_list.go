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

// repoListHandler handles text output for repo ls command.
type repoListHandler struct {
	out       io.Writer
	namespace string
}

// NewRepoListHandler creates a new text handler for repo ls command.
func NewRepoListHandler(out io.Writer, namespace string) metadata.RepoListHandler {
	_ = "STUB: not implemented"
	return *new(metadata.RepoListHandler)
}

// OnRepositoryListed implements metadata.RepoListHandler.
func (h *repoListHandler) OnRepositoryListed(repo string) error {
	_ = "STUB: not implemented"
	// For text format, show only the sub repo (without the namespace prefix) for better readability
	return nil
}

// Render implements metadata.RepoListHandler.
func (h *repoListHandler) Render() error { _ = "STUB: not implemented"; return nil }
