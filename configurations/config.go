package configurations

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("scope-server-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("/opt/scope")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error reading in config: ", err)
		return err
	}
	return nil
}

func InitTestConfig() error {
	viper.SetConfigName("scope-server-test-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("/opt/scope")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error reading in config: ", err)
		return err
	}
	return nil
}
