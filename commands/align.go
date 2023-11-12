package commands

import (
	"github.com/ddefrancesco/scoperunner_server/scopeparser"
)

func NewAlignCommand(com scopeparser.AlignCommandValue) *AlignCommand {
	command := &AlignCommand{
		cmd: com,
	}
	return command
}

type AlignCommand struct {
	cmd scopeparser.AlignCommandValue
}

func (c *AlignCommand) ParseCommand() string {
	switch c.cmd {
	case scopeparser.AltAzCmd:
		return ":AA#"
	case scopeparser.PolarCmd:
		return ":AP#"
	case scopeparser.LandCmd:
		return ":AL#"

	}
	return ""
}
