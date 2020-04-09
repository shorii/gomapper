package gomapper

import (
	"errors"
	"reflect"
)

type INamingPolicy interface {
	Get(obj interface{}, key string) (string, error)
}

type TagNamingPolicy struct {
	TagKey string
}

func (p TagNamingPolicy) Get(obj interface{}, key string) (string, error) {
	valObj := reflect.ValueOf(obj)
	val := valObj.Elem()
	for i := 0; i < val.NumField(); i++ {
		structField := val.Type().Field(i)
		extractedKey := structField.Tag.Get(p.TagKey)
		if key != extractedKey {
			continue
		}
		return structField.Name, nil
	}
	return "", errors.New("No such field")
}
