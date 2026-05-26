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
	"io"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/display/status"
	"oras.land/oras/cmd/oras/internal/option"
)

type pushBlobOptions struct {
	option.Common
	option.Descriptor
	option.Pretty
	option.Target
	option.Terminal

	fileRef   string
	mediaType string
	size      int64
	// Deprecated: verbose is deprecated and will be removed in the future.
	verbose bool
}

func pushCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func pushBlob(cmd *cobra.Command, opts *pushBlobOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// prepare blob content

// outputs blob's descriptor

func doPush(ctx context.Context, statusHandler status.BlobPushHandler, t oras.GraphTarget, desc ocispec.Descriptor, r io.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}
