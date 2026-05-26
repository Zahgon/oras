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
	"oras.land/oras/cmd/oras/internal/option"
)

type attachOptions struct {
	option.Common
	option.Packer
	option.Target
	option.Format
	option.Platform
	option.Terminal

	artifactType      string
	manifestConfigRef string
	concurrency       int
	// Deprecated: verbose is deprecated and will be removed in the future.
	verbose bool
}

func attachCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// no file argument provided

// invalid reference

// buildAttachPackOpts assembles a PackManifestOptions from the resolved
// attach options, subject descriptor, and layer descriptors.
func buildAttachPackOpts(opts *attachOptions, subject ocispec.Descriptor, descs []ocispec.Descriptor) oras.PackManifestOptions {
	_ = "STUB: not implemented"
	return *new(oras.PackManifestOptions)
}

func runAttach(cmd *cobra.Command, opts *attachOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// prepare manifest

// add both pull and push scope hints for dst repository
// to save potential push-scope token requests during copy

// prepare push

// skip duplicated Resolve on subject

// we don't need the CopyError information so we unwrap it here

// Attach

// Export manifest
