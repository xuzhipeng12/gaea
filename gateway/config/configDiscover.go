package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Cfg *ConfigMap

type ConfigMap struct {
	ListenPort string `json:"ListenPort" yaml:"listen_port"`

	LogPath string `json:"LogPath" yaml:"log_path"`

	DBUser string `json:"DBUser" yaml:"db_user"`
	DBPwd  string `json:"DBPwd" yaml:"db_pwd"`
	DBHost string `json:"DBHost" yaml:"db_host"`
	DBPort int    `json:"DBPort,string" yaml:"db_port"`
	DBName string `json:"DBName" yaml:"db_name"`
}

func ConfigRead(configFile string) {
	config := new(ConfigMap)
	config.readConfigFromYaml(configFile)
	fmt.Println("Final config: ", config)

	if config.ListenPort == "" {
		panic("init config fail")
	}
	Cfg = config
}

func (this *ConfigMap) readConfigFromYaml(configFile string) {
	if err := configor.Load(this, configFile); err != nil {
		fmt.Println("read config yaml err: " + err.Error())
	}
	fmt.Println("Config from yaml file: ", this)
}
