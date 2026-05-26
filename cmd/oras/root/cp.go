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

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/display/status"
	"oras.land/oras/cmd/oras/internal/option"
)

type copyOptions struct {
	option.Common
	option.Platform
	option.BinaryTarget
	option.Terminal

	recursive   bool
	concurrency int
	extraRefs   []string
	// Deprecated: verbose is deprecated and will be removed in the future.
	verbose bool
}

func copyCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runCopy(cmd *cobra.Command, opts *copyOptions) error { _ = "STUB: not implemented"; return nil }

// Prepare source

// Prepare destination

// correct source digest

func doCopy(ctx context.Context, copyHandler status.CopyHandler, src oras.ReadOnlyGraphTarget, dst oras.GraphTarget, opts *copyOptions) (desc ocispec.Descriptor, err error) {
	_ = "STUB: not implemented"
	// Prepare copy options
	return *new(ocispec.Descriptor), nil
}

// leave the CopyError to oerrors.Modifier for prefix processing

// recursiveCopy copies an artifact and its referrers from one target to another.
// If the artifact is a manifest list or index, referrers of its manifests are copied as well.
func recursiveCopy(ctx context.Context, src oras.ReadOnlyGraphTarget, dst oras.Target, dstRef string, root ocispec.Descriptor, opts oras.ExtendedCopyGraphOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareCopyOption(ctx context.Context, src oras.ReadOnlyGraphTarget, _ oras.Target, root ocispec.Descriptor, opts oras.ExtendedCopyGraphOptions) (oras.ExtendedCopyGraphOptions, ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(oras.ExtendedCopyGraphOptions), *new(ocispec.Descriptor), nil
}

// no child manifests, thus no child referrers

// no child referrers

// If root has no referrers, we set copyRoot, which is the entry point of
// extended copy, to the first manifest in the index. We also put the root
// and the referrers of the manifests as the predecessors of copyRoot. This
// is to ensure that all these nodes can be copied by calling extended copy.
// Reference: https://github.com/oras-project/oras/issues/1728

// getMountPoint checks if mounting can be performed between two targets and returns
// the repository name to be mounted from if applicable. Mount can be performed if the two
// targets are both remote repositories, are in the same registry and have identical credentials.
func getMountPoint(src oras.ReadOnlyGraphTarget, dst oras.GraphTarget, opts *copyOptions) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
