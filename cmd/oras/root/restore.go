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
	"github.com/spf13/cobra"
	"oras.land/oras/cmd/oras/internal/option"
)

type restoreOptions struct {
	option.Common
	option.Remote
	option.Terminal

	// flags
	input            string
	excludeReferrers bool
	dryRun           bool
	concurrency      int

	// derived options
	repository string
	tags       []string
}

func restoreCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// parse repo and tags

// always print verbose output

// required flag

// optional flags

// apply flags

func runRestore(cmd *cobra.Command, opts *restoreOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// start timing the restore process

// prepare the target registry

// prepare the source OCI store

// resolve tags to restore

// prepare copy options

// count referrers from source

// dry run, skip actual copy
