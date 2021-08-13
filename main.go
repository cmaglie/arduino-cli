// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package main

import (
	"os"

	"github.com/arduino/arduino-cli/cli"
	"github.com/arduino/arduino-cli/cli/errorcodes"
	"github.com/arduino/arduino-cli/configuration"
	"github.com/arduino/arduino-cli/i18n"
	"github.com/arduino/go-properties-orderedmap"
)

func main() {
	// This is only for testing and debugging purposes
	if e := os.Getenv("ARDUINO_CLI_FORCE_OS_SUFFIX"); e != "" {
		properties.SetOSSuffix(e)
	}

	configuration.Settings = configuration.Init(configuration.FindConfigFileInArgsOrWorkingDirectory(os.Args))
	i18n.Init(configuration.Settings.GetString("locale"))
	arduinoCmd := cli.NewCommand()
	if err := arduinoCmd.Execute(); err != nil {
		os.Exit(errorcodes.ErrGeneric)
	}
}
