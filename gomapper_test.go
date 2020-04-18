package gomapper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

type UnsafePointerVal struct{}

type StructVal struct {
	Foo int
	Bar string
}

type PtrVal struct{}

type SampleData struct {
	BoolVal          bool
	IntVal           int
	Int8Val          int8
	Int16Val         int16
	Int32Val         int32
	Int64Val         int64
	UintVal          uint
	Uint8Val         uint8
	Uint16Val        uint16
	Uint32Val        uint32
	Uint64Val        uint64
	UintptrVal       uintptr
	Complex64Val     complex64
	Complex128Val    complex128
	ArrayVal         [3]int
	StringVal        string
	ChanVal          chan int
	InterfaceVal     interface{}
	MapVal           map[string]interface{}
	PtrVal           *PtrVal
	SliceVal         []int
	StructVal1       StructVal
	StructVal2       StructVal
	UnsafePointerVal unsafe.Pointer
}

func Test_priv(t *testing.T) {
	unsafepointerval := unsafe.Pointer(&UnsafePointerVal{})
	structval1 := map[string]interface{}{"Foo": 1, "Bar": "bar"}
	structval2 := StructVal{Foo: 2, Bar: "baz"}
	chanval := make(chan int, 3)
	mapval := map[string]interface{}{"hoge": 1}
	ptrval := &PtrVal{}
	uintptrval := uintptr(unsafepointerval)
	data := map[string]interface{}{
		"BoolVal":          true,
		"IntVal":           1,
		"Int8Val":          2,
		"Int16Val":         3,
		"Int32Val":         4,
		"Int64Val":         5,
		"UintVal":          6,
		"Uint8Val":         7,
		"Uint16Val":        8,
		"Uint32Val":        9,
		"Uint64Val":        10,
		"UintptrVal":       uintptrval,
		"Complex64Val":     1i,
		"Complex128Val":    2i,
		"ArrayVal":         [3]int{11, 12, 13},
		"ChanVal":          chanval,
		"InterfaceVal":     14,
		"MapVal":           mapval,
		"PtrVal":           ptrval,
		"SliceVal":         []int{15, 16, 17},
		"StringVal":        "text",
		"StructVal1":       structval1,
		"StructVal2":       structval2,
		"UnsafePointerVal": unsafepointerval,
	}
	m := NewMapper(nil)
	actual := SampleData{}
	m.Map(data, &actual)

	expected := SampleData{
		BoolVal:          true,
		StringVal:        "text",
		IntVal:           1,
		Int8Val:          2,
		Int16Val:         3,
		Int32Val:         4,
		Int64Val:         5,
		UintVal:          6,
		Uint8Val:         7,
		Uint16Val:        8,
		Uint32Val:        9,
		Uint64Val:        10,
		UintptrVal:       uintptrval,
		Complex64Val:     1i,
		Complex128Val:    2i,
		ArrayVal:         [3]int{11, 12, 13},
		ChanVal:          chanval,
		InterfaceVal:     14,
		MapVal:           mapval,
		PtrVal:           ptrval,
		SliceVal:         []int{15, 16, 17},
		StructVal1:       StructVal{Foo: 1, Bar: "bar"},
		StructVal2:       StructVal{Foo: 2, Bar: "baz"},
		UnsafePointerVal: unsafepointerval,
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
