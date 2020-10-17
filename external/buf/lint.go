package buf

import (
	"io"
	"os"
	"os/exec"
)

func Lint(errFormat string) (int, error) {
	cmd := Cmd("check", "lint", "--error-format", errFormat)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return 0, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}

	err = cmd.Start()
	if err != nil {
		return 0, err
	}

	io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	err = cmd.Wait()
	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode(), nil
	}
	return 0, err
}
