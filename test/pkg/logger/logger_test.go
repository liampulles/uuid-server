package logger_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/uuid-server/pkg/logger"
)

type loggerSelector func(*logger.ServiceImpl) func(format string, v ...interface{})

func TestServiceImpl_Infof_WhenLogLevelIsInfo_ShouldLog(t *testing.T) {
	// Setup fixture and expectations
	actual := ""
	tests := []struct {
		logLevel      string
		fn            loggerSelector
		formatFixture string
		vFixture      []interface{}
		expected      string
	}{
		// Log level set to info
		{
			"iNfo",
			infof,
			"some %s",
			[]interface{}{"data"},
			"INFO some data",
		},
		{
			"iNfo",
			errorf,
			"some %s",
			[]interface{}{"data"},
			"ERROR some data",
		},
		// Log level set to error
		{
			"eRror",
			infof,
			"some %s",
			[]interface{}{"data"},
			"",
		},
		{
			"eRror",
			errorf,
			"some %s",
			[]interface{}{"data"},
			"ERROR some data",
		},
		// Log level unknown -> default to info
		{
			"dsfae",
			infof,
			"some %s",
			[]interface{}{"data"},
			"INFO some data",
		},
		{
			"dgfhg",
			errorf,
			"some %s",
			[]interface{}{"data"},
			"ERROR some data",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Reset fixture
			actual = ""
			loggerServiceFixture := logger.NewServiceImpl(test.logLevel, func(format string, v ...interface{}) {
				actual = fmt.Sprintf(format, v...)
			})

			// Exercise SUT
			test.fn(loggerServiceFixture)(test.formatFixture, test.vFixture...)

			// Verify results
			if actual != test.expected {
				t.Errorf("Unexpected logged\nActual: %s\nExpected: %s", actual, test.expected)
			}
		})
	}
}

func infof(l *logger.ServiceImpl) func(format string, v ...interface{}) {
	return l.Infof
}

func errorf(l *logger.ServiceImpl) func(format string, v ...interface{}) {
	return l.Errorf
}
