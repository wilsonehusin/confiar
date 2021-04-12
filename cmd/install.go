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

var certSrc string
var installTarget string

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs certificates in the defined target(s)",
	Long:  `confiar install -- let your computer trust your certificate`,
	PreRun: func(*cobra.Command, []string) {
		// TODO: find out if requireNameAndIP() is necessary
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.InstallTLS(certSrc, installTarget)
	},
}

func init() {
	installCmd.Flags().StringVarP(&installTarget, "target", "t", "stdout", "installation target")
	installCmd.Flags().StringVarP(&certSrc, "from", "f", "./cert.pem", "where to find the certificate")
	installCmd.Flags().StringVar(&nameList, "fqdn", "", "domain name(s) for certificate (comma separated)")
	installCmd.Flags().StringVar(&ipList, "ip", "", "IP address(es) for certificate (comma separated)")
	rootCmd.AddCommand(installCmd)
}
