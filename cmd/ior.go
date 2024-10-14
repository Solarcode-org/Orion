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
	"os"

	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/builtins"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// iorCmd represents the ior command
var iorCmd = &cobra.Command{
	Use:   "ior",
	Short: "Run an individual Orion source file",
	Long:  `Run an individual Orion source file.`,
	/*Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.*/

	PreRun: toggleDebug,
	Args:   cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Tracef("started `orion ior` with args %v\n", args)

		builtins.MakeFunctions()

		contents, err := os.ReadFile(args[0])
		lib.HandleFatal(err)

		contents = append(contents, "\n"...)

		astree, err := lib.GetAbstractSyntaxTree(contents)
		lib.HandleFatal(err)

		log.Debugf("Parsed into Abstract Syntax Tree: %v", astree)

		for i := 0; i < len(astree); i++ {
			funcCall := astree[i]

			if function, ok := builtins.Functions[funcCall.Name]; ok {
				_, err := function(funcCall.Args)
				if err != nil {
					log.Fatalf("%s: %s\n", funcCall.Name, err)
				}
			} else {
				caser := cases.Title(language.AmericanEnglish)

				if _, ok := builtins.Functions[caser.String(funcCall.Name)]; ok {
					log.Fatalf("Could not find function: %s\nDid you mean: %s?\n", funcCall.Name, caser.String(funcCall.Name))
				}
				log.Fatalf("Could not find function: %s\nMaybe you forgot to add a module prefix?\n", funcCall.Name)
			}
		}

		log.Traceln("ended `orion ior` with exit code 0")
	},
}

func init() {
	rootCmd.AddCommand(iorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
