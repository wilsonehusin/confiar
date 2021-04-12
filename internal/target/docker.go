/*
Copyright © 2021 Wilson Husin <wilsonehusin@gmail.com>

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

package target

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path"

	"github.com/rs/zerolog/log"
)

const dockerCertDir = "/etc/docker/certs.d"
const dockerCertFile = "ca.crt"

type Docker struct {
	CertPath   string
	ExtraHosts []string
	certBytes  []byte
}

func (d *Docker) Install() error {
	certBytes, err := os.ReadFile(d.CertPath)
	if err != nil {
		return err
	}
	d.certBytes = certBytes

	pemBlock, _ := pem.Decode(d.certBytes)
	if pemBlock == nil {
		return fmt.Errorf("failed to parse certificate PEM")
	}

	certData, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return err
	}

	names := append(certData.DNSNames, d.ExtraHosts...)
	for _, name := range names {
		if err := d.installHost(name); err != nil {
			return err
		}
	}
	for _, ipAddr := range certData.IPAddresses {
		if err := d.installHost(ipAddr.String()); err != nil {
			return err
		}
	}

	return nil
}

func (d *Docker) installHost(hostname string) error {
	fullpath := path.Join(dockerCertDir, hostname)
	log.Debug().Str("path", fullpath).Msg("creating directory")
	if err := os.MkdirAll(fullpath, 0755); err != nil {
		return err
	}

	dstpath := path.Join(fullpath, dockerCertFile)
	log.Debug().Str("file", dstpath).Msg("writing certificate")
	if err := os.WriteFile(dstpath, d.certBytes, 0644); err != nil {
		return err
	}

	log.Info().Str("file", dstpath).Str("hostname", hostname).Msg("certificate installed")
	return nil
}
