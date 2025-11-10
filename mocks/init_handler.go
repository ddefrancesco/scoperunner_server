package mocks

type InitRequestMock struct {
	MockGetInitializeCommand func() (string, error)
}

func (m *InitRequestMock) GetInitializeCommand() (string, error) {
	return m.MockGetInitializeCommand()
}
