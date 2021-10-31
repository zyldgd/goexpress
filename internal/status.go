package internal

type Status int

const (
	StatusUnknown Status = iota
	StatusString
	StatusVariable
	StatusNumeric
)

func readVariable() {

}
