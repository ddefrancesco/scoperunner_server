package etxclient

import (
	"log"

	"go.bug.st/serial"
)

type ScopeResponse struct {
	Err      *serial.PortError
	Response []byte
	ExecCmd  string
}

func NewClient() *EtxClient {
	etxclient := &EtxClient{}
	return etxclient
}

func (ec *EtxClient) Connect(serialPort string) (serial.Port, error) {
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	log.Println("ConnectCommand::Connect -> port: " + serialPort)
	port, err := serial.Open(serialPort, mode)
	if err != nil {
		log.Fatal(err)
	}
	return port, err
}

func (ec *EtxClient) Disconnect(port serial.Port) error {
	err := port.Close()
	return err
}

func (ec *EtxClient) ExecReturnNothing(scopecmd string) ScopeResponse {
	// TODO: Open serial
	//       Exec Command scope
	// 		 Close serial
	log.Println("ExecCommand::ExecReturnNothing -> " + scopecmd + " eseguito")
	return ScopeResponse{
		Err:      nil,
		Response: nil,
		ExecCmd:  scopecmd,
	}
}

func (ec *EtxClient) ExecReturnData(scopecmd string) {
	// TODO: Open serial
	//       Exec Command scope
	// 		 Close serial
	log.Println("ExecCommand::ExecReturnData -> " + scopecmd + " eseguito")
}

type EtxClient struct {
}
