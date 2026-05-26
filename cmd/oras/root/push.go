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

package root

import (
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/display/status"
	"oras.land/oras/cmd/oras/internal/option"
)

type pushOptions struct {
	option.Common
	option.Packer
	option.ArtifactPlatform
	option.ImageSpec
	option.Target
	option.Format
	option.Terminal

	extraRefs         []string
	manifestConfigRef string
	artifactType      string
	concurrency       int
	// Deprecated: verbose is deprecated and will be removed in the future.
	verbose bool
}

func pushCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// switch to v1.0 manifest since artifact type is suggested
// by OCI v1.1 artifact guidance but is not presented
// see https://github.com/opencontainers/image-spec/blob/e7f7c0ca69b21688c3cea7c87a04e4503e6099e2/manifest.md?plain=1#L170

func runPush(cmd *cobra.Command, opts *pushOptions) error { _ = "STUB: not implemented"; return nil }

// prepare pack

// prepare push

// add both pull and push scope hints for dst repository
// to save potential push-scope token requests during copy

// we don't need the CopyError information so we unwrap it here

// Push

// Export manifest

func doPush(dst oras.Target, stopTrack status.StopTrackTargetFunc, pack packFunc, copyFunc copyFunc) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}

// Push

type packFunc func() (ocispec.Descriptor, error)
type copyFunc func(desc ocispec.Descriptor) error

func pushArtifact(_ oras.Target, pack packFunc, copyFunc copyFunc) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}

// push
