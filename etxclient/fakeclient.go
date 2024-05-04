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

func (ec *FakeEtxClient) ExecCommand(scopecmd string) interfaces.ETXResponse {
	sr := interfaces.ETXResponse{
		Err:      nil,
		Response: []byte("Command Accepted"),
		ExecCmd:  scopecmd,
	}
	// TODO: Open serial
	//       Exec Command scope
	// 		 Close serial
	switch {
	case strings.Contains(scopecmd, ":Sa"):
		//Set target object altitude to sDD*MM# or sDD*MM’SS# [LX 16”, Autostar, LX200GPS/RCX400]
		// Returns:
		// 1 Object within slew range
		// 0 Object out of slew range
		sr.Response = append(sr.Response, []byte(" 1")...)
	case strings.Contains(scopecmd, ":Sb"):
		/*Set Brighter limit to the ASCII decimal magnitude string. SMM.M
		Returns:
		0 - Valid
		1 – invalid number*/
		sr.Response = []byte("0")

	case strings.Contains(scopecmd, ":SB"):
		/*Set Baud Rate n, where n is an ASCII digit (1..9) with the following interpertation
					1 56.7K
					2 38.4K
					3 28.8K
					4 19.2K
					5 14.4K
					6 9600
					7 4800
					8 2400
					9 1200
		Returns:
		1 At the current baud rate and then changes to the new rate for further communication*/

		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":SC"):
		// 		Change Handbox Date to MM/DD/YY
		// 		Returns: <D><string>
		// 			D = ‘0’ if the date is invalid. The string is the null string.
		// 			D = ‘1’ for valid dates and the string is “Updating Planetary Data# #”
		// Note: For LX200GPS/RCX400 this is the UTC data!
		a := append(sr.Response, []byte(" 1 ")...)
		b := []byte("Updating Planetary Data# ")
		c := append(a, b...)

		sr.Response = []byte(c)

	case strings.Contains(scopecmd, ":Sd"):
		//Set target object declination to sDD*MM or sDD*MM:SS depending on the current precision setting
		// Returns:
		// 1 - Dec Accepted
		// 0 – Dec invalid

		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":SE"):
		//Sets target object to the specificed selenographic latitude on the Moon.
		// Returns 1 - If moon is up and coordinates are accepted.
		// 		   0 – If the coordinates are invalid

		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":Se"):
		//Sets the target object to the specified selenogrphic longitude on the Moon
		// Returns 1 – If the Moon is up and coordinates are accepted.
		// 		   0 – If the coordinates are invalid for any reason.

		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":Sf"):
		//Set faint magnitude limit to sMM.M
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid
		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":SF"):
		//Set FIELD/IDENTIFY field diameter to NNN arc minutes.
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid

		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":Sg"):
		//Set current site’s longitude to DDD*MM an ASCII position string
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid

		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":SG"):
		//Set the number of hours added to local time to yield UTC
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid
		// Generate a random number between 0-1 to determine valid/invalid response
		//rand.Seed(time.Now().UnixNano())
		validReturn(sr)

	case strings.Contains(scopecmd, ":SH"):
		//Set Dightlight Savings Mode [Autostar II Only]. D=1 Sets Daylight savings. D=0 Clears Daylight savings.
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid

		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":Sh"):
		//Set the minimum object elevation limit to DD#
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid
		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":Sl#"):
		//Set the size of the smallest object returned by FIND/BROWSE to NNNN arc minutes
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid
		sr.Response = append(sr.Response, []byte("1")...)

	case strings.Contains(scopecmd, ":SL"):
		/*Set the local Time
		Returns:
			0 – Invalid
			1 - Valid*/
		sr.Response = append(sr.Response, []byte("1")...)

	case strings.Contains(scopecmd, ":So"):
		//Set highest elevation to which the telescope will slew
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid

		sr.Response = append(sr.Response, []byte("1")...)
	case strings.Contains(scopecmd, ":Sq"):
		//Step the quality of limit used in FIND/BROWSE through its cycle of VP … SU. Current setting can be queried with :Gq#
		// Returns: Nothing

	case strings.Contains(scopecmd, ":Sr"):
		//Set target object RA to HH:MM.T or HH:MM:SS depending on the current precision setting.
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid

		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":Ss"):
		//Set the size of the largest object the FIND/BROWSE command will return to NNNN arc minutes
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid
		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":SS"):
		//Sets the local sidereal time to HH:MM:SS
		//Returns:
		//			0 – Invalid
		//			1 - Valid
		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":St#"):
		//Sets the current site latitude to sDD*MM#
		// Returns:
		// 		0 – Invalid
		//      1 - Valid
		sr.Response = append(sr.Response, []byte("1")...)

	case strings.Contains(scopecmd, ":ST"):
		//Sets the current tracking rate to TTT.T hertz, assuming a model where a 60.0 Hertz synchronous motor will cause the RA
		//axis to make exactly one revolution in 24 hours.
		//Returns:
		//			0 – Invalid
		//			1 - Valid
		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":Sw"):
		//Set maximum slew rate to N degrees per second. N is the range (2..8)
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid
		sr.Response = append(sr.Response, []byte(" 1")...)

	case strings.Contains(scopecmd, ":Sy"):
		//Sets the object selection string used by the FIND/BROWSE command.
		// Returns:
		//		0 – Invalid
		// 		1 - Valid
		sr.Response = append(sr.Response, []byte("1")...)

	case strings.Contains(scopecmd, ":Sz"):
		//Sets the target Object Azimuth [LX 16” and LX200GPS/RCX400 only]
		// Returns:
		// 		0 – Invalid
		// 		1 - Valid

		sr.Response = append(sr.Response, []byte("1")...)

	case strings.Contains(scopecmd, "ACK\x06"):
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

	log.Println("EtxClient::ExecCommand -> " + scopecmd + " eseguito")
	return interfaces.ETXResponse{
		Err:      nil,
		Response: sr.Response,
		ExecCmd:  scopecmd,
	}
}

func validReturn(sr interfaces.ETXResponse) {
	r := rand.Float64()
	if r < 0.5 {
		sr.Response = append(sr.Response, []byte(" 1")...)
	} else {
		sr.Response = []byte("0")
	}
}

type FakeEtxClient struct {
}

func parse(s string) (time.Time, error) {
	d, err := time.Parse("01/02/06", s)
	if err != nil {
		return d, err
	}
	return d, nil
}
