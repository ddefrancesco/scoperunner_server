package scopeparser

import "fmt"

type ParkType string

const (
	ParkTypeQuery ParkType = "query"
	ParkTypeSeek  ParkType = "seek"
	ParkTypePark  ParkType = "park"
)

type ParkCommand string

const (
	ParkCmdSeek  ParkCommand = ":hF#"
	ParkCmdPark  ParkCommand = ":hP#"
	ParkCmdQuery ParkCommand = ":h?#"
)

type Parking struct {
	parkType ParkType
}

func NewParking(m ParkType) *Parking {
	park := &Parking{
		parkType: m,
	}
	return park
}

func (p *Parking) InitMap() map[ParkType]ParkCommand {
	items := make(map[ParkType]ParkCommand)
	items[ParkTypeQuery] = ParkCmdQuery
	items[ParkTypeSeek] = ParkCmdSeek
	items[ParkTypePark] = ParkCmdPark
	return items
}

func (p *Parking) ParseMap() (ParkCommand, error) {
	aMap := p.InitMap()
	if _, ok := aMap[p.parkType]; ok {
		return aMap[p.parkType], nil
	}
	return "error", fmt.Errorf("unknown alignment")
}
