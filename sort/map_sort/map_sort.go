package map_sort

import (
	"fmt"
	"log"
	"sort"
)

type NameAge []Student

var _ sort.Interface = (NameAge)(nil)

type Student struct {
	name string
	age  float32
}

func (na NameAge) Len() int {
	return len(na)
}

func (na NameAge) Swap(i, j int) {
	na[i], na[j] = na[j], na[i]
}

func (na NameAge) Less(i, j int) bool {
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

	cs := NameAge{}
	for k, v := range c {
		cs = append(cs, Student{name: k, age: v})
	}

	sort.Sort(cs)

	log.Printf("map sort - SortClassMap after sorting\n")
	for _, s := range cs {
		fmt.Printf("%s\n", s.name)
		fmt.Printf("\tage %.2f\n", s.age)
	}
}
