package config

import (
	"github.com/spf13/viper"
)

// Config 结构体表示配置项
type Config struct {
	Server struct {
		Host string `yaml:"Host"`
		Port int    `yaml:"Port"`
	} `yaml:"Server"`

	MySQL struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		Database string `yaml:"Database"`
	} `yaml:"MySQL"`

	Redis struct {
		Host     string `yaml:"Host"`
		Port     int    `yaml:"Port"`
		Database int    `yaml:"Database"`
	} `yaml:"Redis"`
}

// LoadConfig 加载配置文件
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	Config := &Config{}
	err = viper.Unmarshal(Config)
	if err != nil {
		return nil, err
	}
	return Config, nil
}
