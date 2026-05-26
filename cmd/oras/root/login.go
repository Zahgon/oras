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
	"io"

	"github.com/spf13/cobra"
	"oras.land/oras/cmd/oras/internal/option"
)

type loginOptions struct {
	option.Common
	option.Remote
	Hostname string
}

func loginCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

func runLogin(cmd *cobra.Command, opts loginOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// prompt for credential

// prompt for username

// prompt for token

// prompt for password

func readLine(outWriter io.Writer, prompt string, silent bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
