package basepkg

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadByViper() {
	cfg := "config.json"
	v := viper.New()
	v.SetConfigFile(cfg)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println("read config fail", err)
	}
	fmt.Printf("%v", v.Get("nickname"))
}
