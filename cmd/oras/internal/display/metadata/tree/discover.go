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

package tree

import (
	"io"
	"os"

	"github.com/morikuni/aec"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/internal/tree"
)

var (
	artifactTypeColor = aec.LightYellowF
	digestColor       = aec.GreenF
	annotationsColor  = aec.LightBlackF
)

// discoverHandler handles json metadata output for discover events.
type discoverHandler struct {
	out     io.Writer
	path    string
	root    *tree.Node
	nodes   map[digest.Digest]*tree.Node
	verbose bool
	tty     *os.File
}

// NewDiscoverHandler creates a new handler for discover events.
func NewDiscoverHandler(out io.Writer, path string, root ocispec.Descriptor, verbose bool, tty *os.File) metadata.DiscoverHandler {
	_ = "STUB: not implemented"
	return *new(metadata.DiscoverHandler)
}

// OnDiscovered implements metadata.DiscoverHandler.
func (h *discoverHandler) OnDiscovered(referrer, subject ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

// add artifact type and digest to the referrer

// add annotations to the referrer

// Render implements metadata.DiscoverHandler.
func (h *discoverHandler) Render() error { _ = "STUB: not implemented"; return nil }
