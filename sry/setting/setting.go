package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Port int `mapstructure:"port"`

	Log         *LogConfig   `mapstructure:"log"`
	MysqlConfig *MysqlConfig `mapstructure:"mysql_config"`
	RedisConfig *RedisConfig `mapstructure:"redis_config"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    string `mapstructure:"max_size"`
	MaxAge     string `mapstructure:"max_age"`
	MaxAackups string `mapstructure:"max_ackups"`
}

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func SettingInit() {
	//viper.SetDefault("filedir","./")//设置默认值 默认值是查找配置时优先级最低的
	viper.SetConfigFile("config.yaml") //指定配置文件路径
	//viper.SetConfigName("config.yaml")
	viper.AddConfigPath(".")
	// viper.WriteConfig() //写入add和name
	// viper.SafeWriteConfig()
	//viper.WriteConfigAs()//写入此时定义的配置文件

	//实时监控配置文件
	viper.WatchConfig()

	if err := viper.ReadInConfig(); err != nil {
		// if _, ok := err.(viper.ConfigFileNotFoundError); ok{
		// 	fmt.Println("not fount config file")
		// } else{
		// 	fmt.Println("config file is found, but other fail")
		// }
		panic(err)
	} else {
		fmt.Println("config file exists")
	}

	//把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed: %v", err)
	}
	fmt.Println("Config: ", Conf.Port)

	//回调函数
	//当配置变化之后
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
	})

}
