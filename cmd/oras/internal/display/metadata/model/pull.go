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

package model

import (
	"sync"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// File records metadata of a pulled file.
type File struct {
	// Path is the absolute path of the pulled file.
	Path string `json:"path"`
	Descriptor
}

// newFile creates a new file metadata.
func newFile(name string, outputDir string, desc ocispec.Descriptor, descPath string) (File, error) {
	_ = "STUB: not implemented"
	return *new(File), nil
}

// not likely to go wrong since the file has already be written to file store

type pull struct {
	DigestReference
	Files []File `json:"files"`
}

// NewPull creates a new metadata struct for pull command.
func NewPull(digestReference string, files []File) any { _ = "STUB: not implemented"; return *new(any) }

// Pulled records all pulled files.
type Pulled struct {
	lock  sync.Mutex
	files []File
}

// Files returns all pulled files.
func (p *Pulled) Files() []File { _ = "STUB: not implemented"; return nil }

// Add adds a pulled file.
func (p *Pulled) Add(name string, outputDir string, desc ocispec.Descriptor, descPath string) error {
	_ = "STUB: not implemented"
	return nil
}
