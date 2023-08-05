package utils

import (
	"github.com/spf13/viper"
)


type Config struct {
	Name    string 
	ShopURL string 
	Token   string 
}

var (
	config Config
	name string
	shopUrl string
	token string
)
func LoadConfig(path string) (config Config, err error) {
	viper.SetDefault("name","None")
	viper.SetDefault("shopUrl", "None")
	viper.SetDefault("token", "None")
    viper.AddConfigPath("$HOME")
    viper.SetConfigName(".prestctl")
    viper.SetConfigType("yaml")
    viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
        return
    }
	err = viper.Unmarshal(&config)
    return config, err
}
