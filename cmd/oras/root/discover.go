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
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/option"
)

type discoverOptions struct {
	option.Common
	option.Platform
	option.Target
	option.Format
	option.Terminal

	artifactType string
	depth        int
	// Deprecated: verbose is deprecated and will be removed in the future.
	verbose bool
}

func discoverCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// only show direct referrers for table format

func runDiscover(cmd *cobra.Command, opts *discoverOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// discover artifacts

func fetchAllReferrers(ctx context.Context, repo oras.ReadOnlyGraphTarget, desc ocispec.Descriptor, artifactType string, handler metadata.DiscoverHandler, depth int) error {
	_ = "STUB: not implemented"
	return nil
}
