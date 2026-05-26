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
	"context"
	"errors"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"oras.land/oras-go/v2/content"
)

// Pre-defined annotation keys for annotation file
const (
	AnnotationManifest = "$manifest"
	AnnotationConfig   = "$config"
)

var (
	errAnnotationConflict = errors.New("`--annotation` and `--annotation-file` cannot be both specified")
	errPathValidation     = errors.New("absolute file path detected. If it's intentional, use --disable-path-validation flag to skip this check")
)

// Packer option struct.
type Packer struct {
	Annotation

	ManifestExportPath     string
	PathValidationDisabled bool
	AnnotationFilePath     string

	FileRefs []string
}

// ApplyFlags applies flags to a command flag set.
func (opts *Packer) ApplyFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// ExportManifest saves the pushed manifest to a local file.
func (opts *Packer) ExportManifest(ctx context.Context, fetcher content.Fetcher, desc ocispec.Descriptor) error {
	_ = "STUB: not implemented"
	return nil
}

func (opts *Packer) Parse(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

// Remove the type if specified in the path <file>[:<type>] format

// parseAnnotations loads the manifest annotation map.
func (opts *Packer) parseAnnotations(cmd *cobra.Command) error {
	_ = "STUB: not implemented"
	return nil
}

// decodeJSON decodes file contents into json.
func decodeJSON(filename string, v any) (err error) { _ = "STUB: not implemented"; return nil }
