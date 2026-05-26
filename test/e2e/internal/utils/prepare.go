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

package utils

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// CopyZOTRepo copies oci layout data between repostories.
func CopyZOTRepo(fromRepo string, toRepo string) { _ = "STUB: not implemented"; return }

// PrepareTempOCI prepares an OCI layout root via copying from an ZOT repo and
// return the path.
func PrepareTempOCI(fromZotRepo string) string { _ = "STUB: not implemented"; return "" }

// PrepareTempFiles copies test data into a temp folder and return it.
func PrepareTempFiles() string { _ = "STUB: not implemented"; return "" }

// CopyTestFiles copies test data into dstRoot.
func CopyTestFiles(dstRoot string) error { _ = "STUB: not implemented"; return nil }

// CopyFiles copies files from folder src to folder dest.
func CopyFiles(src string, dest string) error { _ = "STUB: not implemented"; return nil }

// ignore folder

// make sure all parents are created

// copy with original folder structure

// MatchFile reads content from filepath, matches it with want with timeout.
func MatchFile(filepath string, want string, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

// WriteTempFile writes content into name under a temp folder.
func WriteTempFile(name string, content string) (path string) { _ = "STUB: not implemented"; return "" }

func copyFile(srcFile, dstFile string) error { _ = "STUB: not implemented"; return nil }
