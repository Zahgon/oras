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
	"context"
	"errors"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"oras.land/oras-go/v2"
	"oras.land/oras/cmd/oras/internal/display/metadata"
	"oras.land/oras/cmd/oras/internal/option"
)

// outputFormat defines the format of the backup output.
type outputFormat int

const (
	// outputFormatDir indicates the output is a directory.
	outputFormatDir outputFormat = iota
	// outputFormatTar indicates the output is a tar archive.
	outputFormatTar
)

// errTagListNotSupported is returned when the target does not support tag listing.
var errTagListNotSupported = errors.New("the target does not support tag listing")

type backupOptions struct {
	option.Common
	option.Remote
	option.Terminal

	// flags
	output           string
	includeReferrers bool
	concurrency      int

	// derived options
	outputFormat outputFormat
	repository   string
	tags         []string
}

func backupCmd() *cobra.Command { _ = "STUB: not implemented"; return nil }

// parse repo and references

// parse output format

// always print verbose output

// required flags

// optional flags

// apply flags

func runBackup(cmd *cobra.Command, opts *backupOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// start timing the backup process

// test if the output file can be created and fail early if there is an issue

// create a temporary directory as the working directory for OCI store

// this should not happen, just a safeguard

// Prepare copy source and destination

// Resolve tags to back up

// Prepare copy options

// backupTag copies the artifact identified by the tag from src to dst.
func backupTag(ctx context.Context, src oras.ReadOnlyGraphTarget, dst oras.GraphTarget, tag string, root ocispec.Descriptor, copyGraphOpts oras.CopyGraphOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// backupTagWithReferrers copies the artifact identified by tag and its referrers from src to dst.
func backupTagWithReferrers(ctx context.Context, src oras.ReadOnlyGraphTarget, dst oras.GraphTarget, tag string, root ocispec.Descriptor, extCopyGraphOpts oras.ExtendedCopyGraphOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// countReferrers counts the total number of referrers for the given artifact identified by tag, including the referrers
// of its children manifests if the artifact is an image index or manifest list.
func countReferrers(ctx context.Context, target oras.ReadOnlyGraphTarget, tag string, root ocispec.Descriptor, extCopyGraphOpts oras.ExtendedCopyGraphOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If the root is not an image index or manifest list, we have counted all referrers

// count referrers of children manifests

// finalizeBackupOutput finalizes the backup output by removing temporary directories and exporting to a tar archive if needed.
func finalizeBackupOutput(dstRoot string, opts *backupOptions, logger logrus.FieldLogger, metadataHandler metadata.BackupHandler) (returnErr error) {
	_ = "STUB: not implemented"
	// Remove ingest dir for a cleaner output
	return nil
}

// If output format is not a tar, we are done

// exporting the backup to a tar archive

// remove the output file in case of error

// resolveTags resolves tags to their descriptors.
// It returns the resolved tags and their corresponding descriptors.
func resolveTags(ctx context.Context, target oras.ReadOnlyTarget, specifiedTags []string) ([]string, []ocispec.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// resolve the specified tags

// discover all tags in the repository and resolve them

// parseArtifactReferences parses the input string into a repository
// and a slice of tags.
func parseArtifactReferences(artifactRefs string) (string, []string, error) {
	_ = "STUB: not implemented"
	// validate input
	return "", nil, nil
}

// reject digest references early

// validate repository

// clear the tag

// no tags

// validate each tag
