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

type tagOptions struct {
	option.Common
	option.Target

	concurrency int
	targetRefs  []string
}

func tagCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func tagManifest(cmd *cobra.Command, opts *tagOptions) error { _ = "STUB: not implemented"; return nil }

// Since referrer capability has not been set or detected yet,
// nil is the only returned value and thus can be ignored
