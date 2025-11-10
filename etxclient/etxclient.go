package etxclient

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"

	"github.com/ddefrancesco/scoperunner_server/etxclient/interfaces"
	"github.com/spf13/viper"
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
	msb, err := port.GetModemStatusBits()
	if err != nil {
		log.Fatal(err)
	}
	// Check if the CTS line is set
	if !msb.CTS {

		log.Println("CTS line is set")
	}
	// Check if the DSR line is set
	if msb.DSR {
		log.Println("DSR line is set")
	}
	// Check if the RI line is set
	if msb.RI {
		log.Println("RI line is set")
	}
	// Check if the DCD line is set
	if msb.DCD {
		log.Println("DCD line is set")
	}
	port.ResetInputBuffer()
	port.ResetOutputBuffer()
	return port, err
}

func (ec *EtxClient) Disconnect(port serial.Port) error {
	err := port.ResetInputBuffer()
	if err != nil {
		return err
	}
	err = port.ResetOutputBuffer()
	if err != nil {
		return err
	}
	err = port.Close()
	return err
}

func (ec *EtxClient) ExecCommand(scopecmd string) interfaces.ETXResponse {

	// TODO: Open serial
	//       Exec Command scope
	// 		 Close serial

	shouldReturn, returnValue := ec.ConnectTTY(scopecmd)
	if shouldReturn {
		return returnValue
	}

	log.Printf("%v", string(returnValue.Response))
	log.Println("EtxClient::ExecCommand -> " + scopecmd + " eseguito")
	return returnValue
}

func (ec *EtxClient) ConnectTTY(scopecmd string) (bool, interfaces.ETXResponse) {
	log.Println("EtxClient::ConnectTTY -> " + scopecmd)
	serialport := viper.GetString("serialport.name")
	timeout := viper.GetDuration("serialport.timeout") * time.Millisecond
	var response interfaces.ETXResponse
	var accumulator []byte
	port, err := ec.Connect(serialport)
	if err != nil {
		log.Fatal(err)
		port.ResetInputBuffer()
		port.ResetOutputBuffer()

	}
	defer ec.Disconnect(port)
	n, err := port.Write([]byte(scopecmd))
	if err != nil {
		log.Fatal(err)
		port.ResetInputBuffer()
		port.ResetOutputBuffer()

	}

	log.Printf("EtxClient::ConnectTTY -> Sent %v bytes\n", n)
	buff := make([]byte, 128)
	port.ResetInputBuffer()
	port.SetReadTimeout(timeout)

	for {
		n, err := port.Read(buff)
		if err != nil {
			log.Fatal(err)

			return true, interfaces.ETXResponse{
				Err:      &serial.PortError{},
				Response: nil,
				ExecCmd:  scopecmd,
			}
		}

		accumulator = append(accumulator, buff[:n]...)

		fmt.Printf("%s", string(accumulator))
		if n == 0 {
			port.ResetInputBuffer()
			port.ResetOutputBuffer()
			fmt.Println("\nEOF")
			break
		}

	}
	response = interfaces.ETXResponse{
		Err:      nil,
		Response: accumulator,
		ExecCmd:  scopecmd,
	}
	return false, response
}

type EtxClient struct {
}
