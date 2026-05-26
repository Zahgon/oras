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
	"github.com/spf13/cobra"
	"oras.land/oras/cmd/oras/internal/option"
)

type deleteBlobOptions struct {
	option.Common
	option.Confirmation
	option.Descriptor
	option.Pretty
	option.Target
}

func deleteCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func deleteBlob(cmd *cobra.Command, opts *deleteBlobOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// add both pull and delete scope hints for dst repository to save potential delete-scope token requests during deleting
