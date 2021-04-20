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

package cryptographer

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

type GoStd struct {
	priv     *ecdsa.PrivateKey
	derBytes []byte
	outDir   string
}

// inspired by: https://golang.org/src/crypto/tls/generate_cert.go
func (g *GoStd) NewTLSSelfAuthority(names []string, ips []string, outDir string) error {
	log.Info().Strs("names", names).Strs("ips", ips).Str("outDir", outDir).Send()
	g.outDir = outDir
	priv, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate private key")
	}
	g.priv = priv
	keyUsage := x509.KeyUsageDigitalSignature

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate serial number")
	}

	now := time.Now()
	validFrom := now.Add(-1 * time.Hour) // prevent issues from cross-machine time gap
	validUntil := now.Add(365 * 24 * time.Hour)

	log.Info().Time("validFrom", validFrom).Time("validUntil", validUntil).Msg("certificate valid lifetime")

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			// filling in for the luls, mostly in case they're helpful for troubleshooting
			Country:            []string{"Confiar Country"},
			Organization:       []string{"Confiar Organization"},
			OrganizationalUnit: []string{"Confiar Organizational Unit"},
			Locality:           []string{"Confiar Locality"},
			Province:           []string{"Confiar Province"},
		},
		NotBefore:             validFrom,  // ugh
		NotAfter:              validUntil, // ugh x2
		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	template.KeyUsage |= x509.KeyUsageCertSign

	template.DNSNames = append(template.DNSNames, names...)
	for _, ip := range ips {
		template.IPAddresses = append(template.IPAddresses, net.ParseIP(ip))
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &g.priv.PublicKey, g.priv)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate certificate")
	}
	g.derBytes = derBytes

	var writeWaiter sync.WaitGroup
	writeWaiter.Add(2)

	go func() {
		defer writeWaiter.Done()
		if err := g.writeCertFile(); err != nil {
			log.Fatal().Err(err).Str("filename", certFileName).Msg("writing file")
		}
		log.Info().Str("filename", certFileName).Msg("wrote file")
	}()

	go func() {
		defer writeWaiter.Done()
		if err := g.writeKeyFile(); err != nil {
			log.Fatal().Err(err).Str("filename", keyFileName).Msg("writing file")
		}
		log.Info().Str("filename", keyFileName).Msg("wrote file")
	}()

	writeWaiter.Wait()

	for _, filename := range []string{certFileName, keyFileName} {
		if _, err := os.Stat(filename); err != nil {
			log.Fatal().Err(err).Str("filename", filename).Msg("filestat error")
		}
	}

	return nil
}

func (g *GoStd) writeCertFile() error {
	certFile, err := os.Create(path.Join(g.outDir, certFileName))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	if err := pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: g.derBytes}); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	if err := certFile.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}
	return nil
}

func (g *GoStd) writeKeyFile() error {
	keyFile, err := os.OpenFile(path.Join(g.outDir, keyFileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(g.priv)
	if err != nil {
		return fmt.Errorf("failed to marshal private key: %w", err)
	}
	if err := pem.Encode(keyFile, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	if err := keyFile.Close(); err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	}
	return nil
}
