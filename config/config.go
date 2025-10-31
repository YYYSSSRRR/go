package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

type AppConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Source string `mapstructure:"dsn"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// Conf 创建全局变量Conf,来存这些配置
var Conf *Config

// LoadConfig 把config.yaml中的内容映射到结构体Conf中
func LoadConfig() error {

	//设置配置文件的路径和名称
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	//读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("读取配置文件失败：%v", err)
	}

	//将配置文件的内容解析到Conf变量中
	Conf = &Config{}
	err = viper.Unmarshal(Conf)
	if err != nil {
		return fmt.Errorf("解析配置文件失败：%v", err)
	}
	return nil

}
