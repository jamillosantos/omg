package protodep

import (
	"io"
	"os"

	"github.com/pkg/errors"
)

func Up() error {
	cmd := Cmd("up")

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	err = cmd.Wait()
	if err != nil {
		return err
	}

	ec := cmd.ProcessState.ExitCode()
	if ec != 0 {
		return errors.Errorf("process exited with code: %d", ec)
	}
	return nil
}
