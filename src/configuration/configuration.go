package configuration

import (
	"time"

	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	DB     Database
	Server Server
}

type Database struct {
	Host               string        `mapstructure:"host"`
	Port               string        `mapstructure:"port"`
	Password           string        `mapstructure:"password"`
	User               string        `mapstructure:"user"`
	Database           string        `mapstructure:"database"`
	SSLMode            bool          `mapstructure:"sslmode"`
	MaxConnections     int           `mapstructure:"max_connections"`
	MaxIdleConnections int           `mapstructure:"max_idle_connections"`
	ConnectionLifeTime time.Duration `mapstructure:"connection_lifetime"`
}

func (d Database) Params() string {
	var params string

	if d.Host != "" {
		params = params + " host=" + d.Host
	}

	if d.Port != "" {
		params = params + " port=" + d.Port
	}

	if d.Password != "" {
		params = params + " password=" + d.Password
	}

	if d.User != "" {
		params = params + " user=" + d.User
	}

	if d.Database != "" {
		params = params + " dbname=" + d.Database
	}

	if !d.SSLMode {
		params = params + " sslmode=disable"
	}

	return params
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

func (s Server) String() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

func New() (*Configuration, error) {
	configuration := new(Configuration)

	database := viper.New()
	database.SetEnvPrefix("database")

	for key, value := range map[string]interface{}{
		"host":                 "localhost",
		"port":                 "5432",
		"user":                 "postgres",
		"password":             "",
		"database":             "postgres",
		"sslmode":              false,
		"max_connections":      10,
		"max_idle_connections": 10,
		"connection_lifetime":  time.Minute,
	} {
		database.BindEnv(key)
		database.SetDefault(key, value)
	}

	if err := database.Unmarshal(&configuration.DB); err != nil {
		return nil, err
	}

	server := viper.New()
	server.SetEnvPrefix("server")

	for key, value := range map[string]interface{}{
		"host": "localhost",
		"port": "3000",
	} {
		server.BindEnv(key)
		server.SetDefault(key, value)
	}

	if err := server.Unmarshal(&configuration.Server); err != nil {
		return nil, err
	}

	return configuration, nil
}
