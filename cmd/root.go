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

// Package cmd contains the code for the Orion CLI.
// It contains the following commands:
//   - ior
//   - manpages
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfgFile represents the configuration file for this CLI.
var cfgFile string

// verbose represents the verbosity level (0-6)
var verbose uint8

/*
Orion is the CLI tool for the Orion programming language.

The Orion programming language was invented in 2024 by
Arnab Phukan. It was made to combine the speed of Go,
security of Rust and efficiency of Python. It is a
compiled language.

Usage:

	orion [command]

Available Commands:

	completion  Generate the autocompletion script for the specified shell
	encode      Encode a parsed Orion file into bytecode
	help        Help about any command
	ior         Run an individual Orion source file
	manpages    Generate manpages for Orion
	new         Make a new Orion project
	run         Run an Orion project

Flags:

	    --config string   config file (default is $HOME/.orion.yaml)
	-h, --help            help for orion
	-v, --verbose uint8   verbosity level

Use "orion [command] --help" for more information about a command.
*/
var RootCmd = &cobra.Command{
	Use:   "orion",
	Short: "Orion is the CLI tool for the Orion programming language",
	Long: `Orion is the CLI tool for the Orion programming language.

The Orion programming language was invented in 2024 by
Arnab Phukan. It was made to combine the speed of Go,
security of Rust and efficiency of Python. It is a
compiled language.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.orion.yaml)")
	RootCmd.PersistentFlags().Uint8VarP(&verbose, "verbose", "v", 0, "verbosity level")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".orion" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".orion")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
