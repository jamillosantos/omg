package config

import (
	"strings"
)

type GoGRPCOutput struct {
	GoOutput                    `yaml:",inline"`
	RequireUnimplementedServers *bool `yaml:"require_unimplemented_servers"`
}

func (output *GoGRPCOutput) BuildArgs() []string {
	cmd := strings.Builder{}
	cmd.WriteString("--go-grpc_out=")
	if output.RequireUnimplementedServers != nil {
		cmd.WriteString("require_unimplemented_servers=")
		if *output.RequireUnimplementedServers {
			cmd.WriteString("true")
		} else {
			cmd.WriteString("false")
		}
	}
	cmdArgs := strings.Builder{}
	output.GoOutput.buildArgs(&cmdArgs)
	if cmdArgs.Len() > 0 {
		cmd.WriteString(",")
		cmd.WriteString(cmdArgs.String())
	}
	return []string{cmd.String()}
}
