package main

import (
	"github.com/tamarakaufler/go-sort-me-out/sort/embed_sort"
	"github.com/tamarakaufler/go-sort-me-out/sort/map_sort"
	"github.com/tamarakaufler/go-sort-me-out/sort/multi_sort"
	"github.com/tamarakaufler/go-sort-me-out/sort/multiple_key_sort"
	"github.com/tamarakaufler/go-sort-me-out/sort/simple_sort"
)

func main() {
	simple_sort.SortClassByAge(nil)
	simple_sort.SliceSortClassByAge(nil)

	map_sort.SortClassMap(nil)

	multiple_key_sort.SortClassByMathsResults(nil)
	multiple_key_sort.SortClassByPhysicsResults(nil)
	multiple_key_sort.SortClassByEnglishResults(nil)
	multiple_key_sort.SortClassByChemistryResults(nil)

	multi_sort.SortByAllSubjects(nil)

	embed_sort.SortByFirstName(nil)
	embed_sort.SortByLastName(nil)
	embed_sort.SortByAge(nil)
}
