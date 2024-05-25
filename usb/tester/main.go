package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"go.bug.st/serial"
)

func main() {

	// c := &serial.Config{
	// 	Name:        "/dev/cu.usbserial-0001",
	// 	Baud:        9600,
	// 	Parity:      serial.ParityNone,
	// 	StopBits:    serial.Stop1,
	// 	ReadTimeout: 3 * time.Second,
	// }
	// s, err := serial.OpenPort(c)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer s.Close()
	// err = s.Flush()
	// if err != nil {
	// 	log.Fatal(s.Flush().Error())
	// }
	// m, err := s.Write([]byte(":Ga#\n\r"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Sent %v bytes\n", m)
	// buf := make([]byte, 32)
	// n, err := s.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v", string(buf[:n]))

	// err = s.Flush()
	// if err != nil {
	// 	log.Fatal(s.Flush().Error())
	// }
	// err = s.Close()
	// if err != nil {
	// 	log.Fatal(s.Close().Error())
	// }
	// m, err = s.Write([]byte("ACK\x06\n\r"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Sent %v bytes\n", m)
	// buf = make([]byte, 128)
	// n, err = s.Read(buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v", string(buf[:n]))

	// err = s.Flush()
	// if err != nil {
	// 	log.Fatal(s.Flush().Error())
	// }
	// ======================= serial import port code ========================
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// // Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	// Open the first serial port detected at 9600bps N81
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	port, err := serial.Open(ports[1], mode)
	if err != nil {
		log.Fatal(err)
	}
	msb, err := port.GetModemStatusBits()
	if err != nil {
		log.Fatal(err)
	}
	// Check if the CTS line is set
	if !msb.CTS {

		fmt.Println("CTS line is set")
	}
	// Check if the DSR line is set
	if msb.DSR {
		fmt.Println("DSR line is set")
	}
	// Check if the RI line is set
	if msb.RI {
		fmt.Println("RI line is set")
	}
	// Check if the DCD line is set
	if msb.DCD {
		fmt.Println("DCD line is set")
	}
	// Send the string "10,20,30\n\r" to the serial port
	//cmd := initializatitionSequence(prepareSetDateCommand(), prepareSetTimeCommand())
	cmd := ":Q#"
	n, err := port.Write([]byte(cmd))
	if err != nil {
		log.Fatal(err)
	}
	//port.Drain()
	defer port.Close()
	fmt.Printf("Sent %v bytes\n", n)

	// Read and print the response

	buff := make([]byte, 128)
	port.ResetInputBuffer()
	port.SetReadTimeout(time.Millisecond * 1000)

	for {

		//Reads up to 20 bytes
		n, err = port.Read(buff)
		if err != nil {
			log.Fatal(err)
		}
		if n == 0 {
			port.ResetInputBuffer()
			port.ResetOutputBuffer()
			fmt.Println("\nEOF")
			break
		}

		fmt.Printf("%s", string(buff[:n]))

		//If we receive a newline stop reading
		if strings.Contains(string(buff[:n]), "\n") {
			port.ResetInputBuffer()
			port.ResetOutputBuffer()

			break
		}

	}

}

func prepareSetAltitudeCommand(altitude string) string {
	return ":SA" + altitude + "#"
}
func prepareSetDateCommand() string {
	layout := "01/02/06#"
	date := time.Now().Format(layout)
	return ":SC" + date + "#"
}

func prepareSetTimeCommand() string {
	layout := "03:04:05#"
	t := time.Now().Format(layout)
	return ":SL" + t + "#"
}

func prepareSetMagnitudeCommand(magnitude string) string {
	return ":SM" + magnitude + "#"
}

func prepareSetObjectSelectionCommand(object string) string {
	return ":Sy" + object + "#"
}

func prepareSetObjectAzimuthCommand(azimuth string) string {
	return ":Sz" + azimuth + "#"
}

func prepareSetSlewRateCommand(rate string) string {
	return ":SR" + rate + "#"
}

func prepareSetCurrentSiteLongCommand(longitude string) string {
	return ":Sg" + longitude + "#"
}

func prepareSetCurrentSiteLatCommand(latitude string) string {
	return ":Sts" + latitude + "#"
}

func initializatitionSequence(cmds ...string) string {
	return strings.Join(cmds, "")
}
