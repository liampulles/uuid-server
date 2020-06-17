package mock

import (
	"fmt"

	"github.com/liampulles/uuid-server/pkg/logger"
)

// LoggerService mocks logger.Service
type LoggerService struct {
	InfoLogged  []string
	ErrorLogged []string
}

// Check we implement the interface
var _ logger.Service = &LoggerService{}

// Infof implements the interface
func (m *LoggerService) Infof(format string, v ...interface{}) {
	m.InfoLogged = append(m.InfoLogged, fmt.Sprintf(format, v...))
}

// Errorf implements the interface
func (m *LoggerService) Errorf(format string, v ...interface{}) {
	m.ErrorLogged = append(m.ErrorLogged, fmt.Sprintf(format, v...))
}
