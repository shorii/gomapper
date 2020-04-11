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

var (
	refTypeString = reflect.TypeOf(string(""))

	refTypeInt   = reflect.TypeOf(int(0))
	refTypeInt8  = reflect.TypeOf(int8(0))
	refTypeInt16 = reflect.TypeOf(int16(0))
	refTypeInt32 = reflect.TypeOf(int32(0))
	refTypeInt64 = reflect.TypeOf(int64(0))

	refTypeUint   = reflect.TypeOf(uint(0))
	refTypeUint8  = reflect.TypeOf(uint8(0))
	refTypeUint16 = reflect.TypeOf(uint16(0))
	refTypeUint32 = reflect.TypeOf(uint32(0))
	refTypeUint64 = reflect.TypeOf(uint64(0))

	refTypeFloat32 = reflect.TypeOf(float32(0.0))
	refTypeFloat64 = reflect.TypeOf(float64(0.0))
)

func (m Mapper) castFieldType(name string, value interface{}, type_ reflect.Type) (*reflect.Value, error) {
	var val reflect.Value
	assertionErr := errors.New("failed to assert type")
	switch type_ {
	case refTypeString:
		v_, ok := value.(string)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeInt:
		v_, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeInt8:
		v_, ok := value.(int8)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeInt16:
		v_, ok := value.(int16)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeInt32:
		v_, ok := value.(int32)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeInt64:
		v_, ok := value.(int64)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeUint:
		v_, ok := value.(uint)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeUint8:
		v_, ok := value.(uint8)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeUint16:
		v_, ok := value.(uint16)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeUint32:
		v_, ok := value.(uint32)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeUint64:
		v_, ok := value.(uint64)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeFloat32:
		v_, ok := value.(float32)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	case refTypeFloat64:
		v_, ok := value.(float64)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(v_)
	default:
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
	}
	return &val, nil
}
