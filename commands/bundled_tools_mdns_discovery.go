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
	"github.com/arduino/arduino-cli/arduino/resources"
	semver "go.bug.st/relaxed-semver"
)

var (
	mdnsDiscoveryVersion = semver.ParseRelaxed("1.0.0")
	mdnsDiscoveryFlavors = []*cores.Flavor{
		{
			OS: "i686-pc-linux-gnu",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("mdns-discovery-linux32-v%s.tar.bz2", mdnsDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/mdns-discovery-linux32-v%s.tar.bz2", mdnsDiscoveryVersion),
				Size:            0,
				Checksum:        "SHA-256:0000000000000000000000000000000000000000000000000000000000000000",
				CachePath:       "tools",
			},
		},
		{
			OS: "x86_64-pc-linux-gnu",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("mdns-discovery-linux64-v%s.tar.bz2", mdnsDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/mdns-discovery-linux64-v%s.tar.bz2", mdnsDiscoveryVersion),
				Size:            0,
				Checksum:        "SHA-256:0000000000000000000000000000000000000000000000000000000000000000",
				CachePath:       "tools",
			},
		},
		{
			OS: "i686-mingw32",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("mdns-discovery-windows-v%s.zip", mdnsDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/mdns-discovery-windows-v%s.zip", mdnsDiscoveryVersion),
				Size:            0,
				Checksum:        "SHA-256:0000000000000000000000000000000000000000000000000000000000000000",
				CachePath:       "tools",
			},
		},
		{
			OS: "x86_64-apple-darwin",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("mdns-discovery-macosx-v%s.tar.bz2", mdnsDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/mdns-discovery-macosx-v%s.tar.bz2", mdnsDiscoveryVersion),
				Size:            0,
				Checksum:        "SHA-256:0000000000000000000000000000000000000000000000000000000000000000",
				CachePath:       "tools",
			},
		},
		{
			OS: "arm-linux-gnueabihf",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("mdns-discovery-linuxarm-v%s.tar.bz2", mdnsDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/mdns-discovery-linuxarm-v%s.tar.bz2", mdnsDiscoveryVersion),
				Size:            0,
				Checksum:        "SHA-256:0000000000000000000000000000000000000000000000000000000000000000",
				CachePath:       "tools",
			},
		},
		{
			OS: "arm64-linux-gnueabihf",
			Resource: &resources.DownloadResource{
				ArchiveFileName: fmt.Sprintf("mdns-discovery-linuxarm64-v%s.tar.bz2", mdnsDiscoveryVersion),
				URL:             fmt.Sprintf("https://downloads.arduino.cc/tools/mdns-discovery-linuxarm64-v%s.tar.bz2", mdnsDiscoveryVersion),
				Size:            0,
				Checksum:        "SHA-256:0000000000000000000000000000000000000000000000000000000000000000",
				CachePath:       "tools",
			},
		},
	}
)

func getBuiltinMdnsDiscoveryTool(pm *packagemanager.PackageManager) *cores.ToolRelease {
	builtinPackage := pm.Packages.GetOrCreatePackage("builtin")
	mdnsDiscoveryTool := builtinPackage.GetOrCreateTool("mdns-discovery")
	mdnsDiscoveryToolRel := mdnsDiscoveryTool.GetOrCreateRelease(mdnsDiscoveryVersion)
	mdnsDiscoveryToolRel.Flavors = mdnsDiscoveryFlavors
	return mdnsDiscoveryToolRel
}
