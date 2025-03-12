package hud

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type cmd struct {
	Command   []protocol.Command
	Commandvl []protocol.CommandOrigin
	Commandv  []packet.CommandRequest
}

func (t cmd) command() protocol.Command { return protocol.Command{} }

func (t cmd) commandvl() protocol.CommandOrigin { return protocol.CommandOrigin{} }

func (t cmd) commandv() packet.CommandRequest { return packet.CommandRequest{} }
