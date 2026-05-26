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

package descriptor

import (
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// IsManifest checks if a descriptor describes a manifest.
// Adapted from `oras-go`: https://github.com/oras-project/oras-go/blob/d6c837e439f4c567f8003eab6e423c22900452a8/internal/descriptor/descriptor.go#L67
func IsManifest(desc ocispec.Descriptor) bool { _ = "STUB: not implemented"; return false }

// IsImageManifest checks whether a manifest is an image manifest.
func IsImageManifest(desc ocispec.Descriptor) bool { _ = "STUB: not implemented"; return false }

// IsIndex checks if a descriptor describes an image index or Docker manifest list.
func IsIndex(desc ocispec.Descriptor) bool { _ = "STUB: not implemented"; return false }

// ShortDigest converts the digest of the descriptor to a short form for displaying.
func ShortDigest(desc ocispec.Descriptor) (digestString string) {
	_ = "STUB: not implemented"
	return ""
}

// Plain returns a plain descriptor that contains only MediaType, Digest and Size.
// Copied from `oras-go`: https://github.com/oras-project/oras-go/blob/d6c837e439f4c567f8003eab6e423c22900452a8/internal/descriptor/descriptor.go#L81
func Plain(desc ocispec.Descriptor) ocispec.Descriptor {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor)
}

// GetTitleOrMediaType gets a descriptor name using either title or media type.
func GetTitleOrMediaType(desc ocispec.Descriptor) (name string, isTitle bool) {
	_ = "STUB: not implemented"
	return "", false
}

// GenerateContentKey generates a unique key for each content descriptor using
// digest and name.
func GenerateContentKey(desc ocispec.Descriptor) string { _ = "STUB: not implemented"; return "" }
