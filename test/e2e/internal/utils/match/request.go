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

package match

import (
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

// requestHeaderMatcher provides matching for request headers.
// Given a url prefix, it looks at the requests sent to the urls which
// match the given prefix, and check if these requests all contain
// the specified headers.
type requestHeaderMatcher struct {
	urlPrefix string
	headers   []string
}

// MatchCpRequestHeaders returns a request header matcher
// with the given url prefix.
func NewRequestHeaderMatcher(urlPrefix string, headers []string) requestHeaderMatcher {
	_ = "STUB: not implemented"
	return *new(requestHeaderMatcher)
}

// Match matches got with wanted headers.
func (r requestHeaderMatcher) Match(got *gbytes.Buffer) { _ = "STUB: not implemented"; return }

// getRequests parses raw debug output to a string slice
// containing each request that match the given prefix.
func getRequests(urlPrefix string, debugOutput string) []string {
	_ = "STUB: not implemented"
	return nil
}

// trim the response content

// filter with the url prefix

// extract request url to match the prefix

// getRequestHeaders takes a string slice containing requests
// and extract request headers from them.
func getRequestHeaders(reqs []string) []string { _ = "STUB: not implemented"; return nil }

// extract the header content from each request
