package config

import (
	"reflect"
)

type Field struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	NoIndex bool   `yaml:"noindex"`
	IsArray bool   `yaml:"array"`
}

func (field Field) ToStructType() reflect.Type {
	var _type reflect.Type

	switch field.Type {
	case "string":
		if field.IsArray {
			_type = reflect.TypeOf([]string{""})
		} else {
			_type = reflect.TypeOf("")
		}
	case "int":
		if field.IsArray {
			_type = reflect.TypeOf([]int{0})
		} else {
			_type = reflect.TypeOf(0)
		}
	case "bool":
		if field.IsArray {
			_type = reflect.TypeOf([]bool{false})
		} else {
			_type = reflect.TypeOf(false)
		}
	default:
		_type = nil
	}

	return _type
}
