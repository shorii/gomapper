package main

import (
	"fmt"
	"gomapper"
)

type Nested struct {
	Baz int `mapper:"bazkey"`
	Qux int `mapper:"quxkey"`
}

type Sample struct {
	Foo    string `mapper:"fookey"`
	Bar    int    `mapper:"barkey"`
	Nested Nested `mapper:"nestedkey"`
}

func main() {
	policy := gomapper.TagMappingPolicy{TagKey: "mapper"}
	m := gomapper.NewMapper(policy)
	testData := map[string]interface{}{
		"fookey": "foo value",
		"barkey": 9999,
		"nestedkey": map[string]interface{}{
			"bazkey": 9,
			"quxkey": 1,
		},
	}

	sample := Sample{
		Nested: Nested{},
	}
	err := m.Map(testData, &sample)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", sample)
}
