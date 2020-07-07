package main

import "fmt"

func main() {

	var sum int = 0

	list := []int {1,2,3,4,5}

	for i, elem := range list{  // here using := to replace in python, and there are two val for the range: element and the index
		if elem == 3 {
			fmt.Println("here to print index of 3: ", i)
		}

		sum += elem

	}

	fmt.Println("the sum is: ", sum)


	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {  // iterate the key, value; if there is only k, it would be iterate for the keys
        fmt.Printf("%s -> %s\n", k, v)
	}
	
	for i, c := range "go" {  // c would be the unicode of char
        fmt.Println(i, c)
    }

}