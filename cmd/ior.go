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

	"github.com/Solarcode-org/Orion/ast"
	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/builtins"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

/*
Run an individual Orion source file.

Usage:

	orion ior [flags]

Examples:
# Run foo.or

	$ orion ior foo.or

Flags:

	-h, --help   help for ior

Global Flags:

	    --config string    config file (default is $HOME/.orion.yaml)
	-v, --verbose uint32   config file (default is $HOME/.Orion.yaml)
*/
var iorCmd = &cobra.Command{
	Use:   "ior",
	Short: "Run an individual Orion source file",
	Long:  `Run an individual Orion source file.`,
	Example: `# Run foo.or
	$ orion ior foo.or`,

	PreRun: toggleDebug,
	Args:   cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Tracef("started `orion ior` with args %v\n", args)

		contents, err := os.ReadFile(args[0])
		lib.CheckErr(err)

		contents = append(contents, "\n"...)

		astree, errs := lib.GetAbstractSyntaxTree(contents)

		if len(errs) > 0 {
			lib.FailParse(errs)
		}

		log.Tracef("Parsed into Abstract Syntax Tree: %v\n", astree)

		builtins.MakeFunctions()

		for i := 0; i < len(astree); i++ {
			stmt := astree[i]

			switch stmt.Type {
			case ast.Expr_FuncCall:
				lib.RunFunc(*stmt, builtins.Functions)
			}
		}

		log.Traceln("ended `orion ior` with exit code 0")
	},
}

func init() {
	RootCmd.AddCommand(iorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// iorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// iorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
