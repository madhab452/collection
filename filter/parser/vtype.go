package parser

type ValueType string

const (
	TypeBool   ValueType = "bool"
	TypeFloat  ValueType = "float64"
	TypeInt    ValueType = "int64"
	TypeString ValueType = "string"
	TypeAny    ValueType = "any" // let runtime handle such type
)

// todo: use generics
type Value struct {
	Literal     string
	ValueType   ValueType
	BoolValue   bool
	StringValue string
	// more value will be supported later
}

func newValue(vt ValueType, v string) *Value {
	switch vt {
	case TypeBool:
		return boolValue(v)
	case TypeString:
		return &Value{
			ValueType:   TypeString,
			Literal:     v,
			StringValue: v,
		}
	}

	return nil
}

func boolValue(v string) *Value {
	return &Value{
		ValueType: TypeBool,
		Literal:   v,
		BoolValue: func() bool {
			return v == "true"
		}(),
	}
}
