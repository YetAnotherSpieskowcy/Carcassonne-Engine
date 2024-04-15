package building

type Building int64

const (
	None Building = iota
	Monastery
)

func (building Building) String() string {

	switch building {
	case None:
		return "NONE_BUILDING"
	case Monastery:
		return "MONASTERY"
	default:
		return "ERROR"
	}
}
