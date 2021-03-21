# go-sort-me-out
Go examples of different sorting strategies

The examples revolve around a class of students. If that bothers you, please, avoid, so as not to suffer a psychological damage.

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

where c is the underlying type of ByAge, ie []Student. The sort.Sort methods needs an input parameter of sort.Interface type, so we need to do type conversion of []Student into ByAge:

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

## Map Sort

Sorting of a map by its values uses the same approach as the Simple Sort above. A new type NameAge with an underlying type of a list is introduced, which satisfies the sort.Interface interface. The map is then used to create a NameAge variable.

## Multiple Key Sort

We have a struct that contains several fields that would want to individually sort by. As an example, there is a Student with Name and test points in several subjects. We want to be able to sort a class of students by their performace in different subjects.

### Approach 1

It is possible to introduce one sorting type per each subject:

    type ByMaths   []Student
    type ByPhysics []Student
### Approach 2

Another approach is to introduce one generic function type and a sorter struct and implement the sorter.Sort method rather than implementing ByMaths, ByPhysics etc as a sort.Interface type.
The function type is of signature:

    type By func(ch1, ch2 Student) bool

- By function implements Sort method.

- sorter struct, in our case classSorter, encapsulates the list to be sorted and the particular sorting function of type By. The sorter struct implements the sort.Interface methods. See how the classSorter.Less method uses the classSorter.by function.

To sort by Maths grades, we create maths function of the By function signature. To be able to take advantage of the implemented sorter.Sort method, we do type conversion and run the Sort method:

    var maths = func(ch1, ch2 Student) bool {
	    return ch1.Maths < ch2.Maths
    }  

    By(maths).Sort(c)

An advantage of Approach 2 is better readability:

    By(maths).Sort(c)
    By(english).Sort(c) ...

as opposed to

    sort.Sort(ByMaths(c))
    sort.Sort(ByEnglish(c)) ...

## Multi sort

This is continued sorting by provided criteria when sorting by one field encounters field equivalence.

The implementation somewhat resembles the Multiple Key Sort approach.

A new sorting function type is introduced

    type lessFunc func(ch1, ch2 Student) int

While similar to the one in Multiple Key Sort section, the return value here is an integer rather than a boolean. The return value can be:
- -1 when a[i] is smaller than a[j]
-  0 if the are equal
- +1 when a[i] is bigger than a[j]

Again, there is a sorter struct (classSorter) keeping the list to be sorted and this time, rather than one sorting function, there is a list of sorting functions.

The sorter struct

    type sorter struct {
        class []Student
        less  []lessFunc
    }

satisfies the sort.Interface interface. The Less method is a bit more complicated than before. It loops through sorter.less list. The reason why the lessFunc (ie the sorter.less functions) returns-1/0/+1 is to simplify and eliminate additional looping.

OrderedBy function sets up the sorter.less field and the sorter.Sort method sets up the sorter.class field and runs sort.Sort(sorter_instance)
