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

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/oci"
	"oras.land/oras-go/v2/registry/remote"
)

const (
	TargetTypeRemote    = "registry"
	TargetTypeOCILayout = "oci-layout"
)

// Target struct contains flags and arguments specifying one registry or image
// layout.
// Target implements oerrors.Handler interface.
type Target struct {
	Remote
	RawReference string
	Type         string
	Reference    string //contains tag or digest
	// Path contains
	//  - path to the OCI image layout target, or
	//  - registry and repository for the remote target
	Path string

	IsOCILayout bool

	prefix      string
	description string
}

// GetDisplayReference returns full printable reference.
func (target *Target) GetDisplayReference() string { _ = "STUB: not implemented"; return "" }

// setFlagDetails set directional flag prefix and description details
func (target *Target) setFlagDetails(prefix, description string) { _ = "STUB: not implemented"; return }

// ApplyFlags applies flags to a command flag set
// The complete form of the `target` flag is designed to be
//
//	--target type=<type>[[,<key>=<value>][...]]
//
// For better UX, the boolean flag `--oci-layout` is introduced as an alias of
// `--target type=oci-layout`.
// Since there is only one target type besides the default `registry` type,
// the full form is not implemented until a new type comes in.
func (target *Target) ApplyFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// Parse gets target options from user input.
func (target *Target) Parse(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

// parseOCILayoutReference parses the raw in format of <path>[:<tag>|@<digest>]
func (target *Target) parseOCILayoutReference() error { _ = "STUB: not implemented"; return nil }

// `digest` found

// find `tag`

func (target *Target) newOCIStore() (*oci.Store, error) { _ = "STUB: not implemented"; return nil, nil }

func (target *Target) newRepository(common Common, logger logrus.FieldLogger) (*remote.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewTarget generates a new target based on target.
func (target *Target) NewTarget(common Common, logger logrus.FieldLogger) (oras.GraphTarget, error) {
	_ = "STUB: not implemented"
	return *new(oras.GraphTarget), nil
}

// NewBlobDeleter generates a new blob deleter based on target.
func (target *Target) NewBlobDeleter(common Common, logger logrus.FieldLogger) (ResolvableDeleter, error) {
	_ = "STUB: not implemented"
	return *new(ResolvableDeleter), nil
}

// NewManifestDeleter generates a new blob deleter based on target.
func (target *Target) NewManifestDeleter(common Common, logger logrus.FieldLogger) (ResolvableDeleter, error) {
	_ = "STUB: not implemented"
	return *new(ResolvableDeleter), nil
}

// NewReadonlyTarget generates a new read only target based on target.
func (target *Target) NewReadonlyTarget(ctx context.Context, common Common, logger logrus.FieldLogger) (ReadOnlyGraphTagFinderTarget, error) {
	_ = "STUB: not implemented"
	return *new(ReadOnlyGraphTagFinderTarget), nil
}

// EnsureReferenceNotEmpty returns formalized error when the reference is empty.
func (target *Target) EnsureReferenceNotEmpty(cmd *cobra.Command, allowTag bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ModifyError handles error during cmd execution.
func (target *Target) ModifyError(cmd *cobra.Command, err error) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// short circuit for non-remote targets
		nil
}

// handle errors for remote targets

// special handling for not found error returned by registry target

// short circuit if the error is not an ErrorResponse

// raw reference is not registry host

// this should not happen

// not handle if the error is not from the target

// docker.io/xxx -> docker.io/library/xxx
