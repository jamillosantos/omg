package protoc

import (
	"os/exec"

	"github.com/jamillosantos/omg/config"
)

func Run(omgConfig *config.OmgConfig, file string) *exec.Cmd {
	args := make([]string, 0)
	for _, v := range append(omgConfig.BufSources(), omgConfig.Includes...) {
		args = append(args, "-I", v)
	}
	if omgConfig.Go != nil {
		args = append(args, omgConfig.Go.BuildArgs()...)
	}
	if omgConfig.GoGRPC != nil {
		args = append(args, omgConfig.GoGRPC.BuildArgs()...)
	}
	if omgConfig.GRPCGateway != nil {
		args = append(args, omgConfig.GRPCGateway.BuildArgs()...)
	}
	if omgConfig.OpenAPIV2 != nil {
		args = append(args, omgConfig.OpenAPIV2.BuildArgs()...)
	}
	args = append(args, "--descriptor_set_in=/dev/stdin", file)
	return Cmd(args...)
}
