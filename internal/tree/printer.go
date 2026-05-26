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

package tree

import (
	"io"
	"os"
)

// Box-drawing symbols
const (
	EdgeEmpty = "    "
	EdgePipe  = "│   "
	EdgeItem  = "├── "
	EdgeLast  = "└── "
)

// DefaultPrinter prints the tree to the stdout with default settings.
var DefaultPrinter = NewPrinter(os.Stdout)

// Printer prints the tree.
type Printer struct {
	writer io.Writer
}

// NewPrinter create s a new printer.
func NewPrinter(writer io.Writer) *Printer { _ = "STUB: not implemented"; return nil }

// Print prints a tree.
func (p *Printer) Print(root *Node) error { _ = "STUB: not implemented"; return nil }

// print prints a tree recursively.
func (p *Printer) print(prefix string, n *Node) error { _ = "STUB: not implemented"; return nil }

// Print prints the tree using the default printer.
func Print(root *Node) error { _ = "STUB: not implemented"; return nil }
