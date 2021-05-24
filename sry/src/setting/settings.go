package setting

import (
	"fmt"

	"github.com/spf13/viper"
)

//使用viper进行 配置文件的读取
func Init() {
	viper.SetConfigFile("./setting/config.yaml") //指定配置文件路径
	viper.SetConfigName("config.yaml")
	viper.AddConfigPath("/sry/src/setting/")


	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}