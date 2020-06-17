package uuid

import "github.com/google/uuid"

type Service interface {
	GenerateVersion4UUID() (string, error)
}

type ServiceImpl struct {
}

var _ Service = &ServiceImpl{}

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

func (si *ServiceImpl) GenerateVersion4UUID() (string, error) {
	gen, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return gen.String(), nil
}
