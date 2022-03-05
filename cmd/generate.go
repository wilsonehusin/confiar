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

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/wilsonehusin/confiar/internal"
)

var outDir string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate new TLS certificate",
	Long: `confiar generate -- create new TLS certificate

Files will be created in working directory as cert.pem and key.pem, if any of
those files already exist, they will be overwritten.

Specifications:
 	- has itself as certificate authority (CA)
	- is valid starting 1 hour ago until 365 days from now
	- uses ECDSA P-521 (FIPS 186-3) aka. secp521r1
`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateNameAndIP(true); err != nil {
			return err
		}
		return internal.NewTLSSelfAuthority("gostd", names, ips, outDir)
	},
}

func init() {
	generateCmd.Flags().StringVar(&outDir, "out-dir", ".", "directory where certificate will be written to")
	generateCmd.Flags().StringVar(&nameList, "fqdn", "", "domain name(s) for certificate (comma separated)")
	generateCmd.Flags().StringVar(&ipList, "ip", "", "IP address(es) for certificate (comma separated)")
	rootCmd.AddCommand(generateCmd)
}
