package config

import (
    "github.com/spf13/viper"
)

viper.SetConfigType("yaml")

var yamlExample = []byte(`
username: bvtujo
token: authtoken
`)

viper.ReadConfig(bytes.NewBuffer(yamlExample))

