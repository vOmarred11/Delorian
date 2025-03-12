package main

import (
	"github.com/vOmarred11/Delorian/hud"
	"github.com/vOmarred11/Delorian/hud/cheat"
	"github.com/vOmarred11/Delorian/hud/cmd"
)

func main() {
	h := hud.AddCommand{}
	c := hud.AddCheat{}
	h.AddCommand(cmd.Command{Name: "", Desc: "", Alias: ""})
	c.AddCheat(cheat.CheatType{Reach: make(map[int]bool)})
}
