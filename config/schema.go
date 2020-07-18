package config

import (
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Schema struct {
	Tables []Table `yaml:"tables"`
}

type Table struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
}

func (schema Schema) ToBuilder() map[string]StructBuilder {
	definitions := map[string]StructBuilder{}

	for _, table := range schema.Tables {
		fields := []reflect.StructField{}
		for _, field := range table.Fields {
			fields = append(fields, reflect.StructField{
				Name: fmt.Sprintf("FIELD_%s", field.Name),
				Type: field.ToStructType(),
				Tag:  reflect.StructTag(fmt.Sprintf("datastore:\"%s\"", field.Name)),
			})
		}

		internal := reflect.StructOf(fields)
		index := make(map[string]int)
		for i := 0; i < internal.NumField(); i++ {
			index[internal.Field(i).Name] = i
		}

		definitions[table.Name] = StructBuilder{
			internal: internal,
			index:    index,
		}
	}

	return definitions
}

func Load(path string) (Schema, error) {
	if path == "" {
		return Schema{}, errors.New("Couldn't load schema YAML")
	}

	yamlBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Schema{}, err
	}

	schema := Schema{}
	err = yaml.UnmarshalStrict(yamlBytes, &schema)
	return schema, err
}
