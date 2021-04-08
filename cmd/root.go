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
	"path"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var debug bool
var json bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "confiar",
	Short: "confiar -- self-signed TLS certificates made easy",
	Long: `Confiar lets you generate and distribute your self-signed
certificates easily. The goal is to make any application to run as if valid
certificates are in place between hosts.

 !!! YOU SHOULD CONSIDER USING A REAL CERTIFICATE BEFORE ANYTHING ELSE !!!

In scenarios where it doesn't make sense to do so, consider at least having
a centralized certificate authority (CA).

And maybe it doesn't make sense to do that either, so let's do it together!`,
	PersistentPreRun: func(*cobra.Command, []string) {
		prepareLogger()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func prepareLogger() {
	// UTC or GTFO
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	level := zerolog.InfoLevel
	if debug {
		level = zerolog.DebugLevel
		log.Logger = log.With().Caller().Logger()

		// hide GOPATH and just use filename
		zerolog.CallerMarshalFunc = func(file string, line int) string {
			return path.Base(file) + ":" + strconv.Itoa(line)
		}
	}
	zerolog.SetGlobalLevel(level)

	if !json {
		// turn on pretty logging
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "toggle debug logs")
	rootCmd.PersistentFlags().BoolVar(&json, "json", false, "print logs as json")
}
