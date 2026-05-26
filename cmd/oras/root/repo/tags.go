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

package repo

import (
	"github.com/spf13/cobra"
	"oras.land/oras/cmd/oras/internal/option"
)

type showTagsOptions struct {
	option.Common
	option.Target
	option.Format

	last             string
	excludeDigestTag bool
}

func showTagsCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func showTags(cmd *cobra.Command, opts *showTagsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// if a repository path is given, filter the tags under the repository

// if a tag is given, show the associated tags

// if --oci-layout-path is used with a repository path, filter the
// tags under the repository.

// if --exclude-digest-tags is used, skip digest-like tags

// if a tag or digest is given, show the associated tags

// show the tags in the repository

func isDigestTag(tag string) bool { _ = "STUB: not implemented"; return false }
