package config

import (
	"strings"
)

type GoOutput struct {
	Definitions  []string `yaml:"definitions"`
	Dir          string   `yaml:"dir"`
	ImportPrefix string   `yaml:"import_prefix"`
	Paths        string   `yaml:"paths"`
	AnnotateCode *bool    `yaml:"annotate_code"`
}

func (output *GoOutput) buildArgs(cmd *strings.Builder) {
	opts := make([]string, 0)
	if output.ImportPrefix != "" {
		opts = append(opts, "import_prefix="+output.ImportPrefix)
	}
	if output.Paths != "" {
		opts = append(opts, "paths="+output.Paths)
	}
	if output.AnnotateCode != nil {
		opts = append(opts, "annotate_code="+bool2Str(output.AnnotateCode))
	}
	if len(opts) > 0 {
		cmd.WriteString(strings.Join(opts, ","))
		if len(output.Definitions) > 0 {
			cmd.WriteString(",")
		}
	}
	for i, def := range output.Definitions {
		if i > 0 {
			cmd.WriteString(",")
		}
		cmd.WriteString(def)
	}
	cmd.WriteString(":")
	cmd.WriteString(output.Dir)
}

func (output *GoOutput) BuildArgs() []string {
	cmd := strings.Builder{}
	cmd.WriteString("--go_out=")
	output.buildArgs(&cmd)
	return []string{cmd.String()}
}
