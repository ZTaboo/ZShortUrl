package model

type Config struct {
	Base Base `yaml:"base"`
}

type Base struct {
	Url string `yaml:"url"`
}
