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

package graph

import (
	"context"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content"
)

// MediaTypeArtifactManifest specifies the media type for a content descriptor.
const MediaTypeArtifactManifest = "application/vnd.oci.artifact.manifest.v1+json"

// Artifact describes an artifact manifest.
// This structure provides `application/vnd.oci.artifact.manifest.v1+json` mediatype when marshalled to JSON.
//
// This manifest type was introduced in image-spec v1.1.0-rc1 and was removed in
// image-spec v1.1.0-rc3. It is not part of the current image-spec and is kept
// here for Go compatibility.
//
// Reference: https://github.com/opencontainers/image-spec/pull/999
type Artifact struct {
	// MediaType is the media type of the object this schema refers to.
	MediaType string `json:"mediaType"`

	// ArtifactType is the IANA media type of the artifact this schema refers to.
	ArtifactType string `json:"artifactType"`

	// Blobs is a collection of blobs referenced by this manifest.
	Blobs []ocispec.Descriptor `json:"blobs,omitempty"`

	// Subject (reference) is an optional link from the artifact to another manifest forming an association between the artifact and the other manifest.
	Subject *ocispec.Descriptor `json:"subject,omitempty"`

	// Annotations contains arbitrary metadata for the artifact manifest.
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Successors returns the nodes directly pointed by the current node, picking
// out subject and config descriptor if applicable.
// Returning nil when no subject and config found.
func Successors(ctx context.Context, fetcher content.Fetcher, node ocispec.Descriptor) (nodes []ocispec.Descriptor, subject, config *ocispec.Descriptor, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// FindPredecessors returns all predecessors of descs in src concurrently.
func FindPredecessors(ctx context.Context, src oras.ReadOnlyGraphTarget, descs []ocispec.Descriptor, opts oras.ExtendedCopyGraphOptions) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RecursiveFindReferrers finds all referrers of the given descriptors recursively.
func RecursiveFindReferrers(ctx context.Context, src oras.ReadOnlyGraphTarget, descs []ocispec.Descriptor, opts oras.ExtendedCopyGraphOptions) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FilteredSuccessors fetches successors and returns filtered ones.
func FilteredSuccessors(ctx context.Context, desc ocispec.Descriptor, fetcher content.Fetcher, filter func(ocispec.Descriptor) bool) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
