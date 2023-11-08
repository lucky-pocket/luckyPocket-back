package config

var conf RuntimeConfig

type RuntimeConfig struct {
	Web  WebConfig  `yaml:"web"`
	Data DataConfig `yaml:"data"`
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

func Data() DataConfig {
	return conf.Data
}

type DataConfig struct {
	Redis RedisConfig `yaml:"redis"`
	Mysql MysqlConfig `yaml:"mysql"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
	Pass string `yaml:"pass"`
	DB   int    `yaml:"db"`
}

type MysqlConfig struct {
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   string `yaml:"db"`
}
