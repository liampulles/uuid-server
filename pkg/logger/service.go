package logger

import (
	"strings"
)

type formatFunc func(format string, v ...interface{})

// Service provides common logging mechanisms
type Service interface {
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

// ServiceImpl implements Service
type ServiceImpl struct {
	infoFImpl  formatFunc
	errorFImpl formatFunc
}

// Check we implement the interface
var _ Service = &ServiceImpl{}

// NewServiceImpl is a constructor
func NewServiceImpl(logLevel string, actualLogger formatFunc) *ServiceImpl {
	si := &ServiceImpl{
		infoFImpl:  emptyFormatFunc,
		errorFImpl: emptyFormatFunc,
	}
	switch strings.ToUpper(logLevel) {
	default:
		// Default to INFO
		fallthrough
	case "INFO":
		si.infoFImpl = templateFormatFunc("INFO", actualLogger)
		fallthrough
	case "ERROR":
		si.errorFImpl = templateFormatFunc("ERROR", actualLogger)
	}
	return si
}

// Infof implements the interface
func (si *ServiceImpl) Infof(format string, v ...interface{}) {
	si.infoFImpl(format, v...)
}

// Errorf implements the interface
func (si *ServiceImpl) Errorf(format string, v ...interface{}) {
	si.errorFImpl(format, v...)
}

func emptyFormatFunc(format string, v ...interface{}) {
	return
}

func templateFormatFunc(prefix string, actualLogger formatFunc) formatFunc {
	return func(format string, v ...interface{}) {
		actualLogger(prefix+" "+format, v...)
	}
}
