package configuration

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"strings"
	"sync"
)

type DefaultConfiguration struct {
	App AppConfig      `json:"app" yaml:"app"`
	DB  DatabaseConfig `json:"db" yaml:"db"`
}

var (
	once     sync.Once
	instance *DefaultConfiguration
)

func LoadConfigs(cfg *DefaultConfiguration) *DefaultConfiguration {
	once.Do(func() {
		v := viper.New()

		// Viper settings
		v.SetConfigName("config")
		v.AddConfigPath(".")
		v.AddConfigPath("$CONFIG_DIR/")

		// Environment variable settings
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
		v.AllowEmptyEnv(true)
		v.AutomaticEnv()

		// Global configuration
		v.SetDefault("shutdownTimeout", defaultShutdownTimeoutSeconds)
		if _, ok := os.LookupEnv("NO_COLOR"); ok {
			v.SetDefault("no_color", true)
		}

		// Database configuration
		_ = v.BindEnv("db.host")
		v.SetDefault("db.port", defaultPostgresPort)
		_ = v.BindEnv("db.user")
		_ = v.BindEnv("db.password")
		_ = v.BindEnv("db.database")

		err := v.ReadInConfig()
		if err != nil {
			panic(errors.Wrap(err, "Cant read configuration file"))
		}

		err = v.Unmarshal(&cfg)
		if err != nil {
			panic(errors.Wrap(err, "Cant unmarshall configuration"))
		}

		instance = cfg
	})

	return instance

}

func GetConfig() *DefaultConfiguration {
	if instance == nil {
		LoadConfigs(&DefaultConfiguration{})
	}

	return instance
}
