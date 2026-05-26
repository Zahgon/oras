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

package file

import (
	"io"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// PrepareManifestContent prepares the content for manifest from the file path
// or stdin.
func PrepareManifestContent(path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PrepareBlobContent prepares the content descriptor for blob from the file
// path or stdin. Use the input digest and size if they are provided. Will
// return error if the content is from stdin but the content digest and size
// are missing.
func PrepareBlobContent(path string, mediaType string, digestString string, size int64) (desc ocispec.Descriptor, rc io.ReadCloser, err error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), *new(io.ReadCloser), nil
}

// validate digest

// prepares the content descriptor from stdin

// throw err if size or digest is not provided.
