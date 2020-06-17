package uuid_test

type MockUUIDService struct {
	MockErr  error
	MockResp string
}

func (m *MockUUIDService) GenerateVersion4UUID() (string, error) {
	return m.MockResp, m.MockErr
}
