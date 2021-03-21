package simple_sort

import (
	"log"
	"sort"
)

type Child struct {
	Name string
	Age  float32
}

type ByAge []Child

// satisfy sort.Interface interface.
func (b ByAge) Len() int {
	return len(b)
}
func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

var class = []Child{
	{
		Name: "Lucien",
		Age:  17,
	},
	{
		Name: "Amy",
		Age:  16.5,
	},
	{
		Name: "Marianne",
		Age:  19,
	},
	{
		Name: "Mirko",
		Age:  16.5,
	},
	{
		Name: "Marek",
		Age:  18.5,
	},
	{
		Name: "Jane",
		Age:  19,
	},
}

func SortClassByAge(c []Child) {
	if c == nil {
		c = class
	}

	log.Printf("simple sort - SortClassByAge: before sorting %+v\n", c)

	// we have a class c, which has the same underlying type as ByAge.
	// ByAge satisfies the sort.Interface interface (Less/Swap/Less methods) and
	// can be therefore provided to the Sort method.
	// The c variable can be cast into ByAge, therefore sorted by age.
	sort.Sort(ByAge(c))

	log.Printf("simple sort - after sorting SortClassByAge: %+v\n\n", c)
}

func SliceSortClassByAge(c []Child) {
	if c == nil {
		c = class
	}

	log.Printf("simple sort - SliceSortClassByAge: before sorting %+v\n", c)

	// we can use sort.Slice, that accepts two parameters:
	//		a slice to be sorted
	//		a closure with the signature of the Less method of the sort.Interface
	sort.Slice(c, func(i, j int) bool {
		return c[i].Age < c[j].Age
	})
	log.Printf("simple sort - after sorting SliceSortClassByAge: %+v\n\n", c)
}
