package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DevConfig struct {
	LocalPort string `mapstructure:"localport"`
	Database  struct {
		Heroku struct {
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Dbname   string `mapstructure:"dbname"`
		} `mapstructure:"heroku"`
		Local struct {
			User     string `mapstructure:"user"`
			Password string `mapstructure:"password"`
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Dbname   string `mapstructure:"dbname"`
		} `mapstructure:"local"`
	} `mapstructure:"database"`
	Logs struct {
		OutputStdout bool   `mapstructure:"output to stdout"`
		Filename     string `mapstructure:"filename"`
	} `mapstructure:"logs"`
	Sessions struct {
		Name string `mapstructure:"session name"`
	} `mapstructure:"sessions"`
}

var DevConfigStore DevConfig

const (
	devFilename = "devconfig.json"
	devExt      = "json"
)

const configpath = "config/"

func (cfg *DevConfig) FromJson() error {
	viper.AddConfigPath(configpath)
	viper.SetConfigName(devFilename)
	viper.SetConfigType(devExt)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error in config reading")
		cfg.clear()
		return err
	}

	if err := viper.Unmarshal(&DevConfigStore); err != nil {
		fmt.Println("error in config reading")
		cfg.clear()
		return err
	}

	return nil
}

func (cfg *DevConfig) clear() {
	cfg.LocalPort = ""

	cfg.Database.Heroku.User = ""
	cfg.Database.Heroku.Password = ""
	cfg.Database.Heroku.Host = ""
	cfg.Database.Heroku.Password = ""
	cfg.Database.Heroku.Dbname = ""

	cfg.Database.Local.User = ""
	cfg.Database.Local.Password = ""
	cfg.Database.Local.Host = ""
	cfg.Database.Local.Password = ""
	cfg.Database.Local.Dbname = ""

	cfg.Logs.OutputStdout = false
	cfg.Logs.Filename = ""
}
