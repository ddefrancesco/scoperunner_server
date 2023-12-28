package interfaces

import "go.bug.st/serial"

type ScopeResponse struct {
	Err      *serial.PortError
	Response []byte
	ExecCmd  string
}

type SerialClient interface {
	Connect(serialPort string) (serial.Port, error)
	Disconnect(port serial.Port) error
	ExecCommand(scopecmd string) ScopeResponse
	FetchQuery(scopecmd string) ScopeResponse
}
