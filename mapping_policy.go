package gomapper

import (
	"errors"
	"reflect"
)

// IMappingPolicy is interface to decide how to map
type IMappingPolicy interface {
	Get(obj interface{}, key string) (string, error)
}

// DefaultMappingPolicy uses map key to map
type DefaultMappingPolicy struct{}

// Get returns key as it is
func (p DefaultMappingPolicy) Get(obj interface{}, key string) (string, error) {
	return key, nil
}

// TagMappingPolicy uses tag to map
type TagMappingPolicy struct {
	TagKey string
}

// Get returns fieldname which is the same as key
func (p TagMappingPolicy) Get(obj interface{}, key string) (string, error) {
	val := reflect.Indirect(reflect.ValueOf(obj))
	reflectType := val.Type()
	for i := 0; i < reflectType.NumField(); i++ {
		structField := reflectType.Field(i)
		extractedKey := structField.Tag.Get(p.TagKey)
		if key != extractedKey {
			continue
		}
		return structField.Name, nil
	}
	return "", errors.New("No such field")
}
