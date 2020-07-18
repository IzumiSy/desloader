package schema

import (
	"reflect"
)

type StructBuilder struct {
	internal reflect.Type
	index    map[string]int
}

type Instance struct {
	internal reflect.Value
	index    map[string]int
}

func (builder StructBuilder) Instantiate() Instance {
	return Instance{
		internal: reflect.New(builder.internal).Elem(),
		index:    builder.index,
	}
}
