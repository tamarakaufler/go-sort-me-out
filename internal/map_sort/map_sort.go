package map_sort

import (
	"fmt"
	"log"
	"sort"
)

type nameAge []student

var _ sort.Interface = (nameAge)(nil)

type student struct {
	name string
	age  float32
}

func (na nameAge) Len() int {
	return len(na)
}

func (na nameAge) Swap(i, j int) {
	na[i], na[j] = na[j], na[i]
}

func (na nameAge) Less(i, j int) bool {
	return na[i].age < na[j].age
}

var class = map[string]float32{
	"Lucien": 23,
	"Marek":  18,
	"Mirko":  22.5,
	"Radka":  32,
	"Honza":  30,
}

func SortClassMap(c map[string]float32) {
	if c == nil {
		c = class
	}

	cs := nameAge{}
	for k, v := range c {
		cs = append(cs, student{name: k, age: v})
	}

	sort.Sort(cs)

	log.Printf("map sort - SortClassMap after sorting\n")
	for _, s := range cs {
		fmt.Printf("%s\n", s.name)
		fmt.Printf("\tage %.2f\n", s.age)
	}
}
