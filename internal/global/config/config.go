package config

var conf RuntimeConfig

type RuntimeConfig struct {
	Web WebConfig `yaml:"web"`
}

func Web() WebConfig {
	return conf.Web
}

type WebConfig struct {
	HTTP HTTPConfig `yaml:"http"`
}

type HTTPConfig struct {
	Port int `yaml:"port"`
}
