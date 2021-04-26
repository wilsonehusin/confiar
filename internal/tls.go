/*
Copyright © 2021 Wilson Husin

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

package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/wilsonehusin/confiar/internal/cryptographer"
	"github.com/wilsonehusin/confiar/internal/target"
)

var cryptoBackend cryptographer.Cryptographer
var installTarget target.Target

func NewTLSSelfAuthority(backendType string, names []string, ips []string, outDir string) error {
	switch backendType {
	case "gostd":
		cryptoBackend = &cryptographer.GoStd{}
	default:
		return fmt.Errorf("unknown cryptographer backend type: %w", backendType)
	}
	return cryptoBackend.NewTLSSelfAuthority(names, ips, outDir)
}

func InstallTLS(certSrc string, targetType string, extraNames []string, extraIPs []string) error {
	removeTempCert := false

	var certPath string

	if strings.HasPrefix(certSrc, "http://") {
		log.Info().Str("certSrc", certSrc).Msg("downloading to local path")

		resp, err := http.Get(certSrc)
		if err != nil {
			return fmt.Errorf("unable to get certificate from remote: %w", err)
		}
		defer resp.Body.Close()

		tempCertPath, err := ioutil.TempFile(os.TempDir(), "confiar-cert-")
		if err != nil {
			return fmt.Errorf("unable to create local copy of certificate: %w", err)
		}
		if _, err := io.Copy(tempCertPath, resp.Body); err != nil {
			return fmt.Errorf("unable to write certificate to local: %w", err)
		}
		certPath = tempCertPath.Name()
		removeTempCert = true
	} else {
		certPath = certSrc
	}

	log.Info().Str("certPath", certPath).Msg("installing certificate")
	switch targetType {
	case "stdout":
		installTarget = &target.Stdout{
			CertPath: certPath,
		}
	case "docker":
		installTarget = &target.Docker{
			CertPath:   certPath,
			ExtraHosts: append(extraNames, extraIPs...),
		}
	default:
		return fmt.Errorf("unknown installation target: %s", targetType)
	}
	if err := installTarget.Install(); err != nil {
		return fmt.Errorf("install certificate: %w", err)
	}

	if removeTempCert {
		// only remove tempfile on successful install
		// leave the tempfile for debugging on failure
		log.Info().Str("certPath", certPath).Msg("removing cert from local path")
		os.Remove(certPath)
	}

	return nil
}
