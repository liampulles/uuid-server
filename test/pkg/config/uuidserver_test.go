package config_test

import (
	"reflect"
	"testing"

	goConfig "github.com/liampulles/go-config"
	"github.com/liampulles/uuid-server/pkg/config"
)

func TestInitUuidServerConfig_WhenParametersUnset_ShouldUseDefault(t *testing.T) {
	// Setup fixture
	source := goConfig.MapSource(map[string]string{})

	// Setup expectations
	expected := &config.UUIDServerConfig{
		Port:     8080,
		LogLevel: "INFO",
	}

	// Exercise SUT
	actual, err := config.InitUUIDServerConfig(source)

	// Verify results
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response mismatch\nActual: %v\nExpected: %v", actual, expected)
	}
}

func TestInitUuidServerConfig_WhenParametersSet_ShouldUseParameters(t *testing.T) {
	// Setup fixture
	source := goConfig.MapSource(map[string]string{
		"PORT":     "9001",
		"LOGLEVEL": "ERROR",
	})

	// Setup expectations
	expected := &config.UUIDServerConfig{
		Port:     9001,
		LogLevel: "ERROR",
	}

	// Exercise SUT
	actual, err := config.InitUUIDServerConfig(source)

	// Verify results
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response mismatch\nActual: %v\nExpected: %v", actual, expected)
	}
}

func TestInitUuidServerConfig_WhenParameterTypeMismatch_ShouldFail(t *testing.T) {
	// Setup fixture
	source := goConfig.MapSource(map[string]string{
		"PORT":     "not an int",
		"LOGLEVEL": "ERROR",
	})

	// Exercise SUT
	actual, err := config.InitUUIDServerConfig(source)

	// Verify results
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
	if actual != nil {
		t.Errorf("Unexpected response: %v", actual)
	}
}
