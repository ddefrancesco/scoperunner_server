package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

type NotUsed struct {
	Name string
}

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id            string  `csv:"client_id"`
	Name          string  `csv:"client_name"`
	Age           string  `csv:"client_age"`
	NotUsedString string  `csv:"-"`
	NotUsedStruct NotUsed `csv:"-"`
}

type NGCRecord struct {
	//;Type;RA;Dec;Const;MajAx;MinAx;PosAng;B-Mag;V-Mag;J-Mag;H-Mag;K-Mag;SurfBr;Hubble;Pax;Pm-RA;Pm-Dec;RadVel;Redshift;Cstar U-Mag;Cstar B-Mag;Cstar V-Mag;M;NGC;IC;Cstar Names;Identifiers;Common names;NED notes;OpenNGC notes;Sources
	Name         string `csv:"Name"`
	Type         string `csv:"Type"`
	RA           string `csv:"RA"`
	Dec          string `csv:"Dec"`
	Const        string `csv:"Const"`
	MajAx        string `csv:"MajAx"`
	MinAx        string `csv:"MinAx"`
	PosAng       string `csv:"PosAng"`
	BMag         string `csv:"B-Mag"`
	VMag         string `csv:"V-Mag"`
	JMag         string `csv:"J-Mag"`
	HMag         string `csv:"H-Mag"`
	KMag         string `csv:"K-Mag"`
	SurfBr       string `csv:"SurfBr"`
	Hubble       string `csv:"Hubble"`
	Pax          string `csv:"Pax"`
	PmRA         string `csv:"Pm-RA"`
	PmDec        string `csv:"Pm-Dec"`
	RadVel       string `csv:"RadVel"`
	Redshift     string `csv:"Redshift"`
	CstarUMag    string `csv:"Cstar U-Mag"`
	CstarBMag    string `csv:"Cstar B-Mag"`
	CstarVMag    string `csv:"Cstar V-Mag"`
	M            string `csv:"M"`
	NGC          string `csv:"NGC"`
	IC           string `csv:"IC"`
	CstarNames   string `csv:"Cstar Names"`
	Identifiers  string `csv:"Identifiers"`
	CommonNames  string `csv:"Common names"`
	NEDNotes     string `csv:"NED notes"`
	OpenNGCNotes string `csv:"OpenNGC notes"`
	Sources      string `csv:"Sources"`
}

func main() {
	clientsFile, err := os.OpenFile("../NGC.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*NGCRecord{}
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		r.FieldsPerRecord = -1
		return r
	})

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}
	for _, client := range clients {
		fmt.Println("Hello", client.Name)
	}

	// if _, err := clientsFile.Seek(0, 0); err != nil { // Go to the start of the file
	// 	panic(err)
	// }

	// clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
	// clients = append(clients, &Client{Id: "13", Name: "Fred"})
	// clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	// clients = append(clients, &Client{Id: "15", Name: "Danny"})
	// csvContent, err := gocsv.MarshalString(&clients) // Get all clients as CSV string
	//err = gocsv.MarshalFile(&clients, clientsFile) // Use this to save the CSV back to the file
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(csvContent) // Display all clients as CSV string

}
