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

package discoveriesmanager

import (
	"github.com/arduino/arduino-cli/arduino/discovery"
	"github.com/pkg/errors"
)

// DiscoveriesManager is required to handle multiple pluggable-discovery that
// may be shared across platforms
type DiscoveriesManager struct {
	discoveries map[string]*discovery.PluggableDiscovery
}

// New creates a new DiscoveriesManager
func New() *DiscoveriesManager {
	return &DiscoveriesManager{
		discoveries: map[string]*discovery.PluggableDiscovery{},
	}
}

// Add adds a discovery to the list of managed discoveries
func (dm *DiscoveriesManager) Add(disc *discovery.PluggableDiscovery) error {
	id := disc.GetID()
	if _, has := dm.discoveries[id]; has {
		return errors.Errorf("pluggable discovery already added: %s", id)
	}
	dm.discoveries[id] = disc
	disc.Start()
	return disc.StartSync()
}

// ListPorts return the current list of ports detected from all discoveries
func (dm *DiscoveriesManager) ListPorts() ([]*discovery.Port, error) {
	res := []*discovery.Port{}
	for _, disc := range dm.discoveries {
		for _, port := range disc.ListSync() {
			res = append(res, port)
		}
	}
	return res, nil
}
