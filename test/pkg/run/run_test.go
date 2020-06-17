package run_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/liampulles/go-config"
	"github.com/liampulles/uuid-server/pkg/run"
)

func TestRun_WhenConfigFails_ShouldReturnOne(t *testing.T) {
	// Setup fixture
	source := config.MapSource(map[string]string{
		"PORT": "not an int",
	})
	expected := 1

	// Exercise SUT
	actual := run.Run(source, nil)

	// Verify results
	if actual != expected {
		t.Errorf("Unexpected response\nActual: %d\nExpected: %d", actual, expected)
	}
}

func TestRun_WhenServeFails_ShouldReturnTwo(t *testing.T) {
	// Setup fixture
	source := config.MapSource(map[string]string{})
	serveFunc := func(string, http.Handler) error {
		return fmt.Errorf("some error")
	}
	expected := 2

	// Exercise SUT
	actual := run.Run(source, serveFunc)

	// Verify results
	if actual != expected {
		t.Errorf("Unexpected response\nActual: %d\nExpected: %d", actual, expected)
	}
}

func TestRun_WhenServeSucceeds_ShouldReturnZero(t *testing.T) {
	// Setup fixture
	source := config.MapSource(map[string]string{})
	serveFunc := func(string, http.Handler) error {
		return nil
	}
	expected := 0

	// Exercise SUT
	actual := run.Run(source, serveFunc)

	// Verify results
	if actual != expected {
		t.Errorf("Unexpected response\nActual: %d\nExpected: %d", actual, expected)
	}
}
