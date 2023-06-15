package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

var Config = loadConfig(".")

// envConfig 环境配置
type envConfig struct {
	// 系统配置
	AppName         string `mapstructure:"APP_NAME"`    // 项目名称
	GinMode         string `mapstructure:"GIN_MODE"`    // gin运行模式
	ServerPort      int    `mapstructure:"SERVER_PORT"` // 服务运行端口
	PublicPrefix    string // 资源访问前缀
	UploadDirectory string `mapstructure:"UPLOAD_DIRECTORY"` // 上传文件路径
	Version         string // 系统版本
	Secret          string // 系统加密字符
	StaticPath      string // 静态资源URL路径
	StaticDirectory string `mapstructure:"STATIC_DIRECTORY"` // 静态资源地址
	// Redis配置
	RedisAddr     string `mapstructure:"REDIS_ADDR"`
	RedisAuth     string `mapstructure:"REDIS_AUTH"`
	RedisDb       int    `mapstructure:"REDIS_DB"`
	RedisPoolSize int    // Redis连接池大小
	// MySQL配置
	DatabaseUrl            string `mapstructure:"DATABASE_URL"`
	DbPrefix               string `mapstructure:"DB_PREFIX"`
	DbDefaultStringSize    uint   // 数据库string类型字段的默认长度
	DbMaxIdleConns         int    // 数据库空闲连接池最大值
	DbMaxOpenConns         int    // 数据库连接池最大值
	DbConnMaxLifetimeHours int16  // 连接可复用的最大时间(小时)
	Runtime                string `mapstructure:"RUNTIME"` // 日志目录
	Domain                 string `mapstructure:"DOMAIN"`  // 域名
	VerifyTTL              int
}

// loadConfig 加载配置
func loadConfig(path string) envConfig {
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "", "config file path.")
	flag.Parse()
	if cfgPath == "" {
		viper.AddConfigPath(path)
		viper.SetConfigFile(".env")
	} else {
		viper.SetConfigFile(cfgPath)
	}
	viper.AutomaticEnv()
	config := envConfig{
		GinMode: "debug",
		// 服务运行端口
		ServerPort: 8000,
		// 上传文件路径
		UploadDirectory: "/uploads/",
		// Redis源配置
		RedisAddr:     "localhost:6379",
		RedisAuth:     "",
		RedisDb:       2,
		RedisPoolSize: 100,
		// 数据源配置
		DatabaseUrl:            "username:password@tcp(localhost:3306)/database?charset=utf8mb4&parseTime=True&loc=Local",
		DbPrefix:               "lea_",
		DbDefaultStringSize:    256,
		DbMaxIdleConns:         10,
		DbMaxOpenConns:         100,
		DbConnMaxLifetimeHours: 2,
		// 版本
		Version: "v1.0.0",
		// 系统加密字符
		Secret: "yr62SYDYruCLe8XtdXDwp52uvYkbvoFJ",
		// 静态资源URL路径
		StaticPath: "assets",
		// 静态资源本地路径
		StaticDirectory: "static",
		// 缓存日志目录
		Runtime: "runtime",
		// 访问域名
		Domain: "https://www.baidu.com/",
		// 验证码有效期
		VerifyTTL: 300,
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("loadConfig ReadInConfig err:", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("loadConfig Unmarshal err:", err)
	}
	return config
}
