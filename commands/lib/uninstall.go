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

package lib

import (
	"context"
	"fmt"

	"github.com/arduino/arduino-cli/commands"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LibraryUninstall FIXMEDOC
func LibraryUninstall(ctx context.Context, req *rpc.LibraryUninstallRequest, taskCB commands.TaskProgressCB) *status.Status {
	lm := commands.GetLibraryManager(req.GetInstance().GetId())
	ref, err := createLibIndexReference(lm, req)
	if err != nil {
		return status.New(codes.InvalidArgument, err.Error())
	}

	lib := lm.FindByReference(ref)

	if lib == nil {
		taskCB(&rpc.TaskProgress{Message: fmt.Sprintf("Library %s is not installed", req.Name), Completed: true})
	} else {
		taskCB(&rpc.TaskProgress{Name: "Uninstalling " + lib.String()})
		lm.Uninstall(lib)
		taskCB(&rpc.TaskProgress{Completed: true})
	}

	return nil
}
