package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-gin/server/core/internal"
	"go-gin/server/global"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config == "" {
			config = internal.ConfigDefaultFile
			fmt.Println("您在使用默认的配置文件")
		}
	} else {
		config = path[0]
	}
	fmt.Println("xx", config)
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: #{err} \n"))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file change:", e.Name)
		if err = v.Unmarshal(&global.GI_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GI_CONFIG); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root位置去找对应迁移位置，保证root有效
	//global.GI_CONFIG.

	return v
}
