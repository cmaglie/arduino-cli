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

package builder

import (
	"github.com/arduino/arduino-cli/arduino/builder"
	"github.com/arduino/arduino-cli/legacy/builder/types"
	"github.com/pkg/errors"
)

type ContainerSetupHardwareToolsLibsSketchAndProps struct{}

func (s *ContainerSetupHardwareToolsLibsSketchAndProps) Run(ctx *types.Context) error {
	sketchBuildPath, librariesBuildPath, coreBuildPath,
		warningsLevel, err := AddAdditionalEntriesToContext(ctx.BuildPath, ctx.WarningsLevel)
	if err != nil {
		return errors.WithStack(err)
	}
	ctx.SketchBuildPath = sketchBuildPath
	ctx.LibrariesBuildPath = librariesBuildPath
	ctx.CoreBuildPath = coreBuildPath
	ctx.WarningsLevel = warningsLevel

	if ctx.BuildPath.Canonical().EqualsTo(ctx.Sketch.FullPath.Canonical()) {
		return errors.New(tr("Sketch cannot be located in build path. Please specify a different build path"))
	}

	lm, libsResolver, verboseOut, err := builder.LibrariesLoader(
		ctx.UseCachedLibrariesResolution, ctx.LibrariesManager,
		ctx.BuiltInLibrariesDirs, ctx.LibraryDirs, ctx.OtherLibrariesDirs,
		ctx.ActualPlatform, ctx.TargetPlatform,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	ctx.LibrariesManager = lm
	ctx.LibrariesResolver = libsResolver
	if ctx.Verbose {
		ctx.Warn(string(verboseOut))
	}
	return nil
}
