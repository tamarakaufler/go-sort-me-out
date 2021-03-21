# go-sort-me-out
Go examples of different sorting strategies

# Sorting of concrete types - structs

The necessity to sort a list of structs by a particular struct field requires custom implementation of the sorting mechanism. Go provides sort.Sort method accepting a parameter of _sort.Interface_ type, therefore the concrete type to be sorted must be turned into the interface type.

Interface variables of _sort.Interface_ type must implement 3 methods:
- Len() int
- Swap(i,j int)
- Less(i,j int) bool

Example implementations of different approaches are provided in the internal directory.


## Simple Sort

Uses either sort.Sort or sort.Slice methods.

### sort.Sort

A custom type ByAge is introduced for the list of structs:

    type Student struct {
        Name string
        Age  float32
    }

    type ByAge []Student

This type needs to satisfy the sort.Interface. After implementing the three required methods, the sorting is done through:

    sort.Sort(ByAge(c))

where c is the underlying type of ByAge, ie []Student. The sort.Sort methods needs an input parameter of sort.Interface type, so we do type conversion of []Student into ByAge:

    ByAge(c)

### sort.Slice

The sort package provides sort.Slice method, with input of a list and a closure with the signature of Less method. In our case the usage would be:

a) sorting by age

	sort.Slice(c, func(i, j int) bool {
		return c[i].Age < c[j].Age
	}

a) sorting by name

	sort.Slice(c, func(i, j int) bool {
		return c[i].Name < c[j].Name
    }
