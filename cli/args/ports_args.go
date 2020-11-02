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

package args

import (
	"net/url"
	"time"

	"github.com/arduino/arduino-cli/arduino/discovery"
	"github.com/arduino/arduino-cli/arduino/sketches"
	"github.com/arduino/arduino-cli/commands"
	rpc "github.com/arduino/arduino-cli/rpc/commands"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// PortArguments is a structure to contain the port arguments result
type PortArguments struct {
	address  string
	protocol string
}

// AddToCommand adds to the cobra command the flags to choose port and protocol
func (args *PortArguments) AddToCommand(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&args.address, "port", "p", "", "Upload port, e.g.: COM10 or /dev/ttyACM0")
	cmd.Flags().StringVarP(&args.protocol, "protocol", "t", "", "Upload port protocol.")
}

// GetPort returns the Port obtained by parsing command line arguments.
// The extra metadata for the ports is obtained using the pluggable discoveries.
func (args *PortArguments) GetPort(instance *rpc.Instance, sketch *sketches.Sketch) (*discovery.Port, error) {
	// FIXME: make a specification on how a port is specified via command line
	address := args.address
	if address == "" && sketch != nil && sketch.Metadata != nil {
		deviceURI, err := url.Parse(sketch.Metadata.CPU.Port)
		if err != nil {
			return nil, errors.Errorf("invalid Device URL format: %s", err)
		}
		if deviceURI.Scheme == "serial" {
			address = deviceURI.Host + deviceURI.Path
		}
	}
	if address == "" {
		return nil, nil
	}

	logrus.WithField("port", address).Tracef("Upload port")

	// FIXME: commands.GetPackageManager must no be used from "cli"
	//        move the whole "Find port metadata" in "commands"

	// Find port metadata
	var fullPort *discovery.Port
	pm := commands.GetPackageManager(instance.GetId())
	if pm == nil {
		return nil, errors.New("invalid instance")
	}
	timeout := time.Now().Add(5 * time.Second)
	//msg := "Waiting for upload port..."
	for time.Now().Before(timeout) {
		currentPorts := pm.GetDiscoveriesManager().FindPort(args.address, args.protocol)
		if len(currentPorts) == 0 {
			time.Sleep(100 * time.Millisecond)
			//outStream.Write([]byte(msg))
			//msg = "."
			continue
		}
		if len(currentPorts) > 1 {
			return nil, errors.Errorf("ambiguous port %s", args.address)
		}
		fullPort = currentPorts[0]
		break
	}
	if fullPort == nil {
		return nil, errors.Errorf("port %s not found", args.address)
	}
	//if msg == "." {
	//	outStream.Write([]byte(" done!\n"))
	//}

	// Protocol autodetection:
	// - if the port address match only one of the available ports use the protocol of the matching port
	// - if the port address match more than one port exit with an error
	// - in all other cases assume "serial"

	return fullPort, nil
}
