package scopeparser

import (
	"time"

	"github.com/ddefrancesco/scoperunner_server/geocoding"
	"github.com/ddefrancesco/scoperunner_server/geocoding/cache"
	"github.com/ddefrancesco/scoperunner_server/models/commons"
)

var geoCache = cache.New[string, *geocoding.AutostarLatLong]()

type InitializeRequest struct {
	Request commons.RequestAddress
}

type ScopeInitCmd struct {
	ValuesMap map[string]string
}

func NewInitRequest(m commons.RequestAddress) *InitializeRequest {
	initRequest := &InitializeRequest{
		Request: m,
	}
	return initRequest
}

// func (s *InitializeRequest) ParseMap() (map[string]string, error) {
// 	settings := s.InitMap()
// 	var settingsMap ScopeInitCmd
// 	settingsMap.ValuesMap = make(map[string]string)
// 	for k, v := range s.Request {
// 		if _, ok := settings[k]; ok {

// 			settingsMap.ValuesMap[settings[k]] = v.Address

// 		} else {

// 			return nil, ErrUnknownInfoCommand
// 		}
// 	}

// 	return settingsMap.ValuesMap, nil
// }

// func (s *InitializeRequest) InitMap() map[string]string {

// 	return s.initializationDictionary()
// }
// func (s *InitializeRequest) getCacheInstance() (cache.Cache[string,*geocoding.AutostarLatLong], error) {
// 	if cacheInstance, found := geoCache; found {
// 		return cacheInstance, nil
// 	} else {
// 		return nil, ErrCacheNotFound

//		}
//	}
func (s *InitializeRequest) initializationDictionary() map[string]string {

	m := make(map[string]string)
	m["toggle_precision"] = ":U#"
	m["current_date"] = s.SetDateCommand()
	m["utc_offset"] = "SG"
	m["dst"] = "SH"
	m["local_time"] = s.SetTimeCommand()
	m["local_sidereal_time"] = "SS"
	m["current_site_lat"] = "St"
	m["current_site_long"] = "Sg"
	return m
}

func (s *InitializeRequest) SetDateCommand() string {
	layout := "01/02/06#"
	date := time.Now().Format(layout)
	cmd := ":SC" + date
	return cmd
}

func (s *InitializeRequest) SetTimeCommand() string {
	layout := "15:04:05#"
	timeZone := "Europe/Rome"
	loc, _ := time.LoadLocation(timeZone)
	initTime := time.Now().In(loc).Format(layout)
	cmd := ":SL" + initTime
	return cmd
}

func (s *InitializeRequest) SetLatitudeCommand(address geocoding.Address) (string, error) {

	autostarLoc, err := geocoding.GetAutostarLocation(address, geoCache)
	if err != nil {
		return "", err
	}
	cmd := ":St" + autostarLoc.AutostarLat + "#"
	return cmd, nil
}

func (s *InitializeRequest) SetLongitudeCommand(address geocoding.Address) (string, error) {

	autostarLoc, err := geocoding.GetAutostarLocation(address, geoCache)
	if err != nil {
		return "", err
	}
	cmd := ":Sg" + autostarLoc.AutostarLong + "#"
	return cmd, nil
}

func (s *InitializeRequest) SetUTCCommand() string {
	layout := "s03:04:05#"
	initTime := time.Now().Format(layout)
	cmd := "SG" + initTime
	return cmd
}

func (s *InitializeRequest) TogglePrecisionCommand() string {
	return s.initializationDictionary()["toggle_precision"]
}

func (s *InitializeRequest) SetInitializeCommand() (string, error) {
	addrRequest := s.Request
	address := geocoding.Address{Location: addrRequest.Address}
	lat, _ := s.SetLatitudeCommand(address)
	long, _ := s.SetLongitudeCommand(address)

	return s.TogglePrecisionCommand() + s.SetDateCommand() + s.SetTimeCommand() + lat + long, nil
}

func (s *InitializeRequest) SetCurrentDateCommand() string {
	return s.initializationDictionary()["current_date"]

}
