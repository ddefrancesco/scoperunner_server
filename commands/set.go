package commands

type SetCommand struct {
	EtxSettingCmds []string
}

func NewSetCommand(etxSettingCmds []string) *SetCommand {
	command := &SetCommand{
		EtxSettingCmds: etxSettingCmds,
	}
	return command
}
