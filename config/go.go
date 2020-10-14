package config

type GoOutput struct {
	Definitions    []Def
	DestinationDir string
}

type Def struct {
	From string
	To   string
}
