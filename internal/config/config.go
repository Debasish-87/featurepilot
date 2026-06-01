package config

import "github.com/spf13/viper"

type Config struct {
	AppName string

	HTTPPort string

	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string

	RedisHost string
	RedisPort string
}

func Load() (*Config, error) {

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		AppName: viper.GetString("APP_NAME"),

		HTTPPort: viper.GetString("HTTP_PORT"),

		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresPort:     viper.GetString("POSTGRES_PORT"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),
		PostgresDB:       viper.GetString("POSTGRES_DB"),

		RedisHost: viper.GetString("REDIS_HOST"),
		RedisPort: viper.GetString("REDIS_PORT"),
	}

	return cfg, nil
}
