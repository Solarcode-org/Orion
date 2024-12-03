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
	"crypto/sha1"
	"fmt"
	"os"
	"path"

	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/bytecode"
	"github.com/Solarcode-org/Orion/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectParentPath string

/*
Make a new Orion project.

This command creates a directory "project" in "-P" (which
is, by default, the current directory) with a starting
file in the "src" folder

Usage:

	orion new project [flags]

Examples:

	# Make a new project in the current directory.
	$ orion new foo

	# Make a new project in directory baz.
	$ orion new bar -P baz

Flags:

	-h, --help          help for new
	-P, --path string   The path where the project will be created (default $PWD)

Global Flags:

	    --config string   config file (default is $HOME/.orion.yaml)
	-v, --verbose uint8   verbosity level
*/
var newCmd = &cobra.Command{
	Use:   "new [-P path] project",
	Short: "Make a new Orion project",
	Long: `Make a new Orion project.

This command creates a directory "project" in "-P" (which
is, by default, the current directory) with a starting
file in the "src" folder`,
	Example: `
	# Make a new project in the current directory.
	$ orion new foo

	# Make a new project in directory baz.
	$ orion new bar -P baz
`,
	Args:   cobra.ExactArgs(1),
	PreRun: toggleDebug,
	Run: func(cmd *cobra.Command, args []string) {
		log.Traceln("started `orion new`")

		projectPath := path.Join(projectParentPath, args[0])
		srcPath := path.Join(projectPath, "src")
		mainFilePath := path.Join(srcPath, "main.or")
		buildPath := path.Join(projectPath, "_build")
		astFilePath := path.Join(buildPath, "main.ast")
		hashFilePath := path.Join(buildPath, "main.hash")

		log.Tracef(`compiler data:

		projectPath:	%s
		srcPath:	%s
		mainFilePath:	%s
		buildPath:	%s
		astFilePath:	%s
		hashFilePath:	%s
		`, projectPath, srcPath, mainFilePath, buildPath, astFilePath, hashFilePath)

		if err := os.Mkdir(projectPath, os.ModePerm); err != nil {
			utils.CheckErr(err)
		}
		if err := os.Mkdir(srcPath, os.ModePerm); err != nil {
			utils.CheckErr(err)
		}
		if err := os.Mkdir(buildPath, os.ModePerm); err != nil {
			utils.CheckErr(err)
		}

		//region main file
		f, err := os.Create(mainFilePath)

		if err != nil {
			f.Close()
			log.Fatalln(err)
		}

		defer f.Close()

		code := fmt.Sprintf(`package "%s"

world := "world"
Println("Hello", world)
`, projectPath)

		if _, err := f.Write([]byte(code)); err != nil {
			f.Close()
			log.Fatalln(err)
		}
		//endregion main file

		//region ast file
		contents, err := os.ReadFile(mainFilePath)
		utils.CheckErr(err)

		astree, parseErrs, err := lib.ParsedFrom(contents)
		if len(parseErrs) > 0 {
			utils.FailParse(parseErrs)
		}
		utils.CheckErr(err)

		instructions, err := bytecode.EncodedSyntaxTree(astree)
		utils.CheckErr(err)

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
		//endregion ast file

		//region hashed file
		hash := sha1.New()
		hash.Write(instructions)

		hashed := hash.Sum(nil)
		hashFile, err := os.Create(hashFilePath)

		if err != nil {
			hashFile.Close()
			log.Fatalln(err)
		}

		defer hashFile.Close()

		if _, err := hashFile.Write(hashed); err != nil {
			hashFile.Close()
			log.Fatalln(err)
		}
		//endregion hashed file

		log.Traceln("ended `orion new` with exit code 0")
	},
}

func init() {
	RootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	dir, err := os.Getwd()
	cobra.CheckErr(err)

	newCmd.Flags().StringVarP(&projectParentPath, "path", "P", dir, "The path where the project will be created")
}
