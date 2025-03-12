package dash

import "github.com/sandertv/gophertunnel/minecraft/protocol"

type offset struct {
	alias       protocol.Command
	aliasstring string
	msg         string
	offset      byte
	msgimport   any
	action      byte
}
