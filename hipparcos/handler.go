package hipparcos

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/ddefrancesco/scoperunner_server/cache"
	"github.com/ddefrancesco/scoperunner_server/errors"
	"github.com/gocarina/gocsv"
	"github.com/spf13/viper"
)

type HipparcosCatalog []HipparcosRecord

type HipparcosCacheHandler interface {
	GetHipparcosCatalog() (HipparcosCatalog, error)
	SetHipparcosCatalog(catalog HipparcosCatalog) error
	FindHipparcosObject(code string) (*HipparcosRecord, error)
}

type HipparcosRecord struct {
	Catalog string  `csv:"Catalog"`
	Hip     string  `csv:"HIP"`
	RAhms   string  `csv:"RAhms"`
	DEdms   string  `csv:"DEdms"`
	RAdeg   float64 `csv:"RAdeg"`
	DEdeg   float64 `csv:"DEdeg"`
}

func NewHipparcosCatalog() HipparcosCatalog {
	csvFilePath := viper.GetString("hipparcos.csv.path")
	catalog := ReadCsv(csvFilePath)
	// err := catalog.SetNGCCatalog(catalog)
	// if err != nil {
	// 	panic(err)
	// }
	return catalog
}

func (c HipparcosCatalog) GetHipparcosCatalog() (HipparcosCatalog, error) {
	// Get catalog from cache
	var hip_catalog HipparcosCatalog
	return hip_catalog, nil
}

func (c HipparcosCatalog) SetHipparcosCatalog(catalog HipparcosCatalog) error {
	// Store catalog in cache
	catCache := cache.New[string, HipparcosCatalog]()
	catCache.Set("hip_catalog", catalog)

	return nil
}

func (c HipparcosCatalog) FindHipparcosObject(code string) (*HipparcosRecord, error) {
	for _, obj := range c {
		if obj.Hip == code {
			return &obj, nil
		}
	}
	return nil, errors.NewObjectNotFoundInCatalogError(code)
}

func ReadCsv(csvFilePath string) HipparcosCatalog {
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

	var records []*HipparcosRecord
	// Parse the CSV data into the articles slice. If an error occurs, panic.
	if unmarshalError := gocsv.UnmarshalFile(csvFile, &records); unmarshalError != nil {
		panic(unmarshalError)
	}
	var catalog HipparcosCatalog
	for _, record := range records {
		catalog = append(catalog, *record)
	}
	return catalog

}

func hipparcosCatalogNames() map[string]int {
	m := make(map[string]int)

	m["Acamar"] = 13847
	m["Achernar"] = 7588
	m["Acrux	"] = 60718
	m["Adhara	"] = 33579
	m["Agena	"] = 68702
	m["Albireo"] = 95947
	m["Alcor	"] = 65477
	m["Alcyone"] = 17702
	m["Aldebaran"] = 21421
	m["Alderamin"] = 105199
	m["Algenib	"] = 1067
	m["Algieba	"] = 50583
	m["Algol	"] = 14576
	m["Alhena	"] = 31681
	m["Alioth	"] = 62956
	m["Alkaid	"] = 67301
	m["Almaak	"] = 9640
	m["Alnair	"] = 109268
	m["Alnath	"] = 25428
	m["Alnilam	"] = 26311
	m["Alnitak	"] = 26727
	m["Alphard	"] = 46390
	m["Alphekka	"] = 76267
	m["Alpheratz"] = 677
	m["Alshain	"] = 98036
	m["Altair	"] = 97649
	m["Ankaa	"] = 2081
	m["Antares	"] = 80763
	m["Arcturus	"] = 69673
	m["Arneb	"] = 25985
	m["Babcock star"] = 112247
	m["Barnard star"] = 87937
	m["Bellatrix	"] = 25336
	m["Betelgeuse	"] = 27989
	m["Campbell star"] = 96295
	m["Canopus		"] = 30438
	m["Capella		"] = 24608
	m["Caph				"] = 746
	m["Castor			"] = 36850
	m["Cor Caroli		"] = 63125
	m["Cyg X-1			"] = 98298
	m["Deneb			"] = 102098
	m["Denebola			"] = 57632
	m["Diphda			"] = 3419
	m["Dubhe			"] = 54061
	m["Enif				"] = 107315
	m["Etamin			"] = 87833
	m["Fomalhaut		"] = 113368
	m["Groombridge 1830	"] = 57939
	m["Hadar			"] = 68702
	m["Hamal			"] = 9884
	m["Izar				"] = 72105
	m["Kapteyn star		"] = 24186
	m["Kaus Australis	"] = 90185
	m["Kocab			"] = 72607
	m["Kruger 60		"] = 110893
	m["Luyten star		"] = 36208
	m["Markab			"] = 113963
	m["Megrez			"] = 59774
	m["Menkar			"] = 14135
	m["Merak			"] = 53910
	m["Mintaka			"] = 25930
	m["Mira				"] = 10826
	m["Mirach			"] = 5447
	m["Mirphak			"] = 15863
	m["Mizar			"] = 65378
	m["Nihal			"] = 25606
	m["Nunki			"] = 92855
	m["Phad				"] = 58001
	m["Pleione			"] = 17851
	m["Polaris			"] = 11767
	m["Pollux			"] = 37826
	m["Procyon			"] = 37279
	m["Proxima			"] = 70890
	m["Rasalgethi		"] = 84345
	m["Rasalhague		"] = 86032
	m["Red Rectangle	"] = 30089
	m["Regulus			"] = 49669
	m["Rigel			"] = 24436
	m["Rigil Kent		"] = 71683
	m["Sadalmelik		"] = 109074
	m["Sadalsuud		"] = 113963
	m["Sagitta			"] = 106315
	m["Sagittarius A*	"] = 113963
	m["Saiph			"] = 27366
	m["Scheat			"] = 113881
	m["Shaula			"] = 85927
	m["Shedir			"] = 3179
	m["Sheliak			"] = 92420
	m["Sirius			"] = 32349
	m["Spica			"] = 65474
	m["Suhail			"] = 113963
	m["Tarazed			"] = 97278
	m["Tegmine			"] = 106315
	m["Thuban			"] = 68756
	m["Unukalhai		"] = 77070
	m["Van Maanen 2		"] = 3829
	m["Vega				"] = 91262
	m["Vindemiatrix		"] = 63608
	m["Zaurak			"] = 18543
	m["3C 273			"] = 60936

	return m

}
