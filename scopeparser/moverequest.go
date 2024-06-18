package scopeparser

import "errors"

var ErrUnknownMoveCommand = errors.New("Unknown Move/Halt Command")

type MoveRequest struct {
	Request map[string]string
}

func NewMoveRequest(m map[string]string) *MoveRequest {
	request := &MoveRequest{
		Request: m,
	}
	return request
}

type ScopeMoveCmd struct {
	ValuesMap map[string]string
}

func (s *MoveRequest) ParseMap() (map[string]string, error) {
	settings := s.InitMap()
	var settingsMap ScopeMoveCmd
	settingsMap.ValuesMap = make(map[string]string)
	for k, v := range s.Request {
		if _, ok := settings[k]; ok {

			settingsMap.ValuesMap[settings[k]] = v

		} else {

			return nil, ErrUnknownMoveCommand
		}
	}

	return settingsMap.ValuesMap, nil
}

func (s *MoveRequest) InitMap() map[string]string {
	return moveCmdDictionary()
}

func moveCmdDictionary() map[string]string {

	m := make(map[string]string)
	m["quit_north"] = "Qn"
	m["quit_south"] = "Qs"
	m["quit_east"] = "Qe"
	m["quit_west"] = "Qw"
	m["quit_all"] = "Q"
	m["slew_north"] = "Mn"
	m["slew_south"] = "Ms"
	m["slew_east"] = "Me"
	m["slew_west"] = "Mw"
	m["slew_north_ms"] = "Mgn"
	m["slew_south_ms"] = "Mgs"
	m["slew_east_ms"] = "Mge"
	m["slew_west_ms"] = "Mgw"
	m["slew_at_radec"] = "MA"
	m["slew_at_target"] = "MS"
	return m
}
