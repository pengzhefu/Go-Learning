package main

import (
	"fmt"
	// "time"
	)

func main() {

	var str1 string = "string"

	switch str1 {  // if there is a var, so the type for the case should be same.
	case "123":
		fmt.Println("the input is 123")
	
	case "string":
		fmt.Println("match")

	// case string:
	// 	fmt.Println("match string")

	default:  // default case
		fmt.Println("default")

	}

	// switch-case could be the function, so there could be a return value for swich-case
	whatAmI := func(i interface{}) {   // function name := func (parameter interface {}) {function body}
        switch t := i.(type) {  // this type could be only used in switch-case
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
	whatAmI("hey")
	



}