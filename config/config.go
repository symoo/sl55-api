package config

import (
	"log"
	"strings"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

// Config 配置文件
type Config struct {
	Name string
}

// Init 初始化配置
func Init(cfg string) error {

	c := Config {
		Name: cfg,
	}

	if err := c.initConfig(); err != nil {
		return err
	}

	c.watchConfigFile()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("config")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SF55API")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
}

func (c *Config) watchConfigFile() {
	viper.WatchConfig()
	viper.OnConfigChange(func (e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

func (c *Config) initLog() {

}