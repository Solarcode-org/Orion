/*
Copyright © 2024 Arnab Phukan <iamarnab.phukan@gmail.com>

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
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var manpagesPath string

/*
Generate manpages for Orion.

Specifiy the '-P' flag to generate manpages in custom directory (default=./manpages/)

Usage:

	orion manpages [flags]

Examples:

	# Generate manpages in default directory.
	$ orion manpages

	# Generate manpages in foo directory
	$ orion manpages -P foo

Flags:

	-h, --help          help for manpages
	-P, --path string   Custom directory to generate manpages in (default $PWD/manpages)

Global Flags:

	    --config string   config file (default is $HOME/.orion.yaml)
	-v, --verbose uint8   verbosity level
*/
var manpagesCmd = &cobra.Command{
	Use:   "manpages [-P path]",
	Short: "Generate manpages for Orion",
	Long: `Generate manpages for Orion.

Specifiy the '-P' flag to generate manpages in custom directory (default=./manpages/)`,
	Example: `
	# Generate manpages in default directory.
	$ orion manpages
	
	# Generate manpages in foo directory
	$ orion manpages -P foo`,
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()

		header := &doc.GenManHeader{
			Title:   "orion",
			Section: "1",
			Date:    &now,
		}
		cobra.CheckErr(doc.GenManTree(cmd.Root(), header, manpagesPath))
	},
}

func init() {
	RootCmd.AddCommand(manpagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// manpagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	manpagesCmd.Flags().StringVarP(&manpagesPath, "path", "P", "./manpages", "Custom directory to generate manpages in")
}
