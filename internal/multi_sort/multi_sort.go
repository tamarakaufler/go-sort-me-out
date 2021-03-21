package multi_sort

import (
	"fmt"
	"log"
	"sort"
)

type Child struct {
	Name      string
	Maths     int
	Physics   int
	English   int
	Chemistry int
}

// used for sorting by different properties.
//	ch1 < ch2  => -1
//	ch1 == ch2 => 0
//	ch1 > ch2  => 1
type lessFunc func(ch1, ch2 Child) int

type sorter struct {
	class []Child
	less  []lessFunc
}

// sorter has the static type sort.Interface.
var _ sort.Interface = &sorter{}

func (s *sorter) Len() int {
	return len(s.class)
}

func (s *sorter) Swap(i, j int) {
	s.class[i], s.class[j] = s.class[j], s.class[i]
}

func (s *sorter) Less(i, j int) bool {
	ch1, ch2 := s.class[i], s.class[j]

	for k := 0; k < len(s.less); k++ {
		less := s.less[k]

		switch less(ch1, ch2) {
		case -1:
			return true
		case 0:
			return false
		case 1:
			return false
		}
	}
	return false
}

// OrderedBy sets up the sorter.less property as a list of sorting functions.
func OrderedBy(ls ...lessFunc) *sorter {
	return &sorter{
		less: ls,
	}
}

func (s *sorter) Sort(c []Child) {
	s.class = c
	sort.Sort(s)
}

var maths = func(ch1, ch2 Child) int {
	switch {
	case ch1.Maths < ch2.Maths:
		return -1
	case ch1.Maths > ch2.Maths:
		return 1
	}
	return 0
}

var physics = func(ch1, ch2 Child) int {
	switch {
	case ch1.Physics < ch2.Physics:
		return -1
	case ch1.Physics > ch2.Physics:
		return 1
	}
	return 0
}

var english = func(ch1, ch2 Child) int {
	switch {
	case ch1.English < ch2.English:
		return -1
	case ch1.English > ch2.English:
		return 1
	}
	return 0
}

var chemistry = func(ch1, ch2 Child) int {
	switch {
	case ch1.Chemistry < ch2.Chemistry:
		return -1
	case ch1.Chemistry > ch2.Chemistry:
		return 1
	}
	return 0
}

var class = []Child{
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
		Maths:     12,
		Physics:   13,
		English:   16,
		Chemistry: 19,
	},
	{
		Name:      "Mirko",
		Maths:     12,
		Physics:   13,
		English:   16,
		Chemistry: 24,
	},
	{
		Name:      "Marek",
		Maths:     12,
		Physics:   13,
		English:   16,
		Chemistry: 25,
	},
	{
		Name:      "Jane",
		Maths:     12,
		Physics:   16,
		English:   19,
		Chemistry: 26,
	},
}

func SortByAllSubjects(c []Child) {
	if c == nil {
		c = class
	}

	//log.Printf("multi_sort - SortByAllSubjects before %+v\n", c)

	//OrderedBy(maths, english).Sort(c)
	OrderedBy(maths, physics, english, chemistry).Sort(c)

	log.Printf("multi_sort - SortByAllSubjects after\n\n")
	for _, ch := range c {
		fmt.Printf("%s\n", ch.Name)
		fmt.Printf("\tmaths %v\n", ch.Maths)
		fmt.Printf("\tenglish %v\n", ch.English)
		fmt.Printf("\tphysics %v\n", ch.Physics)
		fmt.Printf("\tchemistry %v\n", ch.Chemistry)
	}
}
