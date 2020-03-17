package goexam

import (
	"time"

	"github.com/jinzhu/gorm"
	// gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DatabaseOptions 创建数据库的选项
type DatabaseOptions struct {
	Driver    string `yaml:"driver" mapstructure:"driver"`
	Dsn       string `yaml:"dsn" mapstructure:"dsn"`
	KeepAlive int    `yaml:"keep_alive" mapstructure:"keep_alive"`
	MaxIdles  int    `yaml:"max_idles" mapstructure:"max_idles"`
	MaxOpens  int    `yaml:"max_opens" mapstructure:"max_opens"`
}

// DB is DB
type DB struct {
	*gorm.DB

	// ticker 用于keep alive的定时器
	ticker *time.Ticker
}
