package cache

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/ddefrancesco/scoperunner_server/errors"
	"github.com/ddefrancesco/scoperunner_server/geocoding/cache"
	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

type NGCRecord struct {
	//;Type;RA;Dec;Const;MajAx;MinAx;PosAng;B-Mag;V-Mag;J-Mag;H-Mag;K-Mag;SurfBr;Hubble;Pax;Pm-RA;Pm-Dec;RadVel;Redshift;Cstar U-Mag;Cstar B-Mag;Cstar V-Mag;M;NGC;IC;Cstar Names;Identifiers;Common names;NED notes;OpenNGC notes;Sources
	Name         string  `csv:"Name"`
	Type         string  `csv:"Type"`
	RA           string  `csv:"RA"`
	Dec          string  `csv:"Dec"`
	Const        string  `csv:"Const"`
	MajAx        float32 `csv:"MajAx"`
	MinAx        float32 `csv:"MinAx"`
	PosAng       int     `csv:"PosAng"`
	BMag         float32 `csv:"B-Mag"`
	VMag         float32 `csv:"V-Mag"`
	JMag         float32 `csv:"J-Mag"`
	HMag         float32 `csv:"H-Mag"`
	KMag         float32 `csv:"K-Mag"`
	SurfBr       float32 `csv:"SurfBr"`
	Hubble       string  `csv:"Hubble"`
	Pax          string  `csv:"Pax"`
	PmRA         string  `csv:"Pm-RA"`
	PmDec        string  `csv:"Pm-Dec"`
	RadVel       int     `csv:"RadVel"`
	Redshift     float64 `csv:"Redshift"`
	CstarUMag    float32 `csv:"Cstar U-Mag"`
	CstarBMag    float32 `csv:"Cstar B-Mag"`
	CstarVMag    float32 `csv:"Cstar V-Mag"`
	M            string  `csv:"M"`
	NGC          string  `csv:"NGC"`
	IC           string  `csv:"IC"`
	CstarNames   string  `csv:"Cstar Names"`
	Identifiers  string  `csv:"Identifiers"`
	CommonNames  string  `csv:"Common names"`
	NEDNotes     string  `csv:"NED notes"`
	OpenNGCNotes string  `csv:"OpenNGC notes"`
	Sources      string  `csv:"Sources"`
}

type NGCCatalog []NGCRecord

type NGCCacheHandler interface {
	GetNGCCatalog() (NGCCatalog, error)
	SetNGCCatalog(catalog NGCCatalog) error
	FindNGCObject(code string) (*NGCRecord, error)
}

func NewNGCCatalog() NGCCatalog {
	csvFilePath := viper.GetString("openngc.csv.path")
	catalog := ReadCsv(csvFilePath)
	err := catalog.SetNGCCatalog(catalog)
	if err != nil {
		panic(err)
	}
	return catalog
}

func (h *NGCCatalog) SetNGCCatalog(catalog NGCCatalog) error {

	// Store catalog in cache
	//

	catCache := cache.New[string, NGCCatalog]()
	catCache.Set("catalog", catalog)

	return nil
}
func (h *NGCCatalog) GetNGCCatalog() (NGCCatalog, error) {
	// Get catalog from cache
	var catalog NGCCatalog

	return catalog, nil
}

func (c NGCCatalog) FindNGCObject(name string) (*NGCRecord, error) {
	for _, obj := range c {
		if obj.Name == name {
			return &obj, nil
		}
	}
	return nil, errors.NewObjectNotFoundInCatalogError(name)
}

func ReadCsv(csvFilePath string) NGCCatalog {
	// Try to open the example.csv file in read-write mode.
	csvFile, csvFileError := os.OpenFile(csvFilePath, os.O_RDWR, os.ModePerm)
	// If an error occurs during os.OpenFIle, panic and halt execution.
	if csvFileError != nil {
		panic(csvFileError)
	}
	// Ensure the file is closed once the function returns
	defer csvFile.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		r.FieldsPerRecord = -1
		return r
	})

	var records []*NGCRecord
	// Parse the CSV data into the articles slice. If an error occurs, panic.
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &records); unmarshalError != nil {
		panic(unmarshalError)
	}
	var catalog NGCCatalog
	for _, record := range records {
		catalog = append(catalog, *record)
	}
	return catalog
}
