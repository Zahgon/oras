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

package option

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// BinaryTarget struct contains flags and arguments specifying two registries or
// image layouts.
// BinaryTarget implements errors.Handler interface.
type BinaryTarget struct {
	From        Target
	To          Target
	resolveFlag []string
}

// EnsureSourceTargetReferenceNotEmpty ensures that from target reference is not empty.
func (target *BinaryTarget) EnsureSourceTargetReferenceNotEmpty(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

// EnableDistributionSpecFlag set distribution specification flag as applicable.
func (target *BinaryTarget) EnableDistributionSpecFlag() { _ = "STUB: not implemented"; return }

// ApplyFlags applies flags to a command flag set fs.
func (target *BinaryTarget) ApplyFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Parse parses user-provided flags and arguments into option struct.
func (target *BinaryTarget) Parse(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

// resolve are parsed in array order, latter will overwrite former

// ModifyError handles error during cmd execution.
func (target *BinaryTarget) ModifyError(cmd *cobra.Command, err error) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// extract the inner error

// Example: Error from source registry for "localhost:5000/test:v1":
// Example: Error from destination oci-layout for "oci-dir:v1":

func (target *BinaryTarget) modifyError(cmd *cobra.Command, err error) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
