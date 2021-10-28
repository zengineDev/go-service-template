package configuration

import (
	"fmt"
	"time"
)

type DatabaseConfig struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	SSLMode  string `json:"ssl_mode" yaml:"ssl_mode"`
}

func (c DatabaseConfig) DSN() string {
	maxConLife := time.Hour
	maxConIdle := time.Minute * 30
	health := time.Minute
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s "+
		"pool_max_conns=10 pool_min_conns=1 pool_max_conn_lifetime=%v pool_max_conn_idle_time=%v pool_health_check_period=%v",
		c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode, maxConLife, maxConIdle, health)
}

func (c DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		c.Host, c.Port, c.User, c.Password, c.Database)
}
