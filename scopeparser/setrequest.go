package scopeparser

type SetRequest struct {
	Request map[string]string
}

type ScopeSetCmd struct {
	ValuesMap map[string]string
}

func NewSetRequest(m map[string]string) *SetRequest {
	setRequest := &SetRequest{
		Request: m,
	}
	return setRequest
}

func (s *SetRequest) ParseMap() (map[string]string, error) {
	settings := s.InitMap()
	var settingsMap ScopeSetCmd
	settingsMap.ValuesMap = make(map[string]string)
	for k, v := range s.Request {
		if _, ok := settings[k]; ok {

			settingsMap.ValuesMap[settings[k]] = v

		} else {

			return nil, ErrUnknownInfoCommand
		}
	}

	return settingsMap.ValuesMap, nil
}

func (s *SetRequest) InitMap() map[string]string {
	return settingsDictionary()
}

func settingsDictionary() map[string]string {

	m := make(map[string]string)
	m["altitude"] = "Sa"
	m["brighter_limit"] = "Sb"
	m["baud_speed"] = "SB"
	m["current_date"] = "SC"
	m["target_dec"] = "Sd"
	m["selenographic_lat"] = "SE"
	m["selenographic_long"] = "Se"
	m["fainter_limit"] = "Sf"
	m["field_diameter"] = "SF"
	m["current_site_long"] = "Sg"
	m["utc_offset"] = "SG"
	m["dst"] = "SH"
	m["min_elev_limit"] = "Sh"
	m["smallest_limit"] = "Sl"
	m["local_time"] = "SL"
	m["high_limit"] = "So"
	m["quality"] = "Sq"
	m["target_ra"] = "Sr"
	m["largest_limit"] = "Ss"
	m["local_sidereal_time"] = "SS"
	m["current_site_lat"] = "St"
	m["tracking_rate"] = "ST"
	m["max_slew_rate"] = "Sw"
	m["target_azimuth"] = "Sz"
	m["h12_h24_toggle"] = "H"

	return m
}
