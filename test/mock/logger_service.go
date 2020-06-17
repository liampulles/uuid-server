package mock

import (
	"fmt"

	"github.com/liampulles/uuid-server/pkg/logger"
)

type MockLoggerService struct {
	InfoLogged  []string
	ErrorLogged []string
}

var _ logger.Service = &MockLoggerService{}

func (m *MockLoggerService) Infof(format string, v ...interface{}) {
	m.InfoLogged = append(m.InfoLogged, fmt.Sprintf(format, v...))
}

func (m *MockLoggerService) Errorf(format string, v ...interface{}) {
	m.ErrorLogged = append(m.ErrorLogged, fmt.Sprintf(format, v...))
}
