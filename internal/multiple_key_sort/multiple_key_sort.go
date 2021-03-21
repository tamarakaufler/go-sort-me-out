package multiple_key_sort

import (
	"log"
	"sort"
)

type Student struct {
	Name      string
	Maths     int
	Physics   int
	English   int
	Chemistry int
}

var class = []Student{
	{
		Name:      "Lucien",
		Maths:     32,
		Physics:   30,
		English:   33,
		Chemistry: 29,
	},
	{
		Name:      "Amy",
		Maths:     15,
		Physics:   9,
		English:   16,
		Chemistry: 22,
	},
	{
		Name:      "Marianne",
		Maths:     3,
		Physics:   13,
		English:   6,
		Chemistry: 19,
	},
	{
		Name:      "Mirko",
		Maths:     12,
		Physics:   12,
		English:   6,
		Chemistry: 28,
	},
	{
		Name:      "Marek",
		Maths:     8,
		Physics:   10,
		English:   4,
		Chemistry: 25,
	},
	{
		Name:      "Jane",
		Maths:     11,
		Physics:   16,
		English:   19,
		Chemistry: 26,
	},
}

// used for sorting by different properties.
type By func(ch1, ch2 Student) bool

// classSorter encapsulates the slice to sort and the implementation
// of the Less method for a particular field
type classSorter struct {
	class []Student
	by    func(ch1, ch2 Student) bool
}

// implementation of the sort.Interface
func (cs *classSorter) Len() int {
	return len(cs.class)
}

func (cs *classSorter) Swap(i, j int) {
	cs.class[i], cs.class[j] = cs.class[j], cs.class[i]
}

func (cs *classSorter) Less(i, j int) bool {
	return cs.by(cs.class[i], cs.class[j])
}

// the by receiver is the sorting method for a particular Student property/field.
func (by By) Sort(c []Student) {
	cs := &classSorter{
		class: c,
		by:    by,
	}

	// classSorter implements sort.Interface interface, so can be used as input
	// to the Sort method.
	sort.Sort(cs)
}

// custom Less functions to sort by each property
var maths = func(ch1, ch2 Student) bool {
	return ch1.Maths < ch2.Maths
}
var physics = func(ch1, ch2 Student) bool {
	return ch1.Physics < ch2.Physics
}
var english = func(ch1, ch2 Student) bool {
	return ch1.English < ch2.English
}
var chemistry = func(ch1, ch2 Student) bool {
	return ch1.Chemistry < ch2.Chemistry
}

func SortClassByMathsResults(c []Student) {
	if c == nil {
		c = class
	}

	log.Printf("multiple_key_sort - SortClassByMathsResults before %+v\n", c)

	// maths has the same signature as the By method, so we cast maths to By
	// and call the By.Sort method on c.
	By(maths).Sort(c)

	log.Printf("multiple_key_sort - SortClassByMathsResults after %+v\n\n", c)
}

func SortClassByPhysicsResults(c []Student) {
	if c == nil {
		c = class
	}

	log.Printf("multiple_key_sort - SortClassByPhysicsResults before %+v\n", c)

	// maths has the same signature as the By method, so we cast maths to By
	// and call the By.Sort method on c.
	By(physics).Sort(c)

	log.Printf("multiple_key_sort - SortClassByPhysicsResults after %+v\n\n", c)
}

func SortClassByEnglishResults(c []Student) {
	if c == nil {
		c = class
	}

	log.Printf("multiple_key_sort - SortClassByEnglishResults before %+v\n", c)

	// maths has the same signature as the By method, so we cast maths to By
	// and call the By.Sort method on c.
	By(english).Sort(c)

	log.Printf("multiple_key_sort - SortClassByEnglishResults after %+v\n\n", c)
}

func SortClassByChemistryResults(c []Student) {
	if c == nil {
		c = class
	}

	log.Printf("multiple_key_sort - SortClassByChemistryResults before %+v\n", c)

	// maths has the same signature as the By method, so we cast maths to By
	// and call the By.Sort method on c.
	By(chemistry).Sort(c)

	log.Printf("multiple_key_sort - SortClassByChemistryResults after %+v\n\n", c)
}
