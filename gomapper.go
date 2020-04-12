package gomapper

import (
	"errors"
	"fmt"
	"reflect"
)

type Mapper struct {
	mappingPolicy IMappingPolicy
}

func NewMapper(policy IMappingPolicy) *Mapper {
	mappingPolicy := policy
	if policy == nil {
		mappingPolicy = DefaultMappingPolicy{}
	}
	return &Mapper{mappingPolicy: mappingPolicy}
}

func (m Mapper) Map(data map[string]interface{}, obj interface{}) error {
	for key, value := range data {
		name, err := m.mappingPolicy.Get(obj, key)
		if err != nil {
			return err
		}
		m.setField(obj, name, value)
	}
	return nil
}

func (m Mapper) setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj)
	if structValue.Kind() == reflect.Ptr {
		structValue = structValue.Elem()
	}
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()

	val, err := m.castFieldType(name, value, structFieldType)
	if err != nil {
		return err
	}

	structFieldValue.Set(*val)

	return nil
}

func (m Mapper) castFieldType(name string, value interface{}, type_ reflect.Type) (*reflect.Value, error) {
	var val reflect.Value
	assertionErr := errors.New("failed to assert type")
	kind := type_.Kind()
	switch kind {
	case reflect.Int8:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(int8(v_))
	case reflect.Int16:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(int16(v_))
	case reflect.Int32:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(int32(v_))
	case reflect.Int64:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(int64(v_))
	case reflect.Uint:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint(v_))
	case reflect.Uint8:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint8(v_))
	case reflect.Uint16:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint16(v_))
	case reflect.Uint32:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint32(v_))
	case reflect.Uint64:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint64(v_))
	case reflect.Complex64:
		v_, ok := value.(complex128)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(complex64(v_))
	case reflect.Struct:
		mv_, ok := value.(map[string]interface{})
		if !ok {
			return nil, assertionErr
		}
		v_ := reflect.New(type_).Elem()
		for key, val := range mv_ {
			name, err := m.mappingPolicy.Get(v_.Interface(), key)
			if err != nil {
				return nil, err
			}
			fieldValue := v_.FieldByName(name)
			fieldType := fieldValue.Type()
			refVal, err := m.castFieldType(name, val, fieldType)
			if err != nil {
				return nil, err
			}
			fieldValue.Set(*refVal)
		}
		val = v_
	default:
		val = reflect.ValueOf(value)
	}

	return &val, nil
}
