package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	dms "github.com/ddefrancesco/go-dms/dms"
)

func main() {

	//token := viper.GetString("geocoding.token")

	url := "https://geocoding.openapi.it/geocode"

	payload := strings.NewReader("{\"address\":\"Via Calcutta, Roma RM\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	//req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Authorization", "Bearer 662401dfa874ef890f0df052")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	//fmt.Println(res)
	//fmt.Println(string(body))

	var parsedData map[string]interface{}
	err := json.Unmarshal(body, &parsedData)
	if err != nil {
		fmt.Println(err)
		return
	}
	element := parsedData["element"]
	fmt.Println(element)

	lat := element.(map[string]interface{})["latitude"].(float64)
	fmt.Println(lat)

	long := element.(map[string]interface{})["longitude"].(float64)
	fmt.Println(long)

	dms := convert2DM(lat, long)
	fmt.Println(dms)
	// log.Println("Geocoding call begin")
	// cfg := &swagger.Configuration{
	// 	BasePath: "https://geocoding.openapi.it",
	// }
	// log.Println("Geocoding conf set")
	// client := *swagger.NewAPIClient(cfg)
	// log.Println("Geocoding client set")
	// addr := swagger.Address{
	// 	Address: "Via Calcutta, Roma RM",
	// }
	// geocodeOpts := &swagger.GeocodeApiGeocodeOpts{
	// 	Body: optional.NewInterface(addr),
	// }
	// log.Printf("Geocoding Opts set %v", geocodeOpts.Body.Value())
	// place, resp, err := client.GeocodeApi.Geocode(context.Background(), geocodeOpts)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// log.Println("Geocoding call success " + resp.Status)

	// log.Printf("Geocoding place %v", place)
	// //responseBody, _ := json.Marshal(response)

}

func convert2DM(lat float64, long float64) string {
	//start := time.Now()
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

	//end := time.Now()
	//fmt.Printf("Function took %f seconds.\n", end.Sub(start).Seconds())
	return dmsCoordinate.String()
}
