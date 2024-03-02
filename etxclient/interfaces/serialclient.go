package interfaces

import "go.bug.st/serial"

type ETXResponse struct {
	Err      *serial.PortError
	Response []byte
	ExecCmd  string
}

type SerialClient interface {
	Connect(serialPort string) (serial.Port, error)
	Disconnect(port serial.Port) error
	ExecCommand(scopecmd string) ETXResponse
	FetchQuery(scopecmd string) ETXResponse
}
