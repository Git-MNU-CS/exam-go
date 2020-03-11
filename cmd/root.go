package cmd

import (
	"fmt"

	"os"

	"github.com/goexam/internal/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "exam",
	Short: "Short Ins",
	Long:  `Long Ins`,
}

func init() {
	// 读取配置
	cobra.OnInitialize(initConfig)
	// 添加配置参数
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
}

// Execute is ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile == "" {
		return
	}
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

type (
	// LoggingOps is ..
	LoggingOps struct {
		Mode string `mapstructure:"mode" yaml:"mode"`
	}
	// ServerOps is ..
	ServerOps struct {
		TServer TeacherOps `mapstructure:"t_server" yaml:"t_server"`
	}
	// ApplicationOps is ...
	ApplicationOps struct {
		Logging  LoggingOps               `mapstructure:"logging" yaml:"logging"`
		Server   ServerOps                `mapstructure:"server" yaml:"server"`
		Database services.DatabaseOptions `mapstructure:"database" yaml:"database"`
	}
)

// Load 使用viper加载配置文件
func (opts *ApplicationOps) Load() {
	err := viper.Unmarshal(opts)
	if err != nil {
		// 加入log组件, 改用log记录
		fmt.Printf("failed to parse config file: %s", err)
	}
}

func loadApplocationOps() *ApplicationOps {
	opts := &ApplicationOps{}
	opts.Load()
	return opts
}

// StaticOptions 静态文件配置
type StaticOptions struct {
	Path string `yaml:"path" mapstructure:"path"`
	Root string `yaml:"root" mapstructure:"root"`
}

// ServerOptions 服务器配置
type ServerOptions struct {
	Host   string        `yaml:"host" mapstructure:"path"`
	Port   int           `yaml:"port" mapstructure:"port"`
	Views  string        `yaml:"views" mapstructure:"views"`
	Static StaticOptions `yaml:"static" mapstructure:"static"`
}
