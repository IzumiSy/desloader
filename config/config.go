package config

import (
	"io/ioutil"

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

type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
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
