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
	"github.com/Solarcode-org/Orion/utils"
	"github.com/shopspring/decimal"
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
		utils.CheckErr(err)

		contents = append(contents, "\n"...)

		builtins.MakeFunctions()
		builtins.MakeVariables(verbose)

		astree, parseErrs, err := lib.ParsedFrom(contents)
		if len(parseErrs) > 0 {
			utils.FailParse(parseErrs)
		}
		utils.CheckErr(err)

		log.Tracef("Parsed into Abstract Syntax Tree: %v\n", astree)

		for i := 0; i < len(astree); i++ {
			stmt := astree[i]

			switch stmt.Type {
			case ast.Expr_FuncCall:
				_, err := utils.RunFunc(*stmt, builtins.Functions)

				utils.CheckErr(err)
			case ast.Expr_VariableDef:
				if stmt.Args[0].Type == ast.Expr_FuncCall {
					value, err := utils.RunFunc(*stmt.Args[0], builtins.Functions)
					utils.CheckErr(err)

					builtins.Variables[stmt.Id] = value
				} else if stmt.Args[0].Type == ast.Expr_Variable {
					builtins.Variables[stmt.Id] = builtins.Variables[stmt.Args[0].Id]
				} else {
					builtins.Variables[stmt.Id] = *stmt.Args[0]
				}
			case ast.Expr_VariableTypeDef:
				var val ast.Expr

				if stmt.Args[0].Type == ast.Expr_FuncCall {
					funcVal, err := utils.RunFunc(*stmt.Args[0], builtins.Functions)
					utils.CheckErr(err)

					val = funcVal
				} else if stmt.Args[0].Type == ast.Expr_Variable {
					val = builtins.Variables[stmt.Args[0].Id]
				} else {
					val = *stmt.Args[0]
				}

				switch stmt.Args[1].Type {
				case ast.Expr_String:
					builtins.Variables[stmt.Id] = ast.Expr{
						Type: ast.Expr_String,
						Id:   val.Id,
					}
				case ast.Expr_Number:
					if _, err := decimal.NewFromString(val.Id); err != nil {
						log.Fatalln(err)
					}

					builtins.Variables[stmt.Id] = ast.Expr{
						Type: ast.Expr_Number,
						Id:   val.Id,
					}
				}
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
