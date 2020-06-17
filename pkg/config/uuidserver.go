package config

import (
	goConfig "github.com/liampulles/go-config"
)

type UuidServerConfig struct {
	Port     string
	LogLevel string
}

func InitUuidServerConfig(source goConfig.Source) (*UuidServerConfig, error) {
	typedSource := goConfig.NewTypedSource(source)
	config := &UuidServerConfig{
		Port:     "8080",
		LogLevel: "INFO",
	}

	if err := goConfig.LoadProperties(typedSource,
		goConfig.StrProp("PORT", &config.Port, false),
		goConfig.StrProp("LOGLEVEL", &config.LogLevel, false),
	); err != nil {
		return nil, err
	}

	return config, nil
}
