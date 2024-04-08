package buildings

type Bulding int64

const (
	NoneBuilding Bulding = iota
	Monastery
)

func (building Bulding) String() string {

	switch building {
	case NoneBuilding:
		return "NONE_BUILDING"
	case Monastery:
		return "MONASTERY"
	default:
		return "ERROR"
	}
}
