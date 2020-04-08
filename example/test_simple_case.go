package main

import (
	"fmt"
	"mapper"
)

type HttpProxy struct {
	Host string `mapper:"host"`
	Port int    `mapper:"port"`
}

func main() {
	policy := mapper.TagNamingPolicy{TagKey: "mapper"}
	m := mapper.NewMapper(policy)
	testData := map[string]interface{}{
		"host": "0.0.0.0",
		"port": 7890,
	}

	httpProxy := HttpProxy{}
	err := m.Map(testData, &httpProxy)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", httpProxy)
}
