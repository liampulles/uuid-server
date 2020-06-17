package uuid

import "github.com/google/uuid"

type Service interface {
	GenerateVersion4UUID() (string, error)
}

type ServiceImpl struct {
	newRandomFunc func() (uuid.UUID, error)
}

var _ Service = &ServiceImpl{}

func NewServiceImpl(newRandomFunc func() (uuid.UUID, error)) *ServiceImpl {
	return &ServiceImpl{
		newRandomFunc: newRandomFunc,
	}
}

func (si *ServiceImpl) GenerateVersion4UUID() (string, error) {
	gen, err := si.newRandomFunc()
	if err != nil {
		return "", err
	}
	return gen.String(), nil
}
