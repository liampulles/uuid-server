package uuid

import "github.com/google/uuid"

// Service provides method for generating UUIDs
type Service interface {
	GenerateVersion4UUID() (string, error)
}

// ServiceImpl implements Service
type ServiceImpl struct {
	newRandomFunc func() (uuid.UUID, error)
}

// Check we implement the service
var _ Service = &ServiceImpl{}

// NewServiceImpl is a constructor
func NewServiceImpl(newRandomFunc func() (uuid.UUID, error)) *ServiceImpl {
	return &ServiceImpl{
		newRandomFunc: newRandomFunc,
	}
}

// GenerateVersion4UUID implements the interface
func (si *ServiceImpl) GenerateVersion4UUID() (string, error) {
	gen, err := si.newRandomFunc()
	if err != nil {
		return "", err
	}
	return gen.String(), nil
}
