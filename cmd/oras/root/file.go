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

package root

import (
	"context"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras/cmd/oras/internal/display/status"
)

func loadFiles(ctx context.Context, store *file.Store, annotations map[string]map[string]string, fileRefs []string, displayStatus status.PushHandler) ([]ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get shortest absolute path as unique name

func addFile(ctx context.Context, store *file.Store, name string, mediaType string, filename string) (ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return *new(ocispec.Descriptor), nil
}
