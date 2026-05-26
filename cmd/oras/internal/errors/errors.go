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

package errors

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"oras.land/oras-go/v2/registry/remote/errcode"
)

// OperationType stands for certain type of operations.
type OperationType int

const (
	// OperationTypeParseArtifactReference represents parsing artifact
	// reference operation.
	OperationTypeParseArtifactReference OperationType = iota + 1
)

// RegistryErrorPrefix is the commandline prefix for errors from registry.
const RegistryErrorPrefix = "Error response from registry:"

// UnsupportedFormatTypeError generates the error message for an invalid type.
type UnsupportedFormatTypeError string

// Error implements the error interface.
func (e UnsupportedFormatTypeError) Error() string { _ = "STUB: not implemented"; return "" }

// Error is the error type for CLI error messaging.
type Error struct {
	OperationType  OperationType
	Err            error
	Usage          string
	Recommendation string
}

// Unwrap implements the errors.Wrapper interface.
func (o *Error) Unwrap() error {
	_ = "STUB: not implemented"

	// Error implements the error interface.
	return nil
}

func (o *Error) Error() string { _ = "STUB: not implemented"; return "" }

// CheckArgs checks the args with the checker function.
func CheckArgs(checker func(args []string) (bool, string), Usage string) cobra.PositionalArgs {
	_ = "STUB: not implemented"
	return *new(cobra.PositionalArgs)
}

// Modifier modifies the error during cmd execution.
type Modifier interface {
	ModifyError(cmd *cobra.Command, err error) (modified bool, modifiedErr error)
}

// Command returns an error-handled cobra command.
func Command(cmd *cobra.Command, handler Modifier) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// ReportErrResp returns the inner error message from errResp.Errors.
// If errResp.Errors is empty, it returns the original errResp.
func ReportErrResp(errResp *errcode.ErrorResponse) error { _ = "STUB: not implemented"; return nil }

// Example error string:
// GET "registry.example.com/v2/_catalog": response status code 401: 401

// Example error string:
// unauthorized: authentication required

// UnwrapCopyError extracts the underlying error from an oras.CopyError.
// If err is of type *oras.CopyError, it returns the inner error (copyErr.Err).
// Otherwise, it returns the original error unchanged.
func UnwrapCopyError(err error) error { _ = "STUB: not implemented"; return nil }

// TrimErrBasicCredentialNotFound trims the credentials from err.
// Caller should make sure the err is auth.ErrBasicCredentialNotFound.
func TrimErrBasicCredentialNotFound(err error) error { _ = "STUB: not implemented"; return nil }

// reWrap re-wraps outer to inner by trimming out mid, returns inner if extraction fails.
// +---------- outer ----------+      +------ outer ------+
// |         +---- mid ----+   |      |                   |
// |         |    inner    |   |  =>  |       inner       |
// |         +-------------+   |      |                   |
// +---------------------------+      +-------------------+
func reWrap(outer, mid, inner error) error { _ = "STUB: not implemented"; return nil }

// NewErrEmptyTagOrDigest creates a new error based on the reference string.
func NewErrEmptyTagOrDigest(ref string, cmd *cobra.Command, needsTag bool) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckMutuallyExclusiveFlags checks if any mutually exclusive flags are used
// at the same time, returns an error when detecting used exclusive flags.
func CheckMutuallyExclusiveFlags(fs *pflag.FlagSet, exclusiveFlagSet ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckRequiredTogetherFlags checks if any flags required together are all used,
// returns an error when detecting any flags not used while other flags have been used.
func CheckRequiredTogetherFlags(fs *pflag.FlagSet, requiredTogetherFlags ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkChangedFlags(fs *pflag.FlagSet, flagSet ...string) (changedFlags []string, unchangedFlags []string) {
	_ = "STUB: not implemented"
	return nil, nil
}
