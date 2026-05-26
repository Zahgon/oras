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
	"crypto/tls"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
	"oras.land/oras-go/v2/registry/remote/credentials"
	oerrors "oras.land/oras/cmd/oras/internal/errors"
	onet "oras.land/oras/internal/net"
)

const (
	caFileFlag                 = "ca-file"
	certFileFlag               = "cert-file"
	keyFileFlag                = "key-file"
	usernameFlag               = "username"
	passwordFlag               = "password"
	passwordFromStdinFlag      = "password-stdin"
	identityTokenFlag          = "identity-token"
	identityTokenFromStdinFlag = "identity-token-stdin"
)

// Remote options struct contains flags and arguments specifying one registry.
// Remote implements oerrors.Handler and interface.
type Remote struct {
	DistributionSpec
	CACertFilePath  string
	CertFilePath    string
	KeyFilePath     string
	Insecure        bool
	Configs         []string
	Username        string
	secretFromStdin bool
	Secret          string //nolint:gosec // G117: not a hardcoded secret, this is a CLI option field
	flagPrefix      string

	resolveFlag           []string
	applyDistributionSpec bool
	headerFlags           []string
	headers               http.Header
	warned                map[string]*sync.Map
	plainHTTP             func() (plainHTTP bool, enforced bool)
	store                 credentials.Store
}

// EnableDistributionSpecFlag set distribution specification flag as applicable.
func (remo *Remote) EnableDistributionSpecFlag() { _ = "STUB: not implemented"; return }

// ApplyFlags applies flags to a command flag set.
func (remo *Remote) ApplyFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

func (remo *Remote) applyStdinFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// ApplyFlagsWithPrefix applies flags to a command flag set with a prefix string.
// Commonly used for non-unary remote targets.
func (remo *Remote) ApplyFlagsWithPrefix(fs *pflag.FlagSet, prefix, description string) {
	_ = "STUB: not implemented"
	return
}

// CheckStdinConflict checks if PasswordFromStdin or IdentityTokenFromStdin of a
// *pflag.FlagSet conflicts with read file from input.
func CheckStdinConflict(flags *pflag.FlagSet) error { _ = "STUB: not implemented"; return nil }

// Parse tries to read password with optional cmd prompt.
func (remo *Remote) Parse(cmd *cobra.Command) error { _ = "STUB: not implemented"; return nil }

// readSecret tries to read password or identity token with
// optional cmd prompt.
func (remo *Remote) readSecret(cmd *cobra.Command) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Prompt for credential

// parseResolve parses resolve flag.
func (remo *Remote) parseResolve(baseDial onet.DialFunc) (onet.DialFunc, error) {
	_ = "STUB: not implemented"
	return *new(onet.DialFunc), nil
}

// ipv6 zone is not parsed

// tlsConfig assembles the tls config.
func (remo *Remote) tlsConfig() (*tls.Config, error) { _ = "STUB: not implemented"; return nil, nil }

// authClient assembles a oras auth client.
func (remo *Remote) authClient(_ string, debug bool) (client *auth.Client, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// http.RoundTripper with a retry using the DefaultPolicy
// see: https://pkg.go.dev/oras.land/oras-go/v2/registry/remote/retry#Policy

// ConfigPath returns the config path of the credential store.
func (remo *Remote) ConfigPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (remo *Remote) parseCustomHeaders() error { _ = "STUB: not implemented"; return nil }

// In conformance to the RFC 2616 specification
// Reference: https://www.rfc-editor.org/rfc/rfc2616#section-4.2

// Credential returns a credential based on the remote options.
func (remo *Remote) Credential() auth.Credential {
	_ = "STUB: not implemented"
	return *new(auth.Credential)
}

func (remo *Remote) handleWarning(registry string, logger logrus.FieldLogger) func(warning remote.Warning) {
	_ = "STUB: not implemented"
	return nil
}

// NewRegistry assembles a oras remote registry.
func (remo *Remote) NewRegistry(registry string, common Common, logger logrus.FieldLogger) (reg *remote.Registry, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewRepository assembles a oras remote repository.
func (remo *Remote) NewRepository(reference string, common Common, logger logrus.FieldLogger) (repo *remote.Repository, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isPlainHTTP returns the plain http flag for a given registry.
func (remo *Remote) isPlainHTTP(registry string) bool { _ = "STUB: not implemented"; return false }

// not specified, defaults to plain http for localhost

// ModifyError modifies error during cmd execution.
func (remo *Remote) ModifyError(cmd *cobra.Command, err error) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// DecorateCredentialError decorate error with recommendation.
func (remo *Remote) DecorateCredentialError(err error) *oerrors.Error {
	_ = "STUB: not implemented"
	return nil
}
