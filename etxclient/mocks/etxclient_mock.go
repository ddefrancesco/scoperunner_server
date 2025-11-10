package mocks

import (
	"time"

	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"go.bug.st/serial"
)

type EtxClientMock struct {
	// You can add fields here to customize the behavior of the mock
}

func NewEtxClientMock() *EtxClientMock {
	return &EtxClientMock{}
}

func (ec *EtxClientMock) Connect(serialPort string) (serial.Port, error) {
	// Simulate successful connection
	return &MockSerialPort{}, nil
}

func (ec *EtxClientMock) Disconnect(port serial.Port) error {
	// Simulate successful disconnection
	return nil
}

func (ec *EtxClientMock) ExecCommand(scopecmd string) interfaces.ETXResponse {
	// Simulate command execution based on the scopecmd
	response := interfaces.ETXResponse{
		Err:      nil,
		Response: []byte("Mock response for command: " + scopecmd),
		ExecCmd:  scopecmd,
	}
	return response
}

// MockSerialPort simulates a serial port for testing
type MockSerialPort struct{}

func (m *MockSerialPort) Write(data []byte) (int, error) {
	// Simulate writing data to the port
	return len(data), nil
}

func (m *MockSerialPort) Read(buff []byte) (int, error) {
	// Simulate reading data from the port
	copy(buff, "Mock data response")
	return len("Mock data response"), nil
}

func (m *MockSerialPort) ResetInputBuffer() error {
	// Simulate resetting the input buffer
	return nil
}

func (m *MockSerialPort) ResetOutputBuffer() error {
	// Simulate resetting the output buffer
	return nil
}

func (m *MockSerialPort) Close() error {
	// Simulate closing the port
	return nil
}

func (m *MockSerialPort) Break(duration time.Duration) error {
	// Simulate break signal
	return nil
}

// Implement any additional methods required by the serial.Port interface
func (m *MockSerialPort) SetReadTimeout(timeout time.Duration) error {
	// Simulate setting read timeout
	return nil
}

func (m *MockSerialPort) GetModemStatusBits() (*serial.ModemStatusBits, error) {
	// Simulate getting modem status bits
	return &serial.ModemStatusBits{}, nil
}

func (m *MockSerialPort) Drain() error {
	// Simulate draining the port
	return nil
}

func (m *MockSerialPort) SetDTR(dtr bool) error {
	// Simulate setting DTR
	return nil
}

func (m *MockSerialPort) SetRTS(rts bool) error {
	// Simulate setting RTS
	return nil
}

func (m *MockSerialPort) SetMode(mode *serial.Mode) error {
	// Simulate setting mode
	return nil
}

// Ensure that MockSerialPort implements the serial.Port interface
var _ serial.Port = (*MockSerialPort)(nil)
