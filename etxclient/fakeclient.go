package etxclient

import (
	"log"
	"math/rand"
	"strings"
	"time"

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
		Response: []byte("Command Accepted"),
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
	case strings.Contains(scopecmd, ":GA#"):
		//sDD*MM’SS# Altitude
		sr.Response = []byte("s41*53’30#")

	case strings.Contains(scopecmd, ":Ga#"):
		//HH:MM:SS#
		layout := "03:04:05#"
		sr.Response = []byte(time.Now().Format(layout))

	case strings.Contains(scopecmd, ":Gb#"):
		//sMM.M# Magnitude
		sr.Response = []byte("s12.3#")

	case strings.Contains(scopecmd, ":GC#"):
		//MM/DD/YY#
		layout := "01/02/06#"
		sr.Response = []byte(time.Now().Format(layout))

	case strings.Contains(scopecmd, ":Gc#"):
		//12# or 24#

		sr.Response = []byte("#12")

	case strings.Contains(scopecmd, ":GD#"):
		//sDD*MM# or sDD*MM’SS# s degrees*minutes'seconds # (0-90)*(0-59)'(0-59)

		sr.Response = []byte("s41:12'34#")

	case strings.Contains(scopecmd, ":Gd#"):
		//sDD*MM# or sDD*MM’SS# s degrees*minutes'seconds # (0-90)*(0-59)'(0-59)

		sr.Response = []byte("s41:12'34#")

	case strings.Contains(scopecmd, ":GF#"):
		//NNN#

		sr.Response = []byte("123#")

	case strings.Contains(scopecmd, ":Gf#"):
		//sMM.M#
		sr.Response = []byte("s12.3#")

	case strings.Contains(scopecmd, ":GG#"):
		//sHH# or sHH.H UTC conversion differential

		sr.Response = []byte("s12#")

	case strings.Contains(scopecmd, ":Gg#"):
		//sDDD*MM# Latitude

		sr.Response = []byte("s+41*53#")

	case strings.Contains(scopecmd, ":Gh#"):
		//sDD* High Limit
		sr.Response = []byte("s90*")

	case strings.Contains(scopecmd, ":GL#"):
		//sHH:MM:SS# Local
		layout := "s03:04:05#"
		sr.Response = []byte(time.Now().Format(layout))

	case strings.Contains(scopecmd, ":Gl#"):
		//NNN’#
		sr.Response = []byte("023'#")

	case strings.Contains(scopecmd, ":Go#"):
		//DD*#
		sr.Response = []byte("12*#")

	case strings.Contains(scopecmd, ":Gq#"):
		/*SU# Super
				  EX# Excellent
				  VG# Very Good
		          GD# Good
				  FR# Fair
				  PR# Poor
				  VP# Very Poor*/
		qmap := map[int]string{1: "SU#", 2: "EX#", 3: "VG#", 4: "GD#", 5: "FR#", 6: "PR#", 7: "VP#"}
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		sr.Response = []byte(qmap[r1.Intn(7)])

	case strings.Contains(scopecmd, ":GR#"):
		//HH:MM.T# or HH:MM:SS# Telescope RA
		layout := "03:04:05#"
		sr.Response = []byte(time.Now().Format(layout))
	case strings.Contains(scopecmd, ":Gr#"):
		//HH:MM.T# or HH:MM:SS target RA
		layout := "03:04:05"
		sr.Response = []byte(time.Now().Format(layout))

	case strings.Contains(scopecmd, ":GS#"):
		//HH:MM.T# or HH:MM:SS target RA
		layout := "03:04:05"
		sr.Response = []byte(time.Now().Format(layout))

	case strings.Contains(scopecmd, ":Gs#"):
		//NNN'#
		sr.Response = []byte("123'#")

	case strings.Contains(scopecmd, ":GT#"):
		//TT.T# Tracking Rate
		sr.Response = []byte("15.0#")

	case strings.Contains(scopecmd, ":Gt#"):
		//sDD*MM# Site Latitude
		sr.Response = []byte("s41.58#")

	case strings.Contains(scopecmd, ":GVD#"):
		//mmm dd yyyy# Firmware Date
		sr.Response = []byte("mar 12 2008#")

	case strings.Contains(scopecmd, ":GVN#"):
		//dd.d# Firmware Number
		sr.Response = []byte("12.0#")

	case strings.Contains(scopecmd, ":GVP#"):
		//<stringa># Product Name
		sr.Response = []byte("Autostar#")

	case strings.Contains(scopecmd, ":GVT#"):
		//HH:MM:SS# Firmware Time
		layout := "01/02/06#"

		firmware_time := time.Date(2005, 6, 14, 00, 00, 00, 100, time.Local).Format(layout)
		sr.Response = []byte(firmware_time)

	case strings.Contains(scopecmd, ":Gy#"):
		//GPDCO# eepsky object search string
		sr.Response = []byte("GPDCO#")

	case strings.Contains(scopecmd, ":GZ#"):
		//DDD*MM#T or DDD*MM’SS# Telescope Azimuth
		sr.Response = []byte("000.00#0")

	default:
		sr.Response = []byte("Command Accepted")
	}

	log.Println("EtxClient::FetchQuery -> " + scopecmd + " eseguito")
	return sr

}

type FakeEtxClient struct {
}
