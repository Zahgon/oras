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

package contentutil

import (
	"context"
	"io"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2"
)

type multiReadOnlyTarget struct {
	targets []oras.ReadOnlyTarget
}

// MultiReadOnlyTarget returns a ReadOnlyTarget that combines multiple targets.
func MultiReadOnlyTarget(targets ...oras.ReadOnlyTarget) oras.ReadOnlyTarget {
	_ = "STUB: not implemented"
	return *new(oras.ReadOnlyTarget)
}

// Fetch fetches the content from the targets in order and return first found
// content. If no content is found, it returns ErrNotFound.
func (m *multiReadOnlyTarget) Fetch(ctx context.Context, target ocispec.Descriptor) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Exists returns true if the content exists in any of the targets.
// multiReadOnlyTarget does not implement Exists() because it's read-only.
func (m *multiReadOnlyTarget) Exists(_ context.Context, _ ocispec.Descriptor) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Resolve resolves the reference to a descriptor from the targets in order and
// return first found descriptor. If no descriptor is found, it returns
// ErrNotFound.
func (m *multiReadOnlyTarget) Resolve(ctx context.Context, ref string) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}
