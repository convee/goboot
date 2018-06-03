package conf

import "github.com/BurntSushi/toml"

type Config struct {
	AppName string `toml:"app_name"`
	Log     Log
	Redis   map[string]RedisConfig
	Mysql   map[string]MysqlConfig
}

type Log struct {
	LogName string `toml:"log_name"`
	LogPath string `toml:"log_path"`
}
type RedisConfig struct {
	Address string
}

type MysqlConfig struct {
	Ip          string
	Port        string
	Username    string
	Password    string
	Database    string
	Charset     string
	MaxIdle     int
	MaxOpen     int
	MaxLifetime int
}

var conf Config

func LoadTomlConfig(path string) {
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		panic(err)
	}
	Set(conf)
}

func Set(config Config) {
	conf = config
}

func Get() Config {
	return conf
}
