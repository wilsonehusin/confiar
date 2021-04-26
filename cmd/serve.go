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

var servePort int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves generated certificate over HTTP",
	Long: `confiar serve -- serves generated certificate over HTTP

Sharing the generated certificate with other hosts. Clients who would like to
trust this certificate can run install using --from flag with current host as
address, e.g. --from http://10.11.12.13:8787`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return internal.ServeCertificate(certSrc, servePort)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&certSrc, "from", "f", "./cert.pem", "where to find the certificate")
	serveCmd.Flags().IntVarP(&servePort, "port", "p", 8787, "port to serve the certificate")

	rootCmd.AddCommand(serveCmd)
}
