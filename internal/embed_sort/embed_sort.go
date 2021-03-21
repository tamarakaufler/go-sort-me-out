package embed_sort

import (
	"fmt"
	"log"
	"sort"
)

type Student struct {
	FirstName string
	LastName  string
	Age       float32
}

type Class []Student

// Class satisfies 2 of the sort.Interface methods: Len and Swap.
func (c Class) Len() int {
	return len(c)
}

func (c Class) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// ByFirstName embeds Class, therefore gets its methods. ByFirstName
// satisfies the third sort.Interface method, Less, therefore is a
// variable of the sort.Interface type.
type ByFirstName struct {
	Class
}

// ByLastName embeds Class, therefore gets its methods. ByLastName
// satisfies the third sort.Interface method, Less, therefore is a
// variable of the sort.Interface type.
type ByLastName struct {
	Class
}

// ByAge embeds Class, therefore gets its methods. ByAge
// satisfies the third sort.Interface method, Less, therefore is a
// variable of the sort.Interface type.
type ByAge struct {
	Class
}

func (f ByFirstName) Less(i, j int) bool {
	return f.Class[i].FirstName < f.Class[j].FirstName
}

func (f ByLastName) Less(i, j int) bool {
	return f.Class[i].LastName < f.Class[j].LastName
}

func (f ByAge) Less(i, j int) bool {
	return f.Class[i].Age < f.Class[j].Age
}

var class = []Student{
	{
		FirstName: "Lucien",
		LastName:  "Kaufler",
		Age:       15,
	},
	{
		FirstName: "Amy",
		LastName:  "Brunel",
		Age:       16.5,
	},
	{
		FirstName: "Marianne",
		LastName:  "Maxwell",
		Age:       19,
	},
	{
		FirstName: "Mirko",
		LastName:  "Baumann",
		Age:       16.5,
	},
	{
		FirstName: "Marek",
		LastName:  "McGuyere",
		Age:       18.5,
	},
	{
		FirstName: "Jane",
		LastName:  "Nicholson",
		Age:       19,
	},
}

func SortByLastName(c []Student) {
	if c == nil {
		c = class
	}

	sort.Sort(ByLastName{c})
	log.Printf("\n\nembed_sort - SortByLastName after\n")
	for _, s := range c {
		fmt.Printf("Last Name %v\n", s.LastName)
		fmt.Printf("\tFirst Name %v\n", s.FirstName)
		fmt.Printf("\tAge %v\n", s.Age)
	}
}

func SortByFirstName(c []Student) {
	if c == nil {
		c = class
	}

	sort.Sort(ByFirstName{c})
	log.Printf("\n\nembed_sort - SortByFirstName after\n")
	for _, s := range c {
		fmt.Printf("First Name %v\n", s.FirstName)
		fmt.Printf("\tLast Name %v\n", s.LastName)
		fmt.Printf("\tAge %v\n", s.Age)
	}
}

func SortByAge(c []Student) {
	if c == nil {
		c = class
	}

	sort.Sort(ByAge{c})
	log.Printf("\n\nembed_sort - SortByAge after\n")
	for _, s := range c {
		fmt.Printf("Age %v\n", s.Age)
		fmt.Printf("\tFirst Name %v\n", s.FirstName)
		fmt.Printf("\tLast Name %v\n", s.LastName)
	}
}
