package dash

const (
	ModeTypeDial = iota
	ModeTypeListen
	ModeTypeRanket
	ModeTypeGoRanket
)

type dash struct {
	message string
	cout    cout
	stamp   stamp
	mode    byte
}
