package buf

import (
	"io"
	"os"
	"os/exec"
)

type BuildErrorFormat = string

const (
	BuildErrorFormatText BuildErrorFormat = "text"
	BuildErrorFormatJSON BuildErrorFormat = "json"
)

type BuildRequest struct {
	AsFileDescriptorSet bool
	ErrorFormat         BuildErrorFormat
	ExcludeImports      bool
	ExcludeSourceInfo   bool
	Source              []string
}

func BuildCmd(req *BuildRequest) *exec.Cmd {
	args := []string{"image", "build"}

	if req.AsFileDescriptorSet {
		args = append(args, "--as-file-descriptor-set")
	}
	if req.ErrorFormat != "" {
		args = append(args, "--error-format", req.ErrorFormat)
	}
	if req.ExcludeImports {
		args = append(args, "--exclude-imports")
	}
	if req.ExcludeSourceInfo {
		args = append(args, "--exclude-source-info")
	}
	for _, source := range req.Source {
		args = append(args, "--source", source)
	}

	args = append(args, "-o", "-")

	return Cmd(args...)
}

func Build(req *BuildRequest) (int, error) {
	cmd := BuildCmd(req)

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
