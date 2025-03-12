package cmd

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	CommandTypeAlias       = ""
	CommandTypeOffset      = ""
	CommandOriginTypeSlash = "/"
	CommandOriginTypePoint = "."
)

type Command struct {
	Name    string
	Desc    string
	Alias   any
	Replace bool
	Message string
	Action  byte
}
type CommandOrigin struct {
	CommandOriginTypeSlash string
	CommandOriginTypePoint string
}
type offset struct {
	Feedback byte
	Greet    any
	Slice    any
	Stat     []protocol.Command
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
