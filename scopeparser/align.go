package scopeparser

import "fmt"

type AlignMode string
type AlignCommandValue string

const (
	AltAz AlignMode = "altaz"
	Polar AlignMode = "polar"
	Land  AlignMode = "land"
)

const (
	AltAzCmd AlignCommandValue = ":AA#"
	PolarCmd AlignCommandValue = ":AP#"
	LandCmd  AlignCommandValue = ":AL#"
)

type Alignment struct {
	mode AlignMode
}

func NewAlignment(m AlignMode) *Alignment {
	alignment := &Alignment{
		mode: m,
	}
	return alignment
}

func initItems() map[AlignMode]AlignCommandValue {
	items := make(map[AlignMode]AlignCommandValue)
	items[AltAz] = AltAzCmd
	items[Polar] = PolarCmd
	items[Land] = LandCmd
	return items
}

func (p *Alignment) ParseMap() (AlignCommandValue, error) {
	aMap := initItems()
	if _, ok := aMap[p.mode]; ok {
		return aMap[p.mode], nil
	}
	return "error", fmt.Errorf("unknown alignment")
}
