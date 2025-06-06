// This file is part of arduino-cli.
//
// Copyright 2020-2022 ARDUINO SA (http://www.arduino.cc/)
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

package sketch

import (
	"github.com/arduino/go-paths-helper"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/parser"
	"go.bug.st/f"
)

// updateOrAddYamlRootEntry updates or adds a new entry to the root of the yaml file.
// If the value is empty the entry is removed.
func updateOrAddYamlRootEntry(srcPath *paths.Path, key, newValue string) error {
	// First encode the new value as YAML and parse it to an AST
	newValueYaml, err := yaml.Marshal(map[string]string{key: newValue})
	if err != nil {
		return err
	}
	newValueAst := f.Must(parser.ParseBytes(newValueYaml, parser.ParseComments))

	// If the src file does not exist, we can just write the new value
	if !srcPath.Exist() {
		return srcPath.WriteFile(newValueYaml)
	}

	// Read the source YAML file and parse it to an AST
	srcYaml, err := srcPath.ReadFile()
	if err != nil {
		return err
	}
	srcAst, err := parser.ParseBytes(srcYaml, parser.ParseComments)
	if err != nil {
		return err
	}

	// Perform the merge operation
	keyYmlPath, err := yaml.PathString("$")
	if err != nil {
		return err
	}
	if n, _ := keyYmlPath.FilterFile(srcAst); n == nil {
		// In this case the file is empty, we can just write the new value at the bottom
		srcYaml = append(srcYaml, '\n')
		srcYaml = append(srcYaml, newValueYaml...)
		return srcPath.WriteFile(srcYaml)
	}
	if err := keyYmlPath.MergeFromFile(srcAst, newValueAst); err != nil {
		return err
	}

	// Write back the updated YAML
	return srcPath.WriteFile([]byte(srcAst.String()))
}
