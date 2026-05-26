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
	"github.com/spf13/pflag"
	"oras.land/oras-go/v2"
)

const (
	ImageSpecV1_1 = "v1.1"
	ImageSpecV1_0 = "v1.0"
)

const (
	DistributionSpecReferrersUnknown = "unknown"
	DistributionSpecReferrersTagV1_1 = "v1.1-referrers-tag"
	DistributionSpecReferrersAPIV1_1 = "v1.1-referrers-api"
)

// ImageSpec option struct which implements pflag.Value interface.
type ImageSpec struct {
	Flag        string
	PackVersion oras.PackManifestVersion
}

// Set validates and sets the flag value from a string argument.
func (is *ImageSpec) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Type returns the string value of the inner flag.
func (is *ImageSpec) Type() string {
	_ = "STUB: not implemented"

	// Options returns the string of usable options for the flag.
	return ""
}

func (is *ImageSpec) Options() string { _ = "STUB: not implemented"; return "" }

// String returns the string representation of the flag.
func (is *ImageSpec) String() string {
	_ = "STUB: not implemented"
	// to avoid printing default value in usage doc
	return ""
}

// ApplyFlags applies flags to a command flag set.
func (is *ImageSpec) ApplyFlags(fs *pflag.FlagSet) {
	_ = "STUB: not implemented"
	// default to v1.1, unless --config is used and --artifact-type is not used
	return
}

// referrersState represents the state of Referrers API.
type referrersState = int32

const (
	// ReferrersStateUnknown represents an unknown state of Referrers API.
	ReferrersStateUnknown referrersState = iota
	// ReferrersStateSupported represents that the repository is known to
	// support Referrers API.
	ReferrersStateSupported
	// ReferrersStateUnsupported represents that the repository is known to
	// not support Referrers API.
	ReferrersStateUnsupported
)

// DistributionSpec option struct which implements pflag.Value interface.
type DistributionSpec struct {
	// ReferrersAPI indicates the preference of the implementation of the Referrers API.
	// Set to true for referrers API, false for referrers tag scheme, and nil for auto fallback.
	ReferrersAPI referrersState

	// specFlag should be provided in form of`<version>-<api>-<option>`
	flag string
}

// Set validates and sets the flag value from a string argument.
func (ds *DistributionSpec) Set(value string) error { _ = "STUB: not implemented"; return nil }

// Type returns the string value of the inner flag.
func (ds *DistributionSpec) Type() string {
	_ = "STUB: not implemented"

	// Options returns the string of usable options for the flag.
	return ""
}

func (ds *DistributionSpec) Options() string { _ = "STUB: not implemented"; return "" }

// String returns the string representation of the flag.
func (ds *DistributionSpec) String() string {
	_ = "STUB: not implemented"

	// ApplyFlagsWithPrefix applies flags to a command flag set with a prefix string.
	return ""
}

func (ds *DistributionSpec) ApplyFlagsWithPrefix(fs *pflag.FlagSet, prefix, description string) {
	_ = "STUB: not implemented"
	return
}
