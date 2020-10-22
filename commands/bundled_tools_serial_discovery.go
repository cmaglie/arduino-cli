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

package commands

import (
	"fmt"

	"github.com/arduino/arduino-cli/arduino/cores"
	"github.com/arduino/arduino-cli/arduino/cores/packagemanager"
	"github.com/arduino/arduino-cli/arduino/discovery"
	"github.com/arduino/arduino-cli/arduino/resources"
	semver "go.bug.st/relaxed-semver"
)

var (
	serialDiscoveryVersion = semver.ParseRelaxed("1.1.0")
	serialDiscoveryFlavors = []*cores.Flavor{
		{
			OS: "i686-pc-linux-gnu",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("serial-discovery_v%s_Linux_32bit.tar.bz2", serialDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/serial-discovery_v%s_Linux_32bit.tar.bz2", serialDiscoveryVersion),
				Size:            1589443,
				Checksum:        "SHA-256:e60fa16da2735f80410c636234a405bd1cc9306090cab4e312c4189e38f93f72",
				CachePath:       "tools",
			},
		},
		{
			OS: "x86_64-pc-linux-gnu",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("serial-discovery_v%s_Linux_64bit.tar.bz2", serialDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/serial-discovery_v%s_Linux_64bit.tar.bz2", serialDiscoveryVersion),
				Size:            1611875,
				Checksum:        "SHA-256:6232f852543094e9f73e1994e6888646fdcd24eca15fd4e5bde716a8e23046dc",
				CachePath:       "tools",
			},
		},
		{
			OS: "i686-mingw32",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("serial-discovery_v%s_Windows_32bit.zip", serialDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/serial-discovery_v%s_Windows_32bit.zip", serialDiscoveryVersion),
				Size:            1719070,
				Checksum:        "SHA-256:3efdc744a0ca11c5f9088525eb4363e90e2b6a43a0db23c5c6975a10d739c7cb",
				CachePath:       "tools",
			},
		},
		{
			OS: "x86_64-mingw32",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("serial-discovery_v%s_Windows_64bit.zip", serialDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/serial-discovery_v%s_Windows_64bit.zip", serialDiscoveryVersion),
				Size:            1683799,
				Checksum:        "SHA-256:c6296b92459160f4c0bf7d2e1234cd53fd64f44cb3fa8c3a4b10dd6670466c69",
				CachePath:       "tools",
			},
		},
		{
			OS: "x86_64-apple-darwin",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("serial-discovery_v%s_macOS_64bit.tar.bz2", serialDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/serial-discovery_v%s_macOS_64bit.tar.bz2", serialDiscoveryVersion),
				Size:            1620346,
				Checksum:        "SHA-256:4052a64dd68090726247dea7f03656eae951549df9878362dcedfcef116a9e9f",
				CachePath:       "tools",
			},
		},
		{
			OS: "arm-linux-gnueabihf",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("serial-discovery_v%s_Linux_ARM.tar.bz2", serialDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/serial-discovery_v%s_Linux_ARM.tar.bz2", serialDiscoveryVersion),
				Size:            1511104,
				Checksum:        "SHA-256:fe68fd5abdfebe0f01c48c3eac16d27af46ec2391da87de44571e6ea2dab31e7",
				CachePath:       "tools",
			},
		},
		{
			OS: "arm64-linux-gnueabihf",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("serial-discovery_v%s_Linux_ARM64.tar.bz2", serialDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/serial-discovery_v%s_Linux_ARM64.tar.bz2", serialDiscoveryVersion),
				Size:            1500998,
				Checksum:        "SHA-256:1e6bcb6b7790d38863df15395c96baba239cb56233d7ef2d78bcb2b3efb1bc5d",
				CachePath:       "tools",
			},
		},
	}
)

// WatchListBoards returns a channel that receives events from the bundled discovery tool
func WatchListBoards(pm *packagemanager.PackageManager) (<-chan *discovery.Event, error) {
	t, err := getBuiltinSerialDiscoveryTool(pm)
	if err != nil {
		return nil, err
	}

	if !t.IsInstalled() {
		return nil, fmt.Errorf("missing serial-discovery tool")
	}

	disc, err := discovery.New("serial-discovery", t.InstallDir.Join(t.Tool.Name).String())
	if err != nil {
		return nil, err
	}

	if err = disc.Start(); err != nil {
		return nil, fmt.Errorf("starting discovery: %v", err)
	}

	if err = disc.StartSync(); err != nil {
		return nil, fmt.Errorf("starting sync: %v", err)
	}

	return disc.EventChannel(10), nil
}

func getBuiltinSerialDiscoveryTool(pm *packagemanager.PackageManager) (*cores.ToolRelease, error) {
	builtinPackage := pm.Packages.GetOrCreatePackage("builtin")
	serialDiscoveryTool := builtinPackage.GetOrCreateTool("serial-discovery")
	serialDiscoveryToolRel := serialDiscoveryTool.GetOrCreateRelease(serialDiscoveryVersion)
	serialDiscoveryToolRel.Flavors = serialDiscoveryFlavors
	return pm.Package("builtin").Tool("serial-discovery").Release(serialDiscoveryVersion).Get()
}
