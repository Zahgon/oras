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

package io

import (
	"io"
)

// TarDirectory creates a tar archive from the contents of sourceDir and writes it to the given writer.
func TarDirectory(writer io.Writer, sourceDir string) (tarErr error) {
	_ = "STUB: not implemented"
	// Ensure sourceDir exists and is a directory
	return nil
}

// Create a new tar writer

// IsTarFile loosely checks whether the given file path refers to a tar archive
// by examining its extension and magic number.
func IsTarFile(path string) (bool, error) {
	_ = "STUB: not implemented"
	// loose check: consider *.tar files as tar archives
	return false, nil
}

// check the magic number to determine the file type

// read 5 bytes ("ustar") at the position where the magic number is located

// ustar magic number starts at byte 257
