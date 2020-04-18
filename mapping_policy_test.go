package gomapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SampleStruct struct {
	Foo int `test:"footag"`
}

func Test_DefaultMappingPolicy(t *testing.T) {
	obj := SampleStruct{}
	mappingPolicy := DefaultMappingPolicy{}
	actual, err := mappingPolicy.Get(&obj, "Foo")
	if err != nil {
		t.Error(err)
	}
	expected := "Foo"
	assert.Equal(t, expected, actual)
}

func Test_TagMappingPolicy(t *testing.T) {
	obj := SampleStruct{}
	mappingPolicy := TagMappingPolicy{TagKey: "test"}
	actual, err := mappingPolicy.Get(&obj, "footag")
	if err != nil {
		t.Error(err)
	}
	expected := "Foo"
	assert.Equal(t, expected, actual)
}
