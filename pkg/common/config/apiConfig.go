package config

import (
	"shopping/pkg/db/mysqlutil"
	"shopping/pkg/db/redisutil"
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
