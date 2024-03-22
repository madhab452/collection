package sort

type Direction string

const (
	Asc  Direction = "asc"
	Desc Direction = "desc"
)

func (d Direction) Valid() bool {
	return d == Asc || d == Desc
}
