package dash

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type cout struct {
	command   []protocol.Command
	commandvl []protocol.CommandOrigin
	commandv  []packet.CommandRequest
}
type cstamp struct {
	//response to the protocol
	response []protocol.CommandOutputMessage
	//response to the client
	client []packet.Text
	//response to the console
	console string
}
type stamp struct {
	cout cstamp
}
