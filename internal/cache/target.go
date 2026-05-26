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

package cache

import (
	"context"
	"io"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content"
	"oras.land/oras-go/v2/registry"
)

type closer func() error

func (fn closer) Close() error {
	_ = "STUB: not implemented"

	// Cache target struct.
	return nil
}

type target struct {
	oras.ReadOnlyTarget
	cache content.Storage
}

// New generates a new target storage with caching.
func New(source oras.ReadOnlyTarget, cache content.Storage) oras.ReadOnlyTarget {
	_ = "STUB: not implemented"
	return *new(oras.ReadOnlyTarget)
}

// Fetch fetches the content identified by the descriptor.
func (t *target) Fetch(ctx context.Context, target ocispec.Descriptor) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Fetch from cache

// Fetch from origin with caching

func (t *target) cacheReadCloser(ctx context.Context, rc io.ReadCloser, target ocispec.Descriptor) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

// Exists returns true if the described content exists.
func (t *target) Exists(ctx context.Context, desc ocispec.Descriptor) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Cache referenceTarget struct.
type referenceTarget struct {
	*target
	registry.ReferenceFetcher
}

// FetchReference fetches the content identified by the reference from the
// remote and cache the fetched content.
// Cached content will only be read via Fetch, FetchReference will always fetch
// From origin.
func (t *referenceTarget) FetchReference(ctx context.Context, reference string) (ocispec.Descriptor, io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), *new(io.ReadCloser), nil
}

// skip caching if the content already exists in cache

// get rc from the cache

// no need to do tee'd push

// Fetch from origin with caching
