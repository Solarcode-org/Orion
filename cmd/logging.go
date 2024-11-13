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

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// A PlainFormatter is used when the user has turned off logs (by default).
// Here, the logs with level greater than the debug level [github.com/sirupsen/logrus#DebugLevel]
type PlainFormatter struct {
}

// Formats the log entry into a custom form (for now the log message is returned as is)
func (f *PlainFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}

// toggleDebug is used for toggling logs (based on the `-v` or `verbose` flag)
func toggleDebug(*cobra.Command, []string) {
	if verbose > 0 {
		log.Infof("%s logs enabled\n", log.Level(verbose).String())
		log.SetLevel(log.Level(verbose))
		log.SetFormatter(&log.TextFormatter{})
	} else {
		plainFormatter := new(PlainFormatter)
		log.SetFormatter(plainFormatter)
	}
}
