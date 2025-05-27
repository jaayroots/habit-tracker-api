package enums

type Frequency int

const (
	Daily Frequency = iota + 1
	Weekly
)

func (s Frequency) String() string {
	switch s {
	case Daily:
		return "new"
	case Weekly:
		return "pending"
	default:
		return "unknow"
	}
}

func FetfrequencyMap() map[int]string {
	return map[int]string{
		int(Daily):  Daily.String(),
		int(Weekly): Weekly.String(),
	}
}

func FsValidfrequency(value int) bool {
	switch Frequency(value) {
	case Daily, Weekly:
		return true
	default:
		return false
	}
}
