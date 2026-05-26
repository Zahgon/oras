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

package utils

import (
	"io"
	"time"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"oras.land/oras/test/e2e/internal/utils/match"
)

const (
	orasBinary = "oras"

	// customize your own basic auth file via `htpasswd -cBb <file_name> <user_name> <password>`
	Username       = "hello"
	Password       = "oras-test"
	DefaultTimeout = 10 * time.Second
	// If the command hasn't exited yet, ginkgo session ExitCode is -1
	notResponding = -1
)

// ExecOption provides option used to execute a command.
type ExecOption struct {
	binary  string
	args    []string
	workDir string
	timeout time.Duration

	stdin    io.Reader
	stdout   []match.Matcher
	stderr   []match.Matcher
	exitCode int

	text string
}

// ORAS returns default execution option for oras binary.
func ORAS(args ...string) *ExecOption { _ = "STUB: not implemented"; return nil }

// Binary returns default execution option for customized binary.
func Binary(path string, args ...string) *ExecOption { _ = "STUB: not implemented"; return nil }

// ExpectFailure sets failure exit code checking for the execution.
func (opts *ExecOption) ExpectFailure() *ExecOption {
	_ = "STUB: not implemented"
	// set to 1 but only check if it's positive
	return nil
}

// ExpectBlocking consistently check if the execution is blocked.
func (opts *ExecOption) ExpectBlocking() *ExecOption { _ = "STUB: not implemented"; return nil }

// WithTimeOut sets timeout for the execution.
func (opts *ExecOption) WithTimeOut(timeout time.Duration) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// WithDescription sets description text for the execution.
func (opts *ExecOption) WithDescription(text string) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// WithWorkDir sets working directory for the execution.
func (opts *ExecOption) WithWorkDir(path string) *ExecOption { _ = "STUB: not implemented"; return nil }

// WithInput redirects stdin to r for the execution.
func (opts *ExecOption) WithInput(r io.Reader) *ExecOption { _ = "STUB: not implemented"; return nil }

// MatchKeyWords adds keywords matching to stdout.
func (opts *ExecOption) MatchKeyWords(keywords ...string) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// MatchErrKeyWords adds keywords matching to stderr.
func (opts *ExecOption) MatchErrKeyWords(keywords ...string) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// MatchRequestHeaders adds a debug log matcher that checks for
// the presence of provided headers in all outgoing HTTP requests.
func (opts *ExecOption) MatchRequestHeaders(headers ...string) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// MatchCpRequestHeaders adds a debug log matcher that checks for
// the presence of specific headers in outgoing copy requests towards
// the provided host and repository.
func (opts *ExecOption) MatchCpRequestHeaders(host string, repo string, headers ...string) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// MatchContent adds full content matching to the execution.
func (opts *ExecOption) MatchContent(content string) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// MatchTrimedContent adds trimmed content matching to the execution.
func (opts *ExecOption) MatchTrimmedContent(content string) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// MatchStatus adds full content matching to the execution option.
func (opts *ExecOption) MatchStatus(keys []match.StateKey, verbose bool, successCount int) *ExecOption {
	_ = "STUB: not implemented"
	return nil
}

// redactArgs returns a copy of args with values following password flags replaced by "***".
func redactArgs(args []string) []string { _ = "STUB: not implemented"; return nil }

// Exec run the execution based on opts.
func (opts *ExecOption) Exec() *gexec.Session {
	_ = "STUB: not implemented"

	// this should be a code error but can only be caught during runtime
	return nil
}

// set default description text

// switch working directory

// matching result
