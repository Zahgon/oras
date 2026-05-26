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

package output

import (
	"io"
	"sync"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// Printer prints for status handlers.
type Printer struct {
	Verbose bool

	out  io.Writer
	err  io.Writer
	lock sync.Mutex
}

// NewPrinter creates a new Printer.
func NewPrinter(out io.Writer, err io.Writer) *Printer { _ = "STUB: not implemented"; return nil }

// Write implements the io.Writer interface.
func (p *Printer) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Println prints objects concurrent-safely with newline.
func (p *Printer) Println(a ...any) error { _ = "STUB: not implemented"; return nil }

// Printf prints objects concurrent-safely with newline.
func (p *Printer) Printf(format string, a ...any) error { _ = "STUB: not implemented"; return nil }

// Errors are handled above, so return nil

// PrintVerbose prints when verbose is true.
func (p *Printer) PrintVerbose(a ...any) error { _ = "STUB: not implemented"; return nil }

// PrintStatus prints transfer status.
func (p *Printer) PrintStatus(desc ocispec.Descriptor, status string) error {
	_ = "STUB: not implemented"
	return nil
}
