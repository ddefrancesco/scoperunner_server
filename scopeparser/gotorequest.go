package scopeparser

import (
	"errors"
	"log"
	"strings"

	"github.com/ddefrancesco/scoperunner_server/openngc/cache"
)

var ErrUnknownCommand = errors.New("GoTo Command Failed")

type GotoRequest struct {
	Goto map[string]string
}

func NewGotoRequest(m map[string]string) *GotoRequest {
	request := &GotoRequest{
		Goto: m,
	}
	return request
}

type ScopeGotoCmd struct {
	ValuesMap map[string]string
}

func (s *GotoRequest) ParseMap() (map[string]string, error) {

	settings := s.InitMap()
	var settingsMap ScopeMoveCmd
	settingsMap.ValuesMap = make(map[string]string)
	for k, v := range s.Goto {
		if _, ok := settings[k]; ok {

			settingsMap.ValuesMap[settings[k]] = v

		} else {

			return nil, ErrUnknownMoveCommand
		}
	}

	return settingsMap.ValuesMap, nil
}

func (s *GotoRequest) InitMap() map[string]string {
	return gotoCmdDictionary()
}

func gotoCmdDictionary() map[string]string {

	m := make(map[string]string)

	m["target_dec"] = "Sd"
	m["target_ra"] = "Sr"
	m["slew_at_target"] = "MS"
	m["slew_at_radec"] = "MA"

	return m
}

func (s *GotoRequest) SetGotoRADecCommand() (string, error) {
	gotoReq := s.Goto

	openngc_catalog := cache.NewNGCCatalog()
	openngc_record, err := openngc_catalog.FindNGCObject(gotoReq["goto"])
	//openngc_record, err := openngc_catalog.FindNGCObject("M31")
	if err != nil {
		return "", err
	}
	dec := openngc_record.Dec[:9]
	dec = strings.Replace(dec, ":", "*", 1)

	ra := openngc_record.RA[:8]
	log.Println("Object " + gotoReq["goto"] + " found in catalog")
	log.Println("RA: " + ra)
	log.Println("DEC: " + dec)
	gotoCmd := ":" + s.InitMap()["target_ra"] + ra + "#:" + s.InitMap()["target_dec"] + dec + "#"
	return gotoCmd, nil
}

func (s *GotoRequest) CheckGotoRADecCommand() (string, error) {
	return ":" + s.InitMap()["slew_at_target"] + "#", nil

}

func (s *GotoRequest) SetGotoCommand() (string, error) {
	return ":" + s.InitMap()["slew_at_radec"] + "#", nil

}
