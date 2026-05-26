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
	"context"
	"sync"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/display/status"
	"oras.land/oras/cmd/oras/internal/option"
)

type pullOptions struct {
	option.Cache
	option.Common
	option.Platform
	option.Target
	option.Format
	option.Terminal

	concurrency       int
	KeepOldFiles      bool
	IncludeSubject    bool
	PathTraversal     bool
	Output            string
	ManifestConfigRef string
	// Deprecated: verbose is deprecated and will be removed in the future.
	verbose bool
}

func pullCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runPull(cmd *cobra.Command, opts *pullOptions) (pullError error) {
	_ = "STUB: not implemented"
	return nil
}

// Copy Options

// customize friendly message for path traversal error

func doPull(ctx context.Context, src oras.ReadOnlyTarget, dst oras.GraphTarget, opts oras.CopyOptions, metadataHandler metadata.PullHandler, statusHandler status.PullHandler, po *pullOptions) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}

// empty layer

// unnamed layers are skipped

// skip s if it is unnamed AND has no successors.

// restore named but deduplicated successor nodes

// Copy

// we don't need the CopyError information so we unwrap it here

func notifyOnce(notified *sync.Map, s ocispec.Descriptor, notify func(ocispec.Descriptor) error) error {
	_ = "STUB: not implemented"
	return nil
}
