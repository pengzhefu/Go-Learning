package main

import (
    "fmt"
    "sort"
)

type byLength []string  // created a byLength type that is just an alias for the builtin []string type.

func (s byLength) Len() int {
    return len(s)
}

// here to define how to swap 
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]  // just like python
}

// here to define what the order should be
func (s byLength) Less(i, j int) bool {
    return len(s[i]) > len(s[j])
}

func main() {

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)  // sort func is in-place, change on itself.
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)  // and each sort should have a type function
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
	i := sort.StringsAreSorted(strs)
	fmt.Println("Sorted: ", i)
	

	// custom sort function
	// and here is also an example of interface, sort.Sort is a kind of function whose arg is interface
	// you should define all the function for the interface
	fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(byLength(fruits))  // sort.Sort() function should have 3 func in its interface: Len, Swap, Less
    fmt.Println(fruits)
}