package config

import (
	"github.com/spf13/viper"
)

type ProdConfig struct {
	Database string `mapstructure:"database"`
}

var ProdConfigStore ProdConfig

const (
	prodFilename = "prodconfig.json"
	prodExt      = "json"
)

func (cfg *ProdConfig) FromJson() error {
	viper.AddConfigPath(configpath)
	viper.SetConfigName(prodFilename)
	viper.SetConfigType(prodExt)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		cfg.clear()
		return err
	}

	if err := viper.Unmarshal(&ProdConfigStore); err != nil {
		cfg.clear()
		return err
	}

	return nil
}

func (cfg *ProdConfig) clear() {
	cfg.Database = ""
}
