package converters

import (
	"strings"

	"github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func RequestParamsToInfoArray(params string) ([]scopeparser.Info, error) {
	pArray := strings.Split(params, ",")
	infoMap := InitInfoMap()
	var infoArray []scopeparser.Info
	for _, p := range pArray {
		if _, ok := infoMap[p]; ok {
			infoArray = append(infoArray, infoMap[p])
		} else {
			return nil, scopeparser.ErrUnknownInfoCommand
		}
	}

	return infoArray, nil
}

// func InfoArrayToInfoCommandMap(infoArray []scopeparser.Info) map[scopeparser.Info]scopeparser.InfoCommandValue {
// 	infoMap :=
// 	var infoCommandMap = make(map[scopeparser.Info]scopeparser.InfoCommandValue)
// 	for _, info := range infoArray {
// 		if _, ok := infoMap[info]; ok {
// 			infoCommandMap[info] = infoMap[info]
// 		}
// 	}
// 	return infoCommandMap

// }

func InitInfoMap() map[string]scopeparser.Info {
	infoMap := make(map[string]scopeparser.Info)
	infoMap["altitude"] = scopeparser.InfoAltitude
	infoMap["ltt"] = scopeparser.InfoLTT
	infoMap["browse_bml"] = scopeparser.InfoBrighterMagLimit
	infoMap["current_date"] = scopeparser.InfoCurrentDate
	infoMap["clock_fmt"] = scopeparser.InfoClockFmt
	infoMap["declination"] = scopeparser.InfoDeclination
	infoMap["sel_target_dec"] = scopeparser.InfoSelectedTargetDec
	infoMap["field_diameter"] = scopeparser.InfoFieldDiameter
	infoMap["fainter_mag_limit"] = scopeparser.InfoFainterMagLimit
	infoMap["utc_offset"] = scopeparser.InfoUTCOffset
	infoMap["current_site_long"] = scopeparser.InfoCurrentSiteLong
	infoMap["high_limit"] = scopeparser.InfoHighLimit
	infoMap["local_time_24h"] = scopeparser.InfoLocalTime24h
	infoMap["larger_size_limit"] = scopeparser.InfoLargerSizeLimit
	infoMap["lower_size_limit"] = scopeparser.InfoLowerSizeLimit
	infoMap["minimum_find_quality"] = scopeparser.InfoMinimumQuality
	infoMap["ra"] = scopeparser.InfoRA
	infoMap["sel_target_ra"] = scopeparser.InfoCurrentTargetRA
	infoMap["sidereal_time"] = scopeparser.InfoSiderealTime
	infoMap["smaller_size_limit"] = scopeparser.InfoSmallerSizeLimit
	infoMap["tracking_rate"] = scopeparser.InfoTrackingRate
	infoMap["current_site_lat"] = scopeparser.InfoCurrentSiteLat
	infoMap["firmware_date"] = scopeparser.InfoFirmwareDate
	infoMap["firmware_version"] = scopeparser.InfoFirmwareVersion
	infoMap["product_name"] = scopeparser.InfoProductName
	infoMap["firmware_time"] = scopeparser.InfoFirmwareTime
	infoMap["deepsky"] = scopeparser.InfoDeepsky
	infoMap["azimuth"] = scopeparser.InfoAzimuth
	return infoMap
}
