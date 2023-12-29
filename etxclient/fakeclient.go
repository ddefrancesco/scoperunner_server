package etxclient

import (
	"log"
	"strings"

	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"go.bug.st/serial"
)

func NewFakeClient() *FakeEtxClient {
	etxclient := &FakeEtxClient{}
	return etxclient
}

func (ec *FakeEtxClient) Connect(serialPort string) (serial.Port, error) {
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

func (ec *FakeEtxClient) Disconnect(port serial.Port) error {
	err := port.Close()
	return err
}

func (ec *FakeEtxClient) ExecCommand(scopecmd string) interfaces.ScopeResponse {

	// TODO: Open serial
	//       Exec Command scope
	// 		 Close serial

	log.Println("EtxClient::ExecCommand -> " + scopecmd + " eseguito")
	return interfaces.ScopeResponse{
		Err:      nil,
		Response: nil,
		ExecCmd:  scopecmd,
	}
}

func (ec *FakeEtxClient) FetchQuery(scopecmd string) interfaces.ScopeResponse {
	// TODO: Open serial
	//       Exec Command scope
	// 		 Close serial
	sr := interfaces.ScopeResponse{
		Err:      nil,
		Response: nil,
		ExecCmd:  scopecmd,
	}
	switch {
	case strings.Contains(scopecmd, "ACK"):
		sr.Response = []byte("A") //A,L,P,D
	default:
		sr.Response = []byte("D")
	}

	log.Println("EtxClient::FetchQuery -> " + scopecmd + " eseguito")
	return sr

}

type FakeEtxClient struct {
}
