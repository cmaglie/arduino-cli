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

package packagemanager

import (
	"github.com/arduino/arduino-cli/arduino/cores"
	"github.com/arduino/arduino-cli/arduino/discovery"
	discoverymanager "github.com/arduino/arduino-cli/arduino/discovery/discoveriesmanager"
	"github.com/arduino/go-properties-orderedmap"
	"github.com/pkg/errors"
)

// GetDiscoveriesManager returns the discoveries manager
func (pm *PackageManager) GetDiscoveriesManager() *discoverymanager.DiscoveriesManager {
	return pm.discoveriesManager
}

func (pm *PackageManager) ExtractAllPluggableDiscoveries() error {
	pm.Log.Info("Loading pluggable discoveries")
	for _, platformRelease := range pm.InstalledPlatformReleases() {
		if err := pm.ExtractPluggableDiscoveriesFromPlatformRelease(platformRelease); err != nil {
			return err
		}
	}
	return nil
}

func (pm *PackageManager) ExtractPluggableDiscoveriesFromPlatformRelease(platformRelease *cores.PlatformRelease) error {
	if !platformRelease.IsInstalled() {
		return errors.Errorf("platform %s not installed", platformRelease)
	}

	for id, props := range platformRelease.Properties.SubTree("discovery").FirstLevelOf() {
		discoveryID := platformRelease.String() + ":" + id
		pm.Log.WithField("discoveryID", discoveryID).Debug("instantiating discovery")

		// Resolve discovery command line
		pattern := props.Get("pattern")
		configuration := platformRelease.Properties.Clone()
		configuration.Merge(platformRelease.RuntimeProperties())
		configuration.Merge(props)
		if tools, err := pm.FindToolsRequiredFromPlatformRelease(platformRelease); err == nil {
			for _, tool := range tools {
				configuration.Merge(tool.RuntimeProperties())
			}
		}
		cmd := configuration.ExpandPropsInString(pattern)
		if cmdArgs, err := properties.SplitQuotedString(cmd, `"'`, true); err != nil {
			return err
		} else if d, err := discovery.New(discoveryID, cmdArgs...); err != nil {
			return err
		} else {
			pm.discoveriesManager.Add(d)
		}
	}
	return nil
}
