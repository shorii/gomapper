package main

import (
	"fmt"
	"gomapper"
)

type Nested struct {
	Baz int `mapper:baz`
	Qux int `mapper:qux`
}

type Sample struct {
	Foo    string `mapper:"foo"`
	Bar    int    `mapper:"bar"`
	Nested Nested `mapper:"nested"`
}

func main() {
	policy := gomapper.TagNamingPolicy{TagKey: "mapper"}
	m := gomapper.NewMapper(policy)
	testData := map[string]interface{}{
		"Foo": "foo value",
		"Bar": 9999,
		"Nested": map[string]interface{}{
			"Baz": 9,
			"Qux": 1,
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
