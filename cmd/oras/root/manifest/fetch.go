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
	"github.com/spf13/cobra"
	"oras.land/oras/cmd/oras/internal/option"
)

type fetchOptions struct {
	option.Cache
	option.Common
	option.Descriptor
	option.Platform
	option.Pretty
	option.Target
	option.Format

	mediaTypes []string
	outputPath string
}

func fetchCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func fetchManifest(cmd *cobra.Command, opts *fetchOptions) (fetchErr error) {
	_ = "STUB: not implemented"
	return nil
}

// fetch manifest descriptor only

// fetch manifest descriptor and content
