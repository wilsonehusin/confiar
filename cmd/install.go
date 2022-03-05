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

var installTarget string

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs certificates in the defined target(s)",
	Long: `confiar install -- let your computer trust your certificate

For targets requiring specific hostnames to be mapped to each certificate,
confiar will automatically parse the information from the given certificate.

You can pass additional --fqdn or --ip for hostnames which were not included
in the certificate.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := validateNameAndIP(false); err != nil {
			return err
		}
		return internal.InstallTLS(certSrc, installTarget, names, ips)
	},
}

func init() {
	installCmd.Flags().StringVarP(&installTarget, "target", "t", "stdout", "installation target")
	installCmd.Flags().StringVarP(&certSrc, "from", "f", "./cert.pem", "where to find the certificate")
	installCmd.Flags().StringVar(&nameList, "fqdn", "", "additional domain name(s) for certificate (comma separated)")
	installCmd.Flags().StringVar(&ipList, "ip", "", "additional IP address(es) for certificate (comma separated)")
	rootCmd.AddCommand(installCmd)
}
