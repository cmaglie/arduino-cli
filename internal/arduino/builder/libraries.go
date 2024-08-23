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
	"crypto/md5"
	"encoding/hex"
	"errors"
	"os"
	"slices"
	"strings"
	"time"

	f "github.com/arduino/arduino-cli/internal/algorithms"
	"github.com/arduino/arduino-cli/internal/arduino/builder/cpp"
	"github.com/arduino/arduino-cli/internal/arduino/builder/internal/utils"
	"github.com/arduino/arduino-cli/internal/arduino/libraries"
	"github.com/arduino/arduino-cli/internal/buildcache"
	"github.com/arduino/arduino-cli/internal/i18n"
	"github.com/arduino/go-paths-helper"
	"github.com/arduino/go-properties-orderedmap"
)

// nolint
var (
	FloatAbiCflag = "float-abi"
	FpuCflag      = "fpu"
)

// buildLibraries fixdoc
func (b *Builder) buildLibraries(includesFolders paths.PathList, importedLibraries libraries.List) error {
	includes := f.Map(includesFolders.AsStrings(), cpp.WrapWithHyphenI)
	libs := importedLibraries

	if err := b.librariesBuildPath.MkdirAll(); err != nil {
		return err
	}

	librariesObjectFiles, err := b.compileLibraries(libs, includes)
	if err != nil {
		return err
	}
	b.buildArtifacts.librariesObjectFiles = librariesObjectFiles
	return nil
}

func directoryContainsFile(folder *paths.Path) bool {
	if files, err := folder.ReadDir(); err == nil {
		files.FilterOutDirs()
		return len(files) > 0
	}
	return false
}

func (b *Builder) findExpectedPrecompiledLibFolder(
	library *libraries.Library,
	buildProperties *properties.Map,
) *paths.Path {
	mcu := buildProperties.Get("build.mcu")
	// Add fpu specifications if they exist
	// To do so, resolve recipe.cpp.o.pattern,
	// search for -mfpu=xxx -mfloat-abi=yyy and add to a subfolder
	command, _ := b.prepareCommandForRecipe(buildProperties, "recipe.cpp.o.pattern", true)
	fpuSpecs := ""
	for _, el := range command.GetArgs() {
		if strings.Contains(el, FpuCflag) {
			toAdd := strings.Split(el, "=")
			if len(toAdd) > 1 {
				fpuSpecs += strings.TrimSpace(toAdd[1]) + "-"
				break
			}
		}
	}
	for _, el := range command.GetArgs() {
		if strings.Contains(el, FloatAbiCflag) {
			toAdd := strings.Split(el, "=")
			if len(toAdd) > 1 {
				fpuSpecs += strings.TrimSpace(toAdd[1]) + "-"
				break
			}
		}
	}

	b.logger.Info(i18n.Tr("Library %[1]s has been declared precompiled:", library.Name))

	// Try directory with full fpuSpecs first, if available
	if len(fpuSpecs) > 0 {
		fpuSpecs = strings.TrimRight(fpuSpecs, "-")
		fullPrecompDir := library.SourceDir.Join(mcu).Join(fpuSpecs)
		if fullPrecompDir.Exist() && directoryContainsFile(fullPrecompDir) {
			b.logger.Info(i18n.Tr("Using precompiled library in %[1]s", fullPrecompDir))
			return fullPrecompDir
		}
		b.logger.Info(i18n.Tr(`Precompiled library in "%[1]s" not found`, fullPrecompDir))
	}

	precompDir := library.SourceDir.Join(mcu)
	if precompDir.Exist() && directoryContainsFile(precompDir) {
		b.logger.Info(i18n.Tr("Using precompiled library in %[1]s", precompDir))
		return precompDir
	}
	b.logger.Info(i18n.Tr(`Precompiled library in "%[1]s" not found`, precompDir))
	return nil
}

func (b *Builder) compileLibraries(libraries libraries.List, includes []string) (paths.PathList, error) {
	b.Progress.AddSubSteps(len(libraries))
	defer b.Progress.RemoveSubSteps()

	objectFiles := paths.NewPathList()
	for _, library := range libraries {
		libraryObjectFiles, err := b.compileLibrary(library, includes)
		if err != nil {
			return nil, err
		}
		objectFiles.AddAll(libraryObjectFiles)

		b.Progress.CompleteStep()
	}

	return objectFiles, nil
}

// getCachedLibraryArchiveDirName returns the directory name to be used to store
// the global cached lib.a.
func getCachedLibraryArchiveDirName(fqbn string, optimizationFlags string, libFolder *paths.Path) string {
	fqbnToUnderscore := strings.ReplaceAll(fqbn, ":", "_")
	fqbnToUnderscore = strings.ReplaceAll(fqbnToUnderscore, "=", "_")
	if absCoreFolder, err := libFolder.Abs(); err == nil {
		libFolder = absCoreFolder
	} // silently continue if absolute path can't be detected

	md5Sum := func(data []byte) string {
		md5sumBytes := md5.Sum(data)
		return hex.EncodeToString(md5sumBytes[:])
	}
	hash := md5Sum([]byte(libFolder.String() + optimizationFlags))
	realName := libFolder.Base() + "_" + fqbnToUnderscore + "_" + hash
	if len(realName) > 100 {
		// avoid really long names, simply hash the name again
		realName = md5Sum([]byte(realName))
	}
	return realName
}

func (b *Builder) compileLibrary(library *libraries.Library, includes []string) (paths.PathList, error) {
	if b.logger.Verbose() {
		b.logger.Info(i18n.Tr(`Compiling library "%[1]s"`, library.Name))
	}

	// TODO: libraries with "dot_a_linkage=true" or libraries already providing a
	// precompiled lib.a needs extra boilerplate to be handled. At the moment
	// they are not handled correctly.

	var targetArchivedLibrary *paths.Path
	if b.librariesBuildCachePath != nil {
		archivedLibName := getCachedLibraryArchiveDirName(
			b.buildProperties.Get("build.fqbn"),
			b.buildProperties.Get("compiler.optimization_flags"),
			// TODO: Add included library set here
			library.InstallDir,
		)

		// TODO: if a single library needs a rebuild then all libraries requires a rebuild.
		// This means that the `canUseArchivedLib` must be checked for all libraries
		// in advance in the `compileLibraries` method.

		canUseArchivedLib := func(archivedLib *paths.Path) bool {
			if b.onlyUpdateCompilationDatabase || b.clean {
				return false
			}
			if isOlder, err := utils.DirContentIsOlderThan(library.InstallDir, archivedLib); err != nil || !isOlder {
				// Recreate the archive if ANY of the library files has changed
				return false
			}
			return true
		}

		targetArchivedLibrary = b.librariesBuildCachePath.Join(archivedLibName, "lib.a")
		if canUseArchivedLib(targetArchivedLibrary) {
			// Extend the build cache expiration time
			if _, err := buildcache.New(b.librariesBuildCachePath).GetOrCreate(archivedLibName); errors.Is(err, buildcache.CreateDirErr) {
				return nil, errors.New(i18n.Tr("creating libraries cache folder: %s", err))
			}
			// use archived core
			if b.logger.Verbose() {
				b.logger.Info(i18n.Tr("Using precompiled library: %[1]s", targetArchivedLibrary))
			}
			res := paths.NewPathList()
			res.Add(targetArchivedLibrary)
			return res, nil
		}

		// Create the build cache folder for the core
		if _, err := buildcache.New(b.librariesBuildCachePath).GetOrCreate(archivedLibName); errors.Is(err, buildcache.CreateDirErr) {
			return nil, errors.New(i18n.Tr("creating libraries cache folder: %s", err))
		}
	}

	libraryBuildPath := b.librariesBuildPath.Join(library.DirName)
	if err := libraryBuildPath.MkdirAll(); err != nil {
		return nil, err
	}

	objectFiles := paths.NewPathList()
	archiveFiles := paths.NewPathList()
	if library.Precompiled {
		coreSupportPrecompiled := b.buildProperties.ContainsKey("compiler.libraries.ldflags")
		precompiledPath := b.findExpectedPrecompiledLibFolder(
			library,
			b.buildProperties,
		)

		if !coreSupportPrecompiled {
			b.logger.Info(i18n.Tr("The platform does not support '%[1]s' for precompiled libraries.", "compiler.libraries.ldflags"))
		} else if precompiledPath != nil {
			// Find all libraries in precompiledPath
			libs, err := precompiledPath.ReadDir()
			if err != nil {
				return nil, err
			}

			// Add required LD flags
			libsCmd := library.LDflags + " "
			dynAndStaticLibs := libs.Clone()
			dynAndStaticLibs.FilterSuffix(".a", ".so")
			for _, lib := range dynAndStaticLibs {
				name := strings.TrimSuffix(lib.Base(), lib.Ext())
				if strings.HasPrefix(name, "lib") {
					libsCmd += "-l" + name[3:] + " "
				}
			}

			currLDFlags := b.buildProperties.Get("compiler.libraries.ldflags")
			b.buildProperties.Set("compiler.libraries.ldflags", currLDFlags+" \"-L"+precompiledPath.String()+"\" "+libsCmd+" ")

			// TODO: This codepath is just taken for .a with unusual names that would
			// be ignored by -L / -l methods.
			// Should we force precompiled libraries to start with "lib" ?
			staticLibs := libs.Clone()
			staticLibs.FilterSuffix(".a")
			for _, lib := range staticLibs {
				if !strings.HasPrefix(lib.Base(), "lib") {
					archiveFiles.Add(lib)
				}
			}

			if library.PrecompiledWithSources {
				return archiveFiles, nil
			}
		}
	}

	if library.Layout == libraries.RecursiveLayout {
		libObjectFiles, err := b.compileFiles(
			library.SourceDir, libraryBuildPath,
			true, /** recursive **/
			includes,
		)
		if err != nil {
			return nil, err
		}
		if library.DotALinkage {
			archiveFile, err := b.archiveCompiledFiles(libraryBuildPath.Join(library.DirName+".a"), libObjectFiles)
			if err != nil {
				return nil, err
			}
			archiveFiles.Add(archiveFile)
		} else {
			objectFiles.AddAll(libObjectFiles)
		}
	} else {
		if library.UtilityDir != nil {
			includes = append(includes, cpp.WrapWithHyphenI(library.UtilityDir.String()))
		}
		libObjectFiles, err := b.compileFiles(
			library.SourceDir, libraryBuildPath,
			false, /** recursive **/
			includes,
		)
		if err != nil {
			return nil, err
		}
		objectFiles.AddAll(libObjectFiles)

		if library.UtilityDir != nil {
			utilityBuildPath := libraryBuildPath.Join("utility")
			utilityObjectFiles, err := b.compileFiles(
				library.UtilityDir, utilityBuildPath,
				false, /** recursive **/
				includes,
			)
			if err != nil {
				return nil, err
			}
			objectFiles.AddAll(utilityObjectFiles)
		}
	}

	if len(objectFiles) > 0 {

		// TODO: We must create an object file for each visited directory: this is required because gcc-ar checks
		// if an object file is already in the archive by looking ONLY at the filename WITHOUT the path, so
		// it may happen that a subdir/spi.o inside the archive may be overwritten by a anotherdir/spi.o
		// because thery are both named spi.o.
		// This means that we may need to make multiple .a and cache all of them.

		archiveFile, err := b.archiveCompiledFiles(libraryBuildPath.Join("lib.a"), objectFiles)
		if err != nil {
			return nil, err
		}
		archiveFiles.Add(archiveFile)

		// archive lib.a
		if targetArchivedLibrary != nil && !b.onlyUpdateCompilationDatabase {
			err := archiveFile.CopyTo(targetArchivedLibrary)
			if b.logger.Verbose() {
				if err == nil {
					b.logger.Info(i18n.Tr("Archiving built lib (caching) in: %[1]s", targetArchivedLibrary))
				} else if os.IsNotExist(err) {
					b.logger.Info(i18n.Tr("Unable to cache built lib, please tell %[1]s maintainers to follow %[2]s",
						b.actualPlatform,
						"https://arduino.github.io/arduino-cli/latest/platform-specification/#recipes-to-build-the-corea-archive-file"))
				} else {
					b.logger.Info(i18n.Tr("Error archiving built lib (caching) in %[1]s: %[2]s", targetArchivedLibrary, err))
				}
			}
		}
	}

	return archiveFiles, nil
}

// removeUnusedCompiledLibraries fixdoc
func (b *Builder) removeUnusedCompiledLibraries(importedLibraries libraries.List) error {
	if b.librariesBuildPath.NotExist() {
		return nil
	}

	toLibraryNames := func(libraries []*libraries.Library) []string {
		libraryNames := []string{}
		for _, library := range libraries {
			libraryNames = append(libraryNames, library.DirName)
		}
		return libraryNames
	}

	files, err := b.librariesBuildPath.ReadDir()
	if err != nil {
		return err
	}

	libraryNames := toLibraryNames(importedLibraries)
	for _, file := range files {
		if file.IsDir() {
			if !slices.Contains(libraryNames, file.Base()) {
				if err := file.RemoveAll(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// warnAboutArchIncompatibleLibraries fixdoc
func (b *Builder) warnAboutArchIncompatibleLibraries(importedLibraries libraries.List) {
	archs := []string{b.targetPlatform.Platform.Architecture}
	overrides, _ := b.buildProperties.GetOk("architecture.override_check")
	if overrides != "" {
		archs = append(archs, strings.Split(overrides, ",")...)
	}

	for _, importedLibrary := range importedLibraries {
		if !importedLibrary.SupportsAnyArchitectureIn(archs...) {
			b.logger.Info(
				i18n.Tr("WARNING: library %[1]s claims to run on %[2]s architecture(s) and may be incompatible with your current board which runs on %[3]s architecture(s).",
					importedLibrary.Name,
					strings.Join(importedLibrary.Architectures, ", "),
					strings.Join(archs, ", ")))
		}
	}
}

// printUsedLibraries fixdoc
// TODO here we can completly remove this part as it's duplicated in what we can
// read in the gRPC response
func (b *Builder) printUsedLibraries(importedLibraries libraries.List) {
	if !b.logger.Verbose() || len(importedLibraries) == 0 {
		return
	}

	for _, library := range importedLibraries {
		legacy := ""
		if library.IsLegacy {
			legacy = i18n.Tr("(legacy)")
		}
		if library.Version.String() == "" {
			b.logger.Info(
				i18n.Tr("Using library %[1]s in folder: %[2]s %[3]s",
					library.Name,
					library.InstallDir,
					legacy))
		} else {
			b.logger.Info(
				i18n.Tr("Using library %[1]s at version %[2]s in folder: %[3]s %[4]s",
					library.Name,
					library.Version,
					library.InstallDir,
					legacy))
		}
	}

	// TODO Why is this here?
	time.Sleep(100 * time.Millisecond)
}
