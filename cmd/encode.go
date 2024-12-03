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
	"fmt"
	"os"

	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/builtins"
	"github.com/Solarcode-org/Orion/lib/bytecode"
	"github.com/Solarcode-org/Orion/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

/*
Encode a parsed Orion file into bytecode.

This is really useful for just storing the
syntax tree for later use in the form of cache.

Usage:

	orion encode [flags]

Examples:

	# Encode file "foo.or"
	orion encode foo.or

Flags:

	-h, --help   help for encode

Global Flags:

	    --config string   config file (default is $HOME/.orion.yaml)
	-v, --verbose uint8   verbosity level
*/
var encodeCmd = &cobra.Command{
	Use:   "encode file",
	Short: "Encode a parsed Orion file into bytecode",
	Long: `Encode a parsed Orion file into bytecode.

This is really useful for just storing the
syntax tree for later use in the form of cache.`,
	Example: `
	# Encode file "foo.or"
	orion encode foo.or
`,
	Args:   cobra.ExactArgs(1),
	PreRun: toggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		log.Tracef("started `orion ior` with args %v\n", args)

		contents, err := os.ReadFile(args[0])
		utils.CheckErr(err)

		contents = append(contents, "\n"...)

		builtins.MakeFunctions()
		builtins.MakeVariables(verbose)

		astree, parseErrs, err := lib.ParsedFrom(contents)
		if len(parseErrs) > 0 {
			utils.FailParse(parseErrs)
		}
		utils.CheckErr(err)

		log.Tracef("Parsed into Abstract Syntax Tree: %+#v\n", astree)

		instructions, err := bytecode.EncodedSyntaxTree(astree)
		utils.CheckErr(err)

		astFilePath := fmt.Sprint(args[0], ".ast")

		astFile, err := os.Create(astFilePath)

		if err != nil {
			astFile.Close()
			log.Fatalln(err)
		}

		defer astFile.Close()

		if _, err := astFile.Write(instructions); err != nil {
			astFile.Close()
			log.Fatalln(err)
		}

		log.Traceln("ended `orion encode` with exit code 0")
	},
}

func init() {
	RootCmd.AddCommand(encodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
