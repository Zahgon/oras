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

package blob

import (
	"context"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/option"
)

type fetchBlobOptions struct {
	option.Cache
	option.Common
	option.Descriptor
	option.Pretty
	option.Target
	option.Terminal

	outputPath string
}

func fetchCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func fetchBlob(cmd *cobra.Command, opts *fetchBlobOptions) (fetchErr error) {
	_ = "STUB: not implemented"
	return nil
}

// outputs blob's descriptor if `--descriptor` is used

func (opts *fetchBlobOptions) doFetch(ctx context.Context, src oras.ReadOnlyTarget) (desc ocispec.Descriptor, fetchErr error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}

// fetch blob descriptor only

// fetch blob content

// outputs blob content if "--output -" is used

// save blob content into the local file if the output path is provided

// none TTY output

// TTY output
