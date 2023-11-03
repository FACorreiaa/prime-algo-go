package day_maps

import "fmt"

func DisplayMap() {
	m := make(map[string]string)
	m["first"] = "first"
	m["second"] = "second"
	m["third"] = "third"

	v1 := m["first"]
	fmt.Println("first:", v1)

	delete(m, "third")
	fmt.Println(m)
	for k, v := range m {
		fmt.Println("key: %s, value $s:", k, v)

	}
	clear(m)
	fmt.Println(m)
}
