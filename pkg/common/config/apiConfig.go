package config

import (
	"base_frame/pkg/db/mysqlutil"
	"base_frame/pkg/db/redisutil"
)

type Config struct {
	Name            string `mapstructure:"name"`
	Mode            string `mapstructure:"mode"`
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	DefaultPassword string `mapstructure:"default_password"`

	Mysql *mysqlutil.Config `mapstructure:"mysql"`
	Redis *redisutil.Config `mapstructure:"redis"`
}
