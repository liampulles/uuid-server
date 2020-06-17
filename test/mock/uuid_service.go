package mock

// UUIDService mocks uuid.Service
type UUIDService struct {
	MockErr  error
	MockResp string
}

// GenerateVersion4UUID implements the interface
func (m *UUIDService) GenerateVersion4UUID() (string, error) {
	return m.MockResp, m.MockErr
}
