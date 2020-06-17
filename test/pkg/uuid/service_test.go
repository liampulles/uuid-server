package uuid_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/uuid-server/pkg/uuid"

	googleUUID "github.com/google/uuid"
)

func TestServiceImpl_GenerateVersion4UUID_WhenNewRandomFails_ShouldFail(t *testing.T) {
	// Setup fixture
	serviceImplFixture := uuid.NewServiceImpl(func() (googleUUID.UUID, error) {
		return googleUUID.UUID{}, fmt.Errorf("some error")
	})

	// Exercise SUT
	actual, err := serviceImplFixture.GenerateVersion4UUID()

	// Verify results
	if actual != "" {
		t.Errorf("Expected empty response, got got: %s", actual)
	}
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestServiceImpl_GenerateVersion4UUID_WhenNewRandomSucceeds_ShouldReturnAsString(t *testing.T) {
	// Setup fixture
	expected := "56997ca8-5f7c-4583-afa9-8be557d61007"
	serviceImplFixture := uuid.NewServiceImpl(func() (googleUUID.UUID, error) {
		return googleUUID.Parse(expected)
	})

	// Exercise SUT
	actual, err := serviceImplFixture.GenerateVersion4UUID()

	// Verify results
	if actual != expected {
		t.Errorf("Response mismatch\nActual: %s\nExpected: %s", actual, expected)
	}
	if err != nil {
		t.Errorf("Got an error: %v", err)
	}
}
