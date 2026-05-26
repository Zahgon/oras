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

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/option"
)

type fetchConfigOptions struct {
	option.Cache
	option.Common
	option.Descriptor
	option.Platform
	option.Pretty
	option.Target

	outputPath string
}

func fetchConfigCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func fetchConfig(cmd *cobra.Command, opts *fetchConfigOptions) (fetchErr error) {
	_ = "STUB: not implemented"
	return nil
}

// fetch config descriptor

// fetch config content

// output config content

// save config into the local file if the output path is provided

// output config's descriptor

func fetchConfigDesc(ctx context.Context, src oras.ReadOnlyTarget, reference string, targetPlatform *ocispec.Platform) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	// fetch manifest descriptor and content
	return *new(ocispec.Descriptor), nil
}

// unmarshal manifest content to extract config descriptor
