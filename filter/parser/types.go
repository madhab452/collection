package parser

type vtype int

const (
	IntType vtype = iota + 1
	StringType
	BoolType
	FloatType
	List
)

type Value struct {
	T vtype
	V interface{}
}
