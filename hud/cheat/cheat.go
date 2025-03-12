package cheat

type CheatType struct {
	Fly         map[int]bool
	Automine    map[byte]bool
	Autoclicker map[byte]bool
	Reach       map[int]bool
	Response    string
	SubCmd      string
}
