package config

import "github.com/spf13/viper"

type (
	Config struct {
		Postgres PostgresConfig
		HTTP     HTTPConfig
		Cache    BannerCacheConfig
	}

	PostgresConfig struct {
		//URI      string
		Name     string
		Username string
		Password string
	}

	HTTPConfig struct {
		Host string
		Port string
	}

	BannerCacheConfig struct {
		Size int
		TTL  int
	}
)

func New(configDir string) (*Config, error) {
	if err := parseConfigFile(configDir); err != nil {
		return nil, err
	}

	var config Config
	if err := unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func parseConfigFile(configDir string) error {
	viper.AddConfigPath(configDir)
	viper.SetConfigName("application")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetConfigName("application")
	return viper.MergeInConfig()
}

func unmarshal(config *Config) error {
	if err := viper.UnmarshalKey("http", &config.HTTP); err != nil {
		return err
	}
	return viper.UnmarshalKey("db", &config.Postgres)
}
