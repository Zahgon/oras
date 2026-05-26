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
	"regexp"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/display/status"
	"oras.land/oras/cmd/oras/internal/option"
)

var maxConfigSize int64 = 4 * 1024 * 1024 // 4 MiB

// mediaTypeRegexp is the regular expression pattern required for a valid
// media type, as defined in the image spec schema:
// - https://github.com/opencontainers/image-spec/blob/v1.1.1/schema/defs-descriptor.json#L7
var mediaTypeRegexp = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9!#$&^_.+-]{0,126}/[A-Za-z0-9][A-Za-z0-9!#$&^_.+-]{0,126}$`)

type createOptions struct {
	option.Common
	option.Target
	option.Pretty
	option.Annotation

	artifactType string
	sources      []string
	extraRefs    []string
	outputPath   string
}

func createCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func createIndex(cmd *cobra.Command, opts createOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func fetchSourceManifests(ctx context.Context, displayStatus status.ManifestIndexCreateHandler, target oras.ReadOnlyTarget, sources []string) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPlatform(ctx context.Context, target oras.ReadOnlyTarget, manifest *ocispec.Manifest) (*ocispec.Platform, error) {
	_ = "STUB: not implemented"
	// if config size is larger than 4 MiB, discontinue the fetch
	return nil, nil
}

// fetch config content

// ignore JSON unmarshal errors if the manifest does not have platform information
//nolint:nilerr,nilnil

// ignore if the manifest does not have platform information
//nolint:nilnil

func pushIndex(ctx context.Context, displayStatus status.ManifestIndexCreateHandler, taggedHandler metadata.TaggedHandler,
	target oras.Target, desc ocispec.Descriptor, content []byte, ref string, extraRefs []string, path string) error {
	_ = "STUB: not implemented"
	// push the index
	return nil
}

func enrichDescriptor(ctx context.Context, target oras.ReadOnlyTarget, desc ocispec.Descriptor, manifestBytes []byte) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}

// validateMediaType checks whether mediaType uses valid media type syntax,
// returning a non-nil error if not.
func validateMediaType(mediaType string) error { _ = "STUB: not implemented"; return nil }
