package config

type OmgConfig struct {
	Src      []string  `yaml:"src"`
	Includes []string  `yaml:"includes"`
	Go       *GoOutput `yaml:"go"`
}

var Verbose bool
var Config OmgConfig

func (c *OmgConfig) BufSources() []string {
	result := make([]string, 0)
	for _, s := range c.Src {
		if len(s) > 0 && s[0] == '!' { // ignores negation of files...
			continue
		}
		result = append(result, s)
	}
	return append(result, c.Includes...)
}
