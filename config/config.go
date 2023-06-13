package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Name    string `mapstructure:"name" required:"true"`
		Version string `mapstructure:"version" required:"true"`
		Host    string `mapstructure:"host" required:"true"`
		Port    string `mapstructure:"port" required:"true"`
		AppMode string `mapstructure:"app_mode" required:"true"`
		DbMode  string `mapstructure:"db_mode" required:"true"`
	} `mapstructure:"server"`

	Db struct {
		Sql struct {
			Driver   string `mapstructure:"driver" required:"false"`
			Host     string `mapstructure:"host" required:"false"`
			Name     string `mapstructure:"name" required:"false"`
			Port     string `mapstructure:"port" required:"false"`
			User     string `mapstructure:"user" required:"false"`
			Password string `mapstructure:"password" required:"false"`
		} `mapstructure:"sql"`

		Mongodb struct {
			Uri string `mapstructure:"uri" required:"false"`
			Db  string `mapstructure:"db" required:"false"`
		} `mapstructure:"mongodb"`
	} `mapstructure:"db"`
}

func NewConfig(appMode string) *Config {
	cfg := &Config{}
	viper.AddConfigPath("./config")
	viper.SetConfigName(appMode)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Panic(err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		logrus.Panic(err)
	}
	return cfg
}
