package config

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	configPath = "./configs/"
	configType = "yaml"
	configName = "config.yaml"
)

type Config interface {
}

type ApplicationConfig struct {
	ServerConfig   ServerConfig
	DataBaseConfig DataBaseConfig
	JwtConfig      JwtConfig
}

type ServerConfig struct {
	Host string `yaml:"addr"`
	Port string `yaml:"port"`
}

type DataBaseConfig struct {
	Drivername string `yaml:"drivername"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Dbname     string `yaml:"dbname"`
	Sslmode    string `yaml:"sslmode"`
}

type JwtConfig struct {
	Accesssecret  string `yaml:"accesssecret"`
	Refreshsecret string `yaml:"refreshsecret"`
}

var config = viper.New()

func InitConfig(log *logrus.Logger) (*ApplicationConfig, error) {
	config.AddConfigPath(configPath)
	config.SetConfigName(configName)
	config.SetConfigType(configType)
	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}
	var sCfg ServerConfig
	err := unmarshalByKey("server", &sCfg)
	if err != nil {
		return nil, err
	}
	var dCfg DataBaseConfig
	err = unmarshalByKey("database", &dCfg)
	if err != nil {
		return nil, err
	}
	var jwtCfg JwtConfig
	err = unmarshalByKey("jwt", &jwtCfg)
	if err != nil {
		return nil, err
	}
	applicationConfig := ApplicationConfig{
		ServerConfig:   sCfg,
		DataBaseConfig: dCfg,
		JwtConfig:      jwtCfg,
	}
	if marshal, err := json.MarshalIndent(&applicationConfig, "", "  "); err == nil {
		log.Println(fmt.Sprintf("Load application configuration %s", marshal))
	}
	return &applicationConfig, nil
}

func unmarshalByKey(key string, cfg Config) error {
	return config.UnmarshalKey(key, &cfg)
}
