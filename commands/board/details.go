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

package board

import (
	"context"
	"errors"

	"github.com/arduino/arduino-cli/commands"
	rpc "github.com/arduino/arduino-cli/rpc/commands"
)

// Details returns all details for a board including tools and HW identifiers.
// This command basically gather al the information and translates it into the required grpc struct properties
func Details(ctx context.Context, req *rpc.BoardDetailsReq) (*rpc.BoardDetailsResp, error) {
	pm := commands.GetPackageManager(req.GetInstance().GetId())
	if pm == nil {
		return nil, errors.New("invalid instance")
	}

	boardArg := req.GetBoard()
	if boardArg == "" {
		boardArg = req.GetFqbn() // Deprecated: use FQBN for old client.
	}

	fqbn, _, board, err := pm.FindBoard(boardArg, req.GetBoardConfig())
	boardPlatform := board.PlatformRelease
	boardPackage := boardPlatform.Platform.Package

	// TODO: do not ignore *RegisteredBoard result, but output data from it?
	if err != nil {
		return nil, err
	}

	details := &rpc.BoardDetailsResp{}
	details.Name = board.Name()
	details.Fqbn = board.FQBN()
	details.PropertiesId = board.BoardID
	details.Official = fqbn.Package == "arduino"
	details.Version = board.PlatformRelease.Version.String()

	details.Package = &rpc.Package{
		Name:       boardPackage.Name,
		Maintainer: boardPackage.Maintainer,
		WebsiteURL: boardPackage.WebsiteURL,
		Email:      boardPackage.Email,
		Help:       &rpc.Help{Online: boardPackage.Help.Online},
		Url:        boardPackage.URL,
	}

	details.Platform = &rpc.BoardPlatform{
		Architecture:    boardPlatform.Platform.Architecture,
		Category:        boardPlatform.Platform.Category,
		Url:             boardPlatform.Resource.URL,
		ArchiveFileName: boardPlatform.Resource.ArchiveFileName,
		Checksum:        boardPlatform.Resource.Checksum,
		Size:            boardPlatform.Resource.Size,
		Name:            boardPlatform.Platform.Name,
	}

	details.IdentificationPref = []*rpc.IdentificationPref{}
	vids := board.Properties.SubTree("vid")
	pids := board.Properties.SubTree("pid")
	for id, vid := range vids.AsMap() {
		if pid, ok := pids.GetOk(id); ok {
			idPref := rpc.IdentificationPref{UsbID: &rpc.USBID{VID: vid, PID: pid}}
			details.IdentificationPref = append(details.IdentificationPref, &idPref)
		}
	}

	details.ConfigOptions = []*rpc.ConfigOption{}
	options := board.GetConfigOptions()
	for _, option := range options.Keys() {
		configOption := &rpc.ConfigOption{}
		configOption.Option = option
		configOption.OptionLabel = options.Get(option)
		selected, hasSelected := fqbn.Configs.GetOk(option)

		values := board.GetConfigOptionValues(option)
		for i, value := range values.Keys() {
			configValue := &rpc.ConfigValue{}
			if hasSelected && value == selected {
				configValue.Selected = true
			} else if !hasSelected && i == 0 {
				configValue.Selected = true
			}
			configValue.Value = value
			configValue.ValueLabel = values.Get(value)
			configOption.Values = append(configOption.Values, configValue)
		}

		details.ConfigOptions = append(details.ConfigOptions, configOption)
	}

	details.ToolsDependencies = []*rpc.ToolsDependencies{}
	for _, tool := range boardPlatform.Dependencies {
		toolRelease := pm.FindToolDependency(tool)
		var systems []*rpc.Systems
		if toolRelease != nil {
			for _, f := range toolRelease.Flavors {
				systems = append(systems, &rpc.Systems{
					Checksum:        f.Resource.Checksum,
					Size:            f.Resource.Size,
					Host:            f.OS,
					ArchiveFileName: f.Resource.ArchiveFileName,
					Url:             f.Resource.URL,
				})
			}
		}
		details.ToolsDependencies = append(details.ToolsDependencies, &rpc.ToolsDependencies{
			Name:     tool.ToolName,
			Packager: tool.ToolPackager,
			Version:  tool.ToolVersion.String(),
			Systems:  systems,
		})
	}

	return details, nil
}
