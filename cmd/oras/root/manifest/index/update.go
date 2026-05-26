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

package index

import (
	"context"

	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/display/status"
	"oras.land/oras/cmd/oras/internal/option"
)

type updateOptions struct {
	option.Common
	option.Target
	option.Pretty

	artifactType    string
	addArguments    []string
	mergeArguments  []string
	removeArguments []string
	tags            []string
	outputPath      string
}

func updateCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func updateIndex(cmd *cobra.Command, opts updateOptions) error {
	_ = "STUB: not implemented"
	// if no update flag is used, do nothing
	return nil
}

func fetchIndex(ctx context.Context, handler status.ManifestIndexUpdateHandler, target oras.ReadOnlyTarget, reference string) (ocispec.Index, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Index), nil
}

func addManifests(ctx context.Context, displayStatus status.ManifestIndexUpdateHandler, manifests []ocispec.Descriptor, target oras.ReadOnlyTarget, addArguments []string) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeIndexes(ctx context.Context, displayStatus status.ManifestIndexUpdateHandler, manifests []ocispec.Descriptor, target oras.ReadOnlyTarget, mergeArguments []string) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeManifests(handler status.ManifestIndexUpdateHandler, manifests []ocispec.Descriptor, _ oras.ReadOnlyTarget, opts updateOptions) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	// create a set of digests to speed up the remove
	return nil, nil
}

func doRemoveManifests(originalManifests []ocispec.Descriptor, digestToRemove map[digest.Digest]bool, handler status.ManifestIndexUpdateHandler, indexRef string) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func updateFlagsUsed(flags *pflag.FlagSet) bool { _ = "STUB: not implemented"; return false }

func getPushPath(rawReference string, targetType string, reference string, path string) string {
	_ = "STUB: not implemented"
	return ""
}
