package config

type ProtocConfig struct {
	Go *GoOutput `yaml:"go"`
}

type OmgConfig struct {
	Src []string `yaml:"src"`
}
