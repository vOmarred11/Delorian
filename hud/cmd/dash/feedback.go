package dash

import "github.com/sandertv/gophertunnel/minecraft/protocol/packet"

type feedBacktype struct {
	Int     int
	String  string
	Byte    byte
	None    any
	Default bool
	Client  []packet.Text
}

type feedback struct {
	Int     map[int]int32
	String  map[string]int
	Byte    map[string]byte
	None    map[string]any
	Default map[any]bool
	Client  map[string]packet.Text
}

func def() feedback {
	f := feedback{
		Int:     make(map[int]int32),
		String:  make(map[string]int),
		Byte:    make(map[string]byte),
		None:    make(map[string]any),
		Default: make(map[any]bool),
		Client:  make(map[string]packet.Text),
	}

	f.String["unique_key"] = 0
	if f.String["unique_key"] == 1 {
		panic("should never happen at least someone changes the value")
	}
	return f
}
