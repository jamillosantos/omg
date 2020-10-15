package protoc

import (
	"os/exec"

	"github.com/jamillosantos/omg/config"
)

func Run(bufConf *config.BufConfig, protocConfig *config.ProtocConfig) error {
	args := make([]string, 0)
	for _, v := range bufConf.Build.Roots {
		args = append(args, "-I", v)
	}
	if protocConfig.Go != nil {
		args = append(args, protocConfig.Go.BuildArgs()...)
	}
	cmd := exec.Command("protoc", args...)

	//

	err := cmd.Start()
}
