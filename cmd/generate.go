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
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/wilsonehusin/confiar/internal"
)

var nameList string
var ipList string

var names []string
var ips []string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate HOST_FQDN",
	Short: "Generate new TLS certificate",
	Long: `confiar generate -- create new TLS certificate

Files will be created in working directory as cert.pem and key.pem, if any of
those files already exist, they will be overwritten.

Specifications:
  - has itself as certificate authority (CA)
	- is valid starting 1 hour ago until 30 days from now
	- uses ECDSA P-521 (FIPS 186-3) aka. secp521r1
`,
	Args: cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		if nameList != "" {
			names = strings.Split(nameList, ",")
			for _, name := range names {
				if !internal.ValidFQDN(name) {
					fmt.Fprintf(os.Stderr, "Error: \"%v\" is not a valid fully qualified domain name (FQDN)\n", name)
					os.Exit(1)
				}
			}
		}
		if ipList != "" {
			ips = strings.Split(ipList, ",")
			for _, ip := range ips {
				if !internal.ValidIPAddr(ip) {
					fmt.Fprintf(os.Stderr, "Error: \"%v\" is not a valid IP address\n", ip)
					os.Exit(1)
				}
			}
		}
		if nameList == "" && ipList == "" {
			// both nameList and ipList are empty string
			fmt.Fprintf(os.Stderr, "Error: --fqdn and --ip cannot both be blank\n")
			os.Exit(1)
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.NewTLSSelfAuthority("gostd", names, ips)
	},
}

func init() {
	generateCmd.Flags().StringVar(&nameList, "fqdn", "", "domain name(s) for certificate (comma separated)")
	generateCmd.Flags().StringVar(&ipList, "ip", "", "IP address(es) for certificate (comma separated)")
	rootCmd.AddCommand(generateCmd)
}
