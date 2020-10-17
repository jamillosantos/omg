package protodep

import (
	"os/exec"
)

func Cmd(args ...string) *exec.Cmd {
	return exec.Command("protodep", args...)
}
