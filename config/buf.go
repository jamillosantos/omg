package config

// BufConfig represents the settings needed that are specified at the buf.yaml
type BufConfig struct {
	Build BufBuildConfig `yaml:"build"`
}

// BufBuildConfig represents the settings needed that are specified at the buf.yaml
type BufBuildConfig struct {
	Roots []string `yaml:"roots"`
}
