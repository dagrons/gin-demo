package settings

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	fmt.Print("-1")
	pflag.String("conf_dir", "conf", "configuration location")
	pflag.String("log_dir", "logs", "logs location")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(viper.GetString("conf_dir"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("read config error, err=%v", err))
	}
}
