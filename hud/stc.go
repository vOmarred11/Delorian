package hud

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/vOmarred11/Delorian/hud/cheat"
	"github.com/vOmarred11/Delorian/hud/cmd"
)

type Cmd struct {
	Command   []protocol.Command
	Commandvl []protocol.CommandOrigin
	Commandv  []packet.CommandRequest
}
type Cheat struct {
	Fly         CheatTypeFly
	Reach       CheatTypeReach
	Autoclicker CheatTypeAutoclicker
	Automine    CheatTypeAutomine
}
type AddCommand struct {
	Name             string
	Description      string
	Alias            any
	Replace          bool
	CommandTypeCheat bool
	CommandType      Cmd
}
type AddCheat struct {
	Name             string
	Description      string
	Alias            any
	Replace          bool
	CommandTypeCheat bool
	CheatType        Cheat
}
type CheatTypeReach struct {
	Enabled bool
	Command cmd.Command
	Value   int32
}
type CheatTypeFly struct {
	Enabled bool
	Command cmd.Command
	Value   int32
}
type CheatTypeAutoclicker struct {
	Enabled bool
	Command cmd.Command
	Range   int
	Hold    bool
}
type CheatTypeAutomine struct {
	Command cmd.Command
	Flag    byte
	Enabled bool
}

func (t Cmd) command() protocol.Command { return protocol.Command{} }

func (t Cmd) commandvl() protocol.CommandOrigin { return protocol.CommandOrigin{} }

func (t Cmd) commandv() packet.CommandRequest { return packet.CommandRequest{} }

func (t Cmd) cmd() packet.AvailableCommands { return packet.AvailableCommands{} }

func (c AddCommand) AddCommand(CommandType cmd.Command) {
	AddCmd := true
	if AddCmd {
		panic("should never happen at least someone changes values")
	}
}
func (c AddCheat) AddCheat(CheatType cheat.CheatType) {

}
