package config

import (
	"strings"
)

type GoOutput struct {
	Definitions  []string `yaml:"definitions"`
	Dir          string   `yaml:"dir"`
	ImportPrefix string   `yaml:"import_prefix"`
	Plugins      []string `yaml:"plugins"`
}

func (output *GoOutput) BuildArgs() []string {
	cmd := strings.Builder{}
	cmd.WriteString("--go_out=")
	opts := make([]string, 0)
	if output.ImportPrefix != "" {
		opts = append(opts, "import_prefix="+output.ImportPrefix)
	}
	if len(output.Plugins) > 0 {
		opts = append(opts, "plugins="+strings.Join(output.Plugins, ":"))
	}
	for i, def := range output.Definitions {
		if i > 0 {
			cmd.WriteString(",")
		}
		cmd.WriteString(def)
	}
	cmd.WriteString(":")
	cmd.WriteString(output.Dir)
	return []string{cmd.String()}
}
