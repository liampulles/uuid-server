package config

import (
	goConfig "github.com/liampulles/go-config"
)

// UUIDServerConfig provides configuration for the uuid-server app
type UUIDServerConfig struct {
	Port     int
	LogLevel string
}

// InitUUIDServerConfig populates UUIDServerConfig from a source,
// else uses defaults.
func InitUUIDServerConfig(source goConfig.Source) (*UUIDServerConfig, error) {
	typedSource := goConfig.NewTypedSource(source)
	config := &UUIDServerConfig{
		Port:     8080,
		LogLevel: "INFO",
	}

	if err := goConfig.LoadProperties(typedSource,
		goConfig.IntProp("PORT", &config.Port, false),
		goConfig.StrProp("LOGLEVEL", &config.LogLevel, false),
	); err != nil {
		return nil, err
	}

	return config, nil
}
