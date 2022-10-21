package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.SetDefault("RunMode", "debug")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	// 监控配置文件变化并热加载程序，即不重启程序进程就可以加载最新的配置
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = viper.GetString("RunMode")
}

func LoadServer() {
	HTTPPort = viper.GetInt("server.port")
	ReadTimeout = time.Duration(viper.GetInt("ReadTimeOut")) * time.Second
	WriteTimeout = time.Duration(viper.GetInt("WriteTimeOut")) * time.Second
}

func LoadApp() {
	JwtSecret = viper.GetString("jwt.secret")
	PageSize = viper.GetInt("page_size")
}
