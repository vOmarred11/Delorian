package cmd

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	CommandTypeAlias  = ""
	CommandTypeOffset = ""
)

type Command struct {
	name  string
	desc  string
	alias any
}
type offset struct {
	feedback byte
	greet    any
	slice    any
	stat     []protocol.Command
}
type hts interface {
	commandvl() []protocol.Command
	offset() offset
}

func (c *Command) commandvl() []protocol.Command {
	return []protocol.Command{}
}
func (v *offset) commandv() []protocol.Command {
	return nil
}
