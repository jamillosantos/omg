package buf

import (
	"os/exec"
	"strings"

	"github.com/jamillosantos/omg/internal"
)

func Cmd(args ...string) *exec.Cmd {
	internal.Verbose("executing: buf", strings.Join(args, " "))
	return exec.Command("buf", args...)
}
