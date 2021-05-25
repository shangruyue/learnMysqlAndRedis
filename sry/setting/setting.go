package setting

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

func SettingInit() {
	//viper.SetDefault("filedir","./")//设置默认值
	viper.SetConfigFile("config.yaml") //指定配置文件路径
	//viper.SetConfigName("config.yaml")
	viper.AddConfigPath(".")
	// viper.WriteConfig() //写入add和name
	// viper.SafeWriteConfig()
	//viper.WriteConfigAs()//写入此时定义的配置文件

	//实时监控配置文件
	viper.WatchConfig()

	//回调函数
	//当配置变化之后
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
	})


	if err := viper.ReadInConfig(); err != nil {
		// if _, ok := err.(viper.ConfigFileNotFoundError); ok{
		// 	fmt.Println("not fount config file")
		// } else{
		// 	fmt.Println("config file is found, but other fail")
		// }
		panic(err)
	} else{
		fmt.Println("config file exists")
	}
	
}
