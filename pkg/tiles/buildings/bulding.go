package buildings

type Bulding int64

const (
	None Bulding = iota
	Monastery
)

func (building Bulding) String() string {

	switch building {
	case None:
		return "NONE_BUILDING"
	case Monastery:
		return "MONASTERY"
	default:
		return "ERROR"
	}
}
