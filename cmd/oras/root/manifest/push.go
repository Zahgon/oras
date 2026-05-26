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

package manifest

import (
	"context"

	digest "github.com/opencontainers/go-digest"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2/content"
	"oras.land/oras/cmd/oras/internal/option"
)

type pushOptions struct {
	option.Common
	option.Descriptor
	option.Pretty
	option.Target

	concurrency int
	extraRefs   []string
	fileRef     string
	mediaType   string
	// Deprecated: verbose is deprecated and will be removed in the future.
	verbose bool
}

func pushCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func pushManifest(cmd *cobra.Command, opts pushOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// prepare manifest content

// get manifest media type

// prepare manifest descriptor

// outputs manifest's descriptor

// matchDigest checks whether the manifest's digest matches to it in the remote
// repository.
func matchDigest(ctx context.Context, resolver content.Resolver, reference string, digest digest.Digest) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
