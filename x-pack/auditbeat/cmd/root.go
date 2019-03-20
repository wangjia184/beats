// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"github.com/wangjia184/beats/auditbeat/cmd"
	xpackcmd "github.com/wangjia184/beats/x-pack/libbeat/cmd"

	// Register Auditbeat x-pack modules.
	_ "github.com/wangjia184/beats/x-pack/auditbeat/include"
)

// RootCmd to handle beats CLI.
var RootCmd = cmd.RootCmd

func init() {
	xpackcmd.AddXPack(RootCmd, cmd.Name)
}
