package config

import (
	"flag"
)
func  readConfigPath() string {
	var path string
	flag.StringVar(&path, "configPath", "config/", "Path to config")
	return path
}
