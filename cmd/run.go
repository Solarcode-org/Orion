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
	"bytes"
	"crypto/sha1"
	"os"
	"path"

	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/builtins"
	"github.com/Solarcode-org/Orion/lib/bytecode"
	"github.com/Solarcode-org/Orion/utils"
	"github.com/Solarcode-org/Orion/utils/astrunner"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectPath string
var build bool

/*
Run an Orion project.

This command runs the project in "-P" (which is,
by default, the current directory) with the
starting file in the "src" folder. If the "-B"
flag is passed the project forcefully rebuilds
without comparing the hashes (useful for recovery).

Usage:

	orion run [-P path] [-B | --build] [flags]

Examples:

	# Run the project in the current directory.
	$ orion run

	# Run the project in directory "foo".
	$ orion run -P foo

	# Forcefully rebuild the project in the current directory.
	$ orion run -B

	# Forcefully rebuild the project in directory "bar".
	$ orion run -P bar -B

Flags:

	-B, --build         Force build
	-h, --help          help for run
	-P, --path string   The path to the project (default $PWD)

Global Flags:

	    --config string   config file (default is $HOME/.orion.yaml)
	-v, --verbose uint8   verbosity level
*/
var runCmd = &cobra.Command{
	Use:   "run [-P path] [-B | --build]",
	Short: "Run an Orion project",
	Long: `Run an Orion project.

This command runs the project in "-P" (which is,
by default, the current directory) with the
starting file in the "src" folder. If the "-B"
flag is passed the project forcefully rebuilds
without comparing the hashes (useful for recovery).`,
	Example: `
	# Run the project in the current directory.
	$ orion run
	
	# Run the project in directory "foo".
	$ orion run -P foo
	
	# Forcefully rebuild the project in the current directory.
	$ orion run -B
	
	# Forcefully rebuild the project in directory "bar".
	$ orion run -P bar -B`,
	PreRun: toggleDebug,
	Args:   cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.Traceln("started `orion run`")

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

		if build {
			log.Traceln("forcefully building")

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

			log.Traceln("ended `orion run` with exit code 0")
			//endregion hashed file
		}

		contents, err := os.ReadFile(mainFilePath)
		utils.CheckErr(err)

		previousHash, err := os.ReadFile(hashFilePath)
		utils.CheckErr(err)

		builtins.MakeFunctions()
		builtins.MakeVariables(verbose)

		astree, parseErrs, err := lib.ParsedFrom(contents)
		if len(parseErrs) > 0 {
			utils.FailParse(parseErrs)
		}
		utils.CheckErr(err)

		instructions, err := bytecode.EncodedSyntaxTree(astree)
		utils.CheckErr(err)

		hash := sha1.New()
		hash.Write(instructions)

		hashed := hash.Sum(nil)

		if bytes.Equal(hashed, previousHash) {
			astrunner.RunAST(astree)

			log.Traceln("ended `orion run` with exit code 0")

			return
		}

		//region hashed file
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

		//region ast file
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

		astrunner.RunAST(astree)

		log.Traceln("ended `orion run` with exit code 0")
	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	dir, err := os.Getwd()
	cobra.CheckErr(err)

	runCmd.Flags().StringVarP(&projectPath, "path", "P", dir, "The path to the project")
	runCmd.Flags().BoolVarP(&build, "build", "B", false, "Force build")
}
