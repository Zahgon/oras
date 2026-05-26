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

// FormatType represents a format type.
type FormatType struct {
	// Name is the format type name.
	Name string
	// Usage is the usage string in help doc.
	Usage string
	// HasParams indicates whether the format type has parameters.
	HasParams bool
}

// WithUsage returns a new format type with provided usage string.
func (ft *FormatType) WithUsage(usage string) *FormatType { _ = "STUB: not implemented"; return nil }

// format types
var (
	FormatTypeJSON = &FormatType{
		Name:  "json",
		Usage: "Print in JSON format",
	}
	FormatTypeGoTemplate = &FormatType{
		Name:      "go-template",
		Usage:     "Print output using the given Go template",
		HasParams: true,
	}
	// the table format is deprecated
	FormatTypeTable = &FormatType{
		Name:  "table",
		Usage: "[Deprecated] Get referrers and output in table format",
	}
	FormatTypeTree = &FormatType{
		Name:  "tree",
		Usage: "Get referrers and print in tree format",
	}
	FormatTypeText = &FormatType{
		Name:  "text",
		Usage: "Print in text format",
	}
)

// Format contains input and parsed options for formatted output flags.
type Format struct {
	FormatFlag   string
	Type         string
	Template     string
	allowedTypes []*FormatType
}

// SetTypes sets the default format type and allowed format types.
func (f *Format) SetTypes(defaultType *FormatType, otherTypes ...*FormatType) {
	_ = "STUB: not implemented"
	return
}

// ApplyFlags implements FlagProvider.ApplyFlag.
func (f *Format) ApplyFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// apply flags

// Parse parses the input format flag.
func (f *Format) Parse(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	// print deprecation message for table format
	return nil
}

// flag not specified

// type validation passed

func (f *Format) parseFlag() error { _ = "STUB: not implemented"; return nil }

// template explicitly set

// parse type and add parameter to template
