package inits

import (
	"blog/global"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

func init() {

	var config string
	if configEnv := os.Getenv("BA_CONFIG"); configEnv == "" {
		config = "config.yaml"
	} else {
		config = configEnv
		fmt.Printf("您正在使用BA_CONFIG环境变量,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			panic(err)
		}
	})

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(err)
	}
}
