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
	"os"

	"github.com/arduino/arduino-cli/cli/errorcodes"
	"github.com/arduino/arduino-cli/cli/feedback"
	"github.com/arduino/arduino-cli/commands/board"
	rpc "github.com/arduino/arduino-cli/rpc/commands"
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

// GetAddressAndProtocol returns the Address and Protocol given by command line
// arguments or it will try to figure out the correct one by guessing from the
// available ports.
func (args *PortArguments) GetAddressAndProtocol(instance *rpc.Instance) (string, string) {
	// If no port has been specified just return fields as is, port may be not required
	if args.address == "" {
		return args.address, args.protocol
	}
	// If both address and protocol are specified return them as is
	if args.protocol != "" {
		return args.address, args.protocol
	}

	// Protocol autodetection:
	// - if the port address match only one of the available ports use the protocol of the matching port
	// - if the port address match more than one port exit with an error
	// - in all other cases assume "serial"
	portItems, err := board.List(instance.GetId())
	if err != nil {
		feedback.Errorf("Error getting port list: %v", err)
	} else {
		for _, item := range portItems {
			port := item.GetPort()
			if port.Address != args.address {
				continue
			}
			if args.protocol == "" {
				args.protocol = port.Protocol
			} else {
				feedback.Errorf("Ambiguous port '%s', please specify the protocol.", args.address)
				os.Exit(errorcodes.ErrBadArgument)
			}
		}
	}

	if args.protocol == "" {
		feedback.Print("Using default upload protocol 'serial'")
		args.protocol = "serial"
	}

	return args.address, args.protocol
}
