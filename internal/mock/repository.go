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

// Package mock contains mocking components for unit testing.
package mock

import (
	"context"
	"errors"
	"io"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2/registry/remote"
)

type content struct {
	ocispec.Descriptor
	blob []byte
}

type Repository struct {
	cas                map[string]content
	remote.Repository  // make tests compile
	isFetcher          bool
	isReferenceFetcher bool
	isResolver         bool
}

// WithFetch enables mocking for Fetch.
func (repo *Repository) WithFetch() *Repository { _ = "STUB: not implemented"; return nil }

// WithFetchReference enables mocking for FetchReference.
func (repo *Repository) WithFetchReference() *Repository { _ = "STUB: not implemented"; return nil }

// WithResolve enables mocking for Resolve.
func (repo *Repository) WithResolve() *Repository { _ = "STUB: not implemented"; return nil }

// New returns a new Repository struct.
func New() *Repository { _ = "STUB: not implemented"; return nil }

// Blob mocks a content blob stored in content-addressable storage.
type Blob struct {
	Content   string
	MediaType string
	Tag       string
}

// Remount remounts the underlying CAS of the Repository.
func (repo *Repository) Remount(blobs []Blob) { _ = "STUB: not implemented"; return }

var errNotImplemented = errors.New("not implemented")

// FetchReference mocks the fetching via a reference ref.
func (repo *Repository) FetchReference(_ context.Context, ref string) (ocispec.Descriptor, io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), *new(io.ReadCloser), nil
}

// Fetch mocks fetching the target descriptor.
func (repo *Repository) Fetch(_ context.Context, target ocispec.Descriptor) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Resolve mocks resolving via a reference.
func (repo *Repository) Resolve(_ context.Context, reference string) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}
