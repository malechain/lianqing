package config

import (
	"io/ioutil"
	"log"
	"sync"

	"gopkg.in/yaml.v2"
)

// MysqlConfig mysql配置
type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

// InfuraConfig infura 链接信息
type InfuraConfig struct {
	HTTP string `yaml:"http"`
	Wss  string `yaml:"wss"`
}

// UsdtConfig USDT配置
type UsdtConfig struct {
	Token      string `yaml:"token"`
	PrivateKey string `yaml:"privateKey"`
}

// AddrConfig 生成地址池信息
type AddrConfig struct {
	Password string `yaml:"password"`
}

type CallbackConfig struct {
	Recharge string `yaml:"recharge"`
	Extract  string `yaml:"extract"`
}

type RSAConfig struct {
	Private string `yaml:"private"`
	Public  string `yaml:"public"`
}

// Yaml 整个结构体
type Yaml struct {
	MysqlConfig    `yaml:"mysql"`
	InfuraConfig   `yaml:"infura"`
	UsdtConfig     `yaml:"usdt"`
	AddrConfig     `yaml:"addr"`
	CallbackConfig `yaml:"callback"`
	RSAConfig      `yaml:"rsa"`
}

var (
	conf Yaml
	once sync.Once
)

// GetConfig 返回所有配置信息
func GetConfig() Yaml {
	once.Do(func() {
		yamlFile, err := ioutil.ReadFile("../env.yaml")
		if err != nil {
			log.Fatalln(err)
		}
		if err = yaml.Unmarshal(yamlFile, &conf); err != nil {
			log.Fatalln(err)
		}
	})
	return conf
}
