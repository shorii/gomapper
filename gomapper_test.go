package gomapper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type BoolData struct {
	BoolVal bool
}

func Test_bool(t *testing.T) {
	data := map[string]interface{}{
		"BoolVal": true,
	}
	m := NewMapper(nil)
	actual := BoolData{}
	m.Map(data, &actual)

	expected := BoolData{
		BoolVal: true,
	}

	assert.Equal(t, expected, actual)
}

type StringData struct {
	StringVal string
}

func Test_string(t *testing.T) {
	data := map[string]interface{}{
		"StringVal": "text",
	}
	m := NewMapper(nil)
	actual := StringData{}
	m.Map(data, &actual)

	expected := StringData{
		StringVal: "text",
	}

	assert.Equal(t, expected, actual)
}

type IntData struct {
	IntVal int
}

func Test_int(t *testing.T) {
	data := map[string]interface{}{
		"IntVal": 1,
	}
	m := NewMapper(nil)
	actual := IntData{}
	m.Map(data, &actual)

	expected := IntData{
		IntVal: 1,
	}

	assert.Equal(t, expected, actual)
}

type Int8Data struct {
	Int8Val int8
}

func Test_int8(t *testing.T) {
	data := map[string]interface{}{
		"Int8Val": 1,
	}
	m := NewMapper(nil)
	actual := Int8Data{}
	m.Map(data, &actual)

	expected := Int8Data{
		Int8Val: 1,
	}

	assert.Equal(t, expected, actual)
}

type Int16Data struct {
	Int16Val int16
}

func Test_int16(t *testing.T) {
	data := map[string]interface{}{
		"Int16Val": 1,
	}
	m := NewMapper(nil)
	actual := Int16Data{}
	m.Map(data, &actual)

	expected := Int16Data{
		Int16Val: 1,
	}

	assert.Equal(t, expected, actual)
}

type Int32Data struct {
	Int32Val int32
}

func Test_int32(t *testing.T) {
	data := map[string]interface{}{
		"Int32Val": 1,
	}
	m := NewMapper(nil)
	actual := Int32Data{}
	m.Map(data, &actual)

	expected := Int32Data{
		Int32Val: 1,
	}

	assert.Equal(t, expected, actual)
}

type Int64Data struct {
	Int64Val int64
}

func Test_int64(t *testing.T) {
	data := map[string]interface{}{
		"Int64Val": 1,
	}
	m := NewMapper(nil)
	actual := Int64Data{}
	m.Map(data, &actual)

	expected := Int64Data{
		Int64Val: 1,
	}

	assert.Equal(t, expected, actual)
}

type UintData struct {
	UintVal uint
}

func Test_uint(t *testing.T) {
	data := map[string]interface{}{
		"UintVal": 1,
	}
	m := NewMapper(nil)
	actual := UintData{}
	m.Map(data, &actual)

	expected := UintData{
		UintVal: 1,
	}

	assert.Equal(t, expected, actual)
}

type Uint8Data struct {
	Uint8Val uint8
}

func Test_uint8(t *testing.T) {
	data := map[string]interface{}{
		"Uint8Val": 1,
	}
	m := NewMapper(nil)
	actual := Uint8Data{}
	m.Map(data, &actual)

	expected := Uint8Data{
		Uint8Val: 1,
	}

	assert.Equal(t, expected, actual)
}

type Uint16Data struct {
	Uint16Val uint16
}

func Test_uint16(t *testing.T) {
	data := map[string]interface{}{
		"Uint16Val": 1,
	}
	m := NewMapper(nil)
	actual := Uint16Data{}
	m.Map(data, &actual)

	expected := Uint16Data{
		Uint16Val: 1,
	}

	assert.Equal(t, expected, actual)
}

type Uint32Data struct {
	Uint32Val uint32
}

func Test_uint32(t *testing.T) {
	data := map[string]interface{}{
		"Uint32Val": 1,
	}
	m := NewMapper(nil)
	actual := Uint32Data{}
	m.Map(data, &actual)

	expected := Uint32Data{
		Uint32Val: 1,
	}

	assert.Equal(t, expected, actual)
}

type Uint64Data struct {
	Uint64Val uint64
}

func Test_uint64(t *testing.T) {
	data := map[string]interface{}{
		"Uint64Val": 1,
	}
	m := NewMapper(nil)
	actual := Uint64Data{}
	m.Map(data, &actual)

	expected := Uint64Data{
		Uint64Val: 1,
	}

	assert.Equal(t, expected, actual)
}

type Complex64Data struct {
	Complex64Val complex64
}

func Test_complex64(t *testing.T) {
	data := map[string]interface{}{
		"Complex64Val": 2i,
	}
	m := NewMapper(nil)
	actual := Complex64Data{}
	m.Map(data, &actual)

	expected := Complex64Data{
		Complex64Val: 2i,
	}

	assert.Equal(t, expected, actual)
}

type Complex128Data struct {
	Complex128Val complex128
}

func Test_complex128(t *testing.T) {
	data := map[string]interface{}{
		"Complex128Val": 2i,
	}
	m := NewMapper(nil)
	actual := Complex128Data{}
	m.Map(data, &actual)

	expected := Complex128Data{
		Complex128Val: 2i,
	}

	assert.Equal(t, expected, actual)
}

type NestedStructData struct {
	Foo string
}

type ArrayData struct {
	ArrayVal              [3]int
	ArrayStructFromStruct [3]NestedStructData
	ArrayStructFromMap    [3]NestedStructData
}

func Test_array(t *testing.T) {
	data := map[string]interface{}{
		"ArrayVal": [3]int{1, 2, 3},
		"ArrayStructFromStruct": [3]NestedStructData{
			NestedStructData{Foo: "fooval1"},
			NestedStructData{Foo: "fooval2"},
			NestedStructData{Foo: "fooval3"},
		},
		"ArrayStructFromMap": [3]map[string]interface{}{
			map[string]interface{}{"Foo": "fooval7"},
			map[string]interface{}{"Foo": "fooval8"},
			map[string]interface{}{"Foo": "fooval9"},
		},
	}
	m := NewMapper(nil)
	actual := ArrayData{}
	m.Map(data, &actual)
	expected := ArrayData{
		ArrayVal: [3]int{1, 2, 3},
		ArrayStructFromStruct: [3]NestedStructData{
			NestedStructData{Foo: "fooval1"},
			NestedStructData{Foo: "fooval2"},
			NestedStructData{Foo: "fooval3"},
		},
		ArrayStructFromMap: [3]NestedStructData{
			NestedStructData{Foo: "fooval7"},
			NestedStructData{Foo: "fooval8"},
			NestedStructData{Foo: "fooval9"},
		},
	}
	assert.Equal(t, expected, actual)
}

type ChanData struct {
	ChanVal chan int
}

func Test_chan(t *testing.T) {
	chanval := make(chan int, 3)
	data := map[string]interface{}{
		"ChanVal": chanval,
	}
	m := NewMapper(nil)
	actual := ChanData{}
	m.Map(data, &actual)

	expected := ChanData{
		ChanVal: chanval,
	}

	assert.Equal(t, expected, actual)
}

type InterfaceData struct {
	InterfaceVal interface{}
}

func Test_interface(t *testing.T) {
	data := map[string]interface{}{
		"InterfaceVal": 1,
	}
	m := NewMapper(nil)
	actual := InterfaceData{}
	m.Map(data, &actual)

	expected := InterfaceData{
		InterfaceVal: 1,
	}

	assert.Equal(t, expected, actual)
}

type MapData struct {
	MapVal interface{}
}

func Test_map(t *testing.T) {
	data := map[string]interface{}{
		"MapVal": map[string]interface{}{
			"foo": "fooval",
			"bar": "barval",
		},
	}
	m := NewMapper(nil)
	actual := MapData{}
	m.Map(data, &actual)

	expected := MapData{
		MapVal: map[string]interface{}{
			"foo": "fooval",
			"bar": "barval",
		},
	}

	assert.Equal(t, expected, actual)
}

type PtrData struct {
	PtrVal *NestedStructData
}

func Test_ptr(t *testing.T) {
	ptrval := &NestedStructData{
		Foo: "fooval",
	}
	data := map[string]interface{}{
		"PtrVal": ptrval,
	}
	m := NewMapper(nil)
	actual := PtrData{}
	m.Map(data, &actual)

	expected := PtrData{
		PtrVal: ptrval,
	}

	assert.Equal(t, expected, actual)
}

type SliceData struct {
	SliceVal              []int
	SliceStructFromStruct []NestedStructData
	SliceStructFromMap    []NestedStructData
}

func Test_slice(t *testing.T) {
	data := map[string]interface{}{
		"SliceVal": []int{1, 2, 3},
		"SliceStructFromStruct": []NestedStructData{
			NestedStructData{Foo: "fooval1"},
			NestedStructData{Foo: "fooval2"},
			NestedStructData{Foo: "fooval3"},
		},
		"SliceStructFromMap": []map[string]interface{}{
			map[string]interface{}{"Foo": "fooval7"},
			map[string]interface{}{"Foo": "fooval8"},
			map[string]interface{}{"Foo": "fooval9"},
		},
	}
	m := NewMapper(nil)
	actual := SliceData{}
	m.Map(data, &actual)
	expected := SliceData{
		SliceVal: []int{1, 2, 3},
		SliceStructFromStruct: []NestedStructData{
			NestedStructData{Foo: "fooval1"},
			NestedStructData{Foo: "fooval2"},
			NestedStructData{Foo: "fooval3"},
		},
		SliceStructFromMap: []NestedStructData{
			NestedStructData{Foo: "fooval7"},
			NestedStructData{Foo: "fooval8"},
			NestedStructData{Foo: "fooval9"},
		},
	}
	assert.Equal(t, expected, actual)
}

type StructData struct {
	StructValFromStruct NestedStructData
	StructValFromMap    NestedStructData
}

func Test_struct(t *testing.T) {
	data := map[string]interface{}{
		"StructValFromStruct": NestedStructData{
			Foo: "fooval1",
		},
		"StructValFromMap": map[string]interface{}{
			"Foo": "fooval2",
		},
	}
	m := NewMapper(nil)
	actual := StructData{}
	m.Map(data, &actual)

	expected := StructData{
		StructValFromStruct: NestedStructData{
			Foo: "fooval1",
		},
		StructValFromMap: NestedStructData{
			Foo: "fooval2",
		},
	}

	assert.Equal(t, expected, actual)
}

type UnsafePtrData struct {
	UnsafePtrVal unsafe.Pointer
}

func Test_unsafeptr(t *testing.T) {
	unsafeptrval := unsafe.Pointer(&NestedStructData{})
	data := map[string]interface{}{
		"UnsafePtrVal": unsafeptrval,
	}
	m := NewMapper(nil)
	actual := UnsafePtrData{}
	m.Map(data, &actual)

	expected := UnsafePtrData{
		UnsafePtrVal: unsafeptrval,
	}

	assert.Equal(t, expected, actual)
}

type FuncData struct {
	FuncVal func() error
}

func Test_func(t *testing.T) {
	funcval := func() error { return nil }
	data := map[string]interface{}{
		"FuncVal": funcval,
	}
	m := NewMapper(nil)
	actual := FuncData{}
	m.Map(data, &actual)
	printedActual := fmt.Sprintf("%#v", actual)

	expected := FuncData{
		FuncVal: funcval,
	}
	printedExpected := fmt.Sprintf("%#v", expected)

	assert.Equal(t, printedExpected, printedActual)
}
