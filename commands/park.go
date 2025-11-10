package commands

import "github.com/ddefrancesco/scoperunner_server/scopeparser"

func NewParkingCommand(com scopeparser.ParkCommand) *ParkingCommand {
	command := &ParkingCommand{
		cmd: com,
	}
	return command
}

type ParkingCommand struct {
	cmd scopeparser.ParkCommand
}

func (c *ParkingCommand) ParseCommand() string {
	switch c.cmd {
	case scopeparser.ParkCmdPark:
		return ":hP#"
	case scopeparser.ParkCmdSeek:
		return ":hF#"
	case scopeparser.ParkCmdQuery:
		return ":h?#"

	}
	return ""
}
