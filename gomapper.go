package gomapper

import (
	"errors"
	"fmt"
	"reflect"
)

// Mapper map to struct
type Mapper struct {
	mappingPolicy IMappingPolicy
}

// NewMapper create new Mapper
func NewMapper(policy IMappingPolicy) *Mapper {
	mappingPolicy := policy
	if policy == nil {
		mappingPolicy = DefaultMappingPolicy{}
	}
	return &Mapper{mappingPolicy: mappingPolicy}
}

// Map map data to struct
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

func (m Mapper) castFieldType(name string, value interface{}, reflectType reflect.Type) (*reflect.Value, error) {
	var val reflect.Value
	assertionErr := errors.New("failed to assert type")
	kind := reflectType.Kind()
	switch kind {
	case reflect.Int8:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(int8(castedVal))
	case reflect.Int16:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(int16(castedVal))
	case reflect.Int32:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(int32(castedVal))
	case reflect.Int64:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(int64(castedVal))
	case reflect.Uint:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint(castedVal))
	case reflect.Uint8:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint8(castedVal))
	case reflect.Uint16:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint16(castedVal))
	case reflect.Uint32:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint32(castedVal))
	case reflect.Uint64:
		castedVal, ok := value.(int)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(uint64(castedVal))
	case reflect.Complex64:
		castedVal, ok := value.(complex128)
		if !ok {
			return nil, assertionErr
		}
		val = reflect.ValueOf(complex64(castedVal))
	case reflect.Struct:
		mapData, ok := value.(map[string]interface{})
		if ok {
			newValue, err := m.castStruct(mapData, reflectType)
			if err != nil {
				return nil, assertionErr
			}
			val = *newValue
		} else {
			val = reflect.ValueOf(value)
			if val.Kind() != reflect.Struct {
				return nil, errors.New("Invalid value type")
			}
		}
	case reflect.Array:
		arrayVal := reflect.ValueOf(value)
		items := reflect.New(reflectType).Elem()
		if arrayVal.Len() != items.Len() {
			return nil, errors.New("Invalid value type")
		}
		structType := reflectType.Elem()
		for i := 0; i < items.Len(); i++ {
			element := arrayVal.Index(i)
			mapData, ok := element.Interface().(map[string]interface{})
			if ok {
				newValue, err := m.castStruct(mapData, structType)
				if err != nil {
					return nil, assertionErr
				}
				items.Index(i).Set(*newValue)
			} else {
				items.Index(i).Set(element)
			}
		}
		val = items
	case reflect.Slice:
		sliceVal := reflect.ValueOf(value)
		slice := reflect.MakeSlice(reflectType, 0, 0)
		structType := reflectType.Elem()
		for i := 0; i < sliceVal.Cap(); i++ {
			element := sliceVal.Index(i)
			mapData, ok := element.Interface().(map[string]interface{})
			if ok {
				newValue, err := m.castStruct(mapData, structType)
				if err != nil {
					return nil, assertionErr
				}
				slice = reflect.Append(slice, *newValue)
			} else {
				slice = reflect.Append(slice, element)
			}
		}
		val = slice
	default:
		val = reflect.ValueOf(value)
	}

	return &val, nil
}

func (m Mapper) castStruct(mapData map[string]interface{}, structType reflect.Type) (*reflect.Value, error) {
	newValue := reflect.New(structType).Elem()
	for mapKey, mapValue := range mapData {
		name, err := m.mappingPolicy.Get(newValue.Interface(), mapKey)
		if err != nil {
			return nil, err
		}
		fieldValue := newValue.FieldByName(name)
		fieldType := fieldValue.Type()
		refVal, err := m.castFieldType(name, mapValue, fieldType)
		if err != nil {
			return nil, err
		}
		fieldValue.Set(*refVal)
	}
	return &newValue, nil
}
