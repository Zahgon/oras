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

package testutils

import (
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2/content"
	"oras.land/oras-go/v2/content/memory"
)

// MockFetcher implements content.Fetcher and populates a memory store.
type MockFetcher struct {
	store       *memory.Store
	Fetcher     content.Fetcher
	Subject     ocispec.Descriptor
	Config      ocispec.Descriptor
	OciImage    ocispec.Descriptor
	ImageLayer  ocispec.Descriptor
	DockerImage ocispec.Descriptor
	Index       ocispec.Descriptor
}

// NewMockFetcher creates a MockFetcher and populates it.
func NewMockFetcher() (mockFetcher MockFetcher) {
	_ = "STUB: not implemented"
	return *new(MockFetcher)
}

// PushBlob pushes a blob to the memory store.
func (mf *MockFetcher) PushBlob(mediaType string, blob []byte) ocispec.Descriptor {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor)
}

func (mf *MockFetcher) pushImage(subject *ocispec.Descriptor, mediaType string, config ocispec.Descriptor, layers ...ocispec.Descriptor) ocispec.Descriptor {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor)
}

// PushOCIImage pushes the given subject, config and layers as a OCI image.
func (mf *MockFetcher) PushOCIImage(subject *ocispec.Descriptor, config ocispec.Descriptor, layers ...ocispec.Descriptor) ocispec.Descriptor {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor)
}

// PushDockerImage pushes the given subject, config and layers as a Docker image.
func (mf *MockFetcher) PushDockerImage(config ocispec.Descriptor, layers ...ocispec.Descriptor) ocispec.Descriptor {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor)
}

// PushIndex pushes the manifests as an index.
func (mf *MockFetcher) PushIndex(manifests ...ocispec.Descriptor) ocispec.Descriptor {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor)
}
