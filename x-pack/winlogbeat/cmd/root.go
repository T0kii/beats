// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"github.com/T0kii/beats/v7/libbeat/cmd"
	winlogbeatCmd "github.com/T0kii/beats/v7/winlogbeat/cmd"

	// Register fields.
	_ "github.com/T0kii/beats/v7/x-pack/libbeat/include"
	_ "github.com/T0kii/beats/v7/x-pack/winlogbeat/include"

	// Enable pipelines.
	_ "github.com/T0kii/beats/v7/x-pack/winlogbeat/module"
)

// Name of this beat.
var Name = winlogbeatCmd.Name

// RootCmd to handle beats cli
var RootCmd *cmd.BeatsRootCmd

func init() {
	settings := winlogbeatCmd.WinlogbeatSettings()
	settings.ElasticLicensed = true
	RootCmd = winlogbeatCmd.Initialize(settings)
	RootCmd.ExportCmd.AddCommand(GenExportPipelineCmd(settings))
}
