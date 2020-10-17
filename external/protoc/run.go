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
	args = append(args, "--descriptor_set_in=/dev/stdin", file)
	return Cmd(args...)
}
