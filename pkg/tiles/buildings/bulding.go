package buildings

type Bulding int64

const (
	NONE_BULDING Bulding = iota
	MONASTERY
)

func (building Bulding) ToString() string {

	switch building {
	case NONE_BULDING:
		return "NONE_BUILDING"
	case MONASTERY:
		return "MONASTERY"
	default:
		return "ERROR"
	}
}
