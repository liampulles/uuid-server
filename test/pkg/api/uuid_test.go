package api_test

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/liampulles/uuid-server/pkg/api"
	"github.com/liampulles/uuid-server/test/mock"
)

func TestUUIDHandler_ServeHTTP_WhenUUIDServiceFails_ShouldReturnInternalServiceErrorAndLogIt(t *testing.T) {
	// Setup fixture
	mockLogger := &mock.LoggerService{}
	uuidHandler := api.NewUUIDHandler(
		mockLogger,
		&mock.UUIDService{
			MockErr: fmt.Errorf("some error"),
		})
	responseWriter := &mockResponseWriter{
		setWriter: bytes.NewBufferString(""),
	}

	// Setup expectations
	expectedCode := 500
	expectedBytes := "Encountered an error: some error"
	expectedErrorLog := []string{"UUID Service error: some error"}

	// Exercise SUT
	uuidHandler.ServeHTTP(responseWriter, nil)

	// Verify results
	if responseWriter.setStatusCode != expectedCode {
		t.Errorf("Status code mismatch\nActual: %d\nExpected: %d", responseWriter.setStatusCode, expectedCode)
	}
	if responseWriter.setWriter.String() != expectedBytes {
		t.Errorf("Written bytes mismatch\nActual: %s\nExpected: %s", responseWriter.setWriter.String(), expectedBytes)
	}
	if !reflect.DeepEqual(mockLogger.ErrorLogged, expectedErrorLog) {
		t.Errorf("Error log mismatch\nActual: %#v\nExpected: %#v", mockLogger.ErrorLogged, expectedErrorLog)
	}
}

func TestUUIDHandler_ServeHTTP_WhenUUIDServiceFailsAndResponseWriterFails_ShouldLogBoth(t *testing.T) {
	// Setup fixture
	mockLogger := &mock.LoggerService{}
	uuidHandler := api.NewUUIDHandler(
		mockLogger,
		&mock.UUIDService{
			MockErr: fmt.Errorf("some error"),
		})
	responseWriter := &mockResponseWriter{
		writeErr: fmt.Errorf("writing error"),
	}

	// Setup expectations
	expectedCode := 500
	expectedErrorLog := []string{
		"Could not write error to response: writing error",
		"UUID Service error: some error",
	}

	// Exercise SUT
	uuidHandler.ServeHTTP(responseWriter, nil)

	// Verify results
	if responseWriter.setStatusCode != expectedCode {
		t.Errorf("Status code mismatch\nActual: %d\nExpected: %d", responseWriter.setStatusCode, expectedCode)
	}
	if !reflect.DeepEqual(mockLogger.ErrorLogged, expectedErrorLog) {
		t.Errorf("Error log mismatch\nActual: %#v\nExpected: %#v", mockLogger.ErrorLogged, expectedErrorLog)
	}
}

func TestUUIDHandler_ServeHTTP_WhenUUIDServicePassesButResponseWriterFails_ShouldSetStatusCodeButLogError(t *testing.T) {
	// Setup fixture
	mockLogger := &mock.LoggerService{}
	uuidHandler := api.NewUUIDHandler(
		mockLogger,
		&mock.UUIDService{
			MockResp: "some uuid",
		})
	responseWriter := &mockResponseWriter{
		writeErr: fmt.Errorf("writing error"),
	}

	// Setup expectations
	expectedCode := 200
	expectedErrorLog := []string{
		"Could not write response: writing error",
	}

	// Exercise SUT
	uuidHandler.ServeHTTP(responseWriter, nil)

	// Verify results
	if responseWriter.setStatusCode != expectedCode {
		t.Errorf("Status code mismatch\nActual: %d\nExpected: %d", responseWriter.setStatusCode, expectedCode)
	}
	if !reflect.DeepEqual(mockLogger.ErrorLogged, expectedErrorLog) {
		t.Errorf("Error log mismatch\nActual: %#v\nExpected: %#v", mockLogger.ErrorLogged, expectedErrorLog)
	}
}

func TestUUIDHandler_ServeHTTP_WhenUUIDServicePasses_ShouldSetStatusCodeOkAndWriteResponse(t *testing.T) {
	// Setup fixture
	mockLogger := &mock.LoggerService{}
	uuidHandler := api.NewUUIDHandler(
		mockLogger,
		&mock.UUIDService{
			MockResp: "some uuid",
		})
	responseWriter := &mockResponseWriter{
		setWriter: bytes.NewBufferString(""),
	}

	// Setup expectations
	expectedCode := 200
	expectedResponse := "some uuid"

	// Exercise SUT
	uuidHandler.ServeHTTP(responseWriter, nil)

	// Verify results
	if responseWriter.setStatusCode != expectedCode {
		t.Errorf("Status code mismatch\nActual: %d\nExpected: %d", responseWriter.setStatusCode, expectedCode)
	}
	if responseWriter.setWriter.String() != expectedResponse {
		t.Errorf("Response mismatch\nActual: %s\nExpected: %s", responseWriter.setWriter.String(), expectedResponse)
	}
}

type mockResponseWriter struct {
	writeErr      error
	setStatusCode int
	setWriter     *bytes.Buffer
}

func (m *mockResponseWriter) Header() http.Header {
	return nil
}

func (m *mockResponseWriter) Write(b []byte) (int, error) {
	if m.writeErr != nil {
		return -1, m.writeErr
	}
	return m.setWriter.Write(b)
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.setStatusCode = statusCode
}
