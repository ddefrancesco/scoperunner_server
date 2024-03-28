package etxclient

import (
	"fmt"
	"log"

	"go.bug.st/serial"

	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
)

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

func (ec *EtxClient) ExecCommand(scopecmd string) interfaces.ETXResponse {

	// TODO: Open serial
	//       Exec Command scope
	// 		 Close serial

	n, buff, shouldReturn, returnValue := ec.ConnectTTY(scopecmd)
	if shouldReturn {
		return returnValue
	}
	log.Printf("%v", string(buff[:n]))
	log.Println("EtxClient::ExecCommand -> " + scopecmd + " eseguito")
	return returnValue
}

func (ec *EtxClient) ConnectTTY(scopecmd string) (int, []byte, bool, interfaces.ETXResponse) {
	port, err := ec.Connect("/dev/ttyUSB0")
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Disconnect(port)
	n, err := port.Write([]byte(scopecmd))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sent %v bytes\n", n)
	buff := make([]byte, 1024)
	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)
			return 0, nil, true, interfaces.ETXResponse{
				Err:      &serial.PortError{},
				Response: nil,
				ExecCmd:  scopecmd,
			}
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}
	}
	return n, buff, false, interfaces.ETXResponse{}
}

type EtxClient struct {
}
