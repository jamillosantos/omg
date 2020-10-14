package buf

import (
	"bytes"
	"io"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func Ls() ([]string, error) {
	cmd := exec.Command("buf", "ls-files")
	/*
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return nil, err
		}
	*/
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	buff := bytes.NewBuffer(nil)
	_, err = buff.ReadFrom(stdout)
	if err != nil {
		return nil, err
	}

	r := make([]string, 0)
	for {
		l, err := buff.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		r = append(r, strings.TrimSpace(l))
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	ec := cmd.ProcessState.ExitCode()
	if ec != 0 {
		return nil, errors.Errorf("process exited with code: %d", ec)
	}
	return r, nil
}
