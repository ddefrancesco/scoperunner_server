package geocoding

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	dms "github.com/ddefrancesco/go-dms/dms"
	cache "github.com/ddefrancesco/scoperunner_server/geocoding/cache"
	"github.com/spf13/viper"
)

type AutostarLatLong struct {
	AutostarLat  string `json:"autostar-lat"`
	AutostarLong string `json:"autostar-long"`
}
type Address struct {
	Location string `json:"location"`
}

func GetAutostarLocation(address Address, geoCache *cache.Cache[string, *AutostarLatLong]) (*AutostarLatLong, error) {
	token := viper.GetString("geocoding.token")
	url := viper.GetString("geocoding.url")
	// Insert caching here
	//geoCache := cache.New[string, *AutostarLatLong]()
	cachedValue, found := geoCache.Get(address.Location)
	//setGeoLocationCache(address.Address, *dms, geoCache)
	if !found {
		payload := strings.NewReader("{\"address\":\"" + address.Location + "\"}")

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		return parseResponse(address, body, geoCache)
	}
	return cachedValue, nil
}

func parseResponse(address Address, body []byte, geoCache *cache.Cache[string, *AutostarLatLong]) (*AutostarLatLong, error) {
	var parsedData map[string]interface{}
	err := json.Unmarshal(body, &parsedData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	element := parsedData["element"]
	fmt.Println(element)

	lat := element.(map[string]interface{})["latitude"].(float64)
	fmt.Println(lat)

	long := element.(map[string]interface{})["longitude"].(float64)
	fmt.Println(long)

	dms := convert2DM(lat, long)
	geoCache.Set(address.Location, dms)
	return dms, nil
}

func convert2DM(lat float64, long float64) *AutostarLatLong {

	dmsCoordinate, err := dms.NewDMS(dms.DecimalDegrees{
		Latitude:  lat,
		Longitude: long,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DMS coordinates: %+v\n", dmsCoordinate.String())
	dmsLat := dmsCoordinate.AutostarLatitude(dmsCoordinate.Latitude)
	fmt.Printf("DMS Lat. %v\n", dmsLat)
	dmsLong := dmsCoordinate.AutostarLongitude(dmsCoordinate.Longitude)
	fmt.Printf("DMS Long. %v\n", dmsLong)

	return &AutostarLatLong{AutostarLat: dmsLat, AutostarLong: dmsLong}
}
