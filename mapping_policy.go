package gomapper

import (
	"errors"
	"reflect"
)

type IMappingPolicy interface {
	Get(obj interface{}, key string) (string, error)
}

type DefaultMappingPolicy struct{}

func (p DefaultMappingPolicy) Get(obj interface{}, key string) (string, error) {
	return key, nil
}

type TagMappingPolicy struct {
	TagKey string
}

func (p TagMappingPolicy) Get(obj interface{}, key string) (string, error) {
	val := reflect.Indirect(reflect.ValueOf(obj))
	type_ := val.Type()
	for i := 0; i < type_.NumField(); i++ {
		structField := type_.Field(i)
		extractedKey := structField.Tag.Get(p.TagKey)
		if key != extractedKey {
			continue
		}
		return structField.Name, nil
	}
	return "", errors.New("No such field")
}
