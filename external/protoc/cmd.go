package protoc

import (
	"os/exec"
	"strings"

	"github.com/jamillosantos/omg/internal"
)

func Cmd(args ...string) *exec.Cmd {
	internal.Verbose("executing: protoc", strings.Join(args, " "))
	return exec.Command("protoc", args...)
}
