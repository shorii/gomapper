package main

import (
	"fmt"
	"gomapper"
)

type Sample struct {
	Foo string `mapper:"foo"`
	Bar int    `mapper:"bar"`
}

func main() {
	policy := gomapper.TagNamingPolicy{TagKey: "mapper"}
	m := gomapper.NewMapper(policy)
	testData := map[string]interface{}{
		"foo": "0.0.0.0",
		"bar": 7890,
	}

	sample := Sample{}
	err := m.Map(testData, &sample)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", sample)
}
