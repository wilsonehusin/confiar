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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/wilsonehusin/confiar/internal"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"buildinfo"},
	Short:   "version",
	Long:    `versions related to this build`,
	Run: func(*cobra.Command, []string) {
		versions := *internal.BuildInfo()
		fmt.Printf("%v: %v\n", "Version", versions["Version"])
		for k, v := range versions {
			if k == "Version" {
				continue
			}
			fmt.Printf("%v: %v\n", k, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
