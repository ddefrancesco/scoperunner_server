package mocks

import (
	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"github.com/stretchr/testify/mock"
)

type EtxMock struct {
	mock.Mock
}

func (mock *EtxMock) ExecCommand(scopecmd string) interfaces.ETXResponse {
	args := mock.Called(scopecmd)
	return args.Get(0).(interfaces.ETXResponse)

}
func (mock *EtxMock) FetchQuery(scopecmd string) interfaces.ETXResponse {
	args := mock.Called(scopecmd)
	return args.Get(0).(interfaces.ETXResponse)
}
