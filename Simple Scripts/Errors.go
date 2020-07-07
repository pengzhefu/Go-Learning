package main

import (
    "errors"
    "fmt"
)

func f1(arg int) (int, error) {
    if arg == 42 {

        return -1, errors.New("can't work with 42")

    }

    return arg + 3, nil
}
// customize the own error struct
type argError struct {
    arg  int
    prob string
}


func (e *argError) Error() string {  // the name Error could not be changed, since the op will find it automatically
	// and in order to find it automatically, the defined of input(using pointer or struct itself should be same when its infered to the automatic situation)
	// in this case, the method defines the pointer, so the return type should be pointer when we need to use Error Method automatically
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {  // since the second return type of f2 func is error, 
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"}  // return a val & a pointer of struct, it could call the Error() method of struct automatically
    }
    return arg + 3, nil
}

func main() {

    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

	_, e := f2(42)
	fmt.Println("the e is: ", e)
    if ae, ok := e.(*argError); ok {
		fmt.Println(ae)
        fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println(ok)
	}
	
	fmt.Println("build a struct of argError")
	test := argError{10, "string"}
	res1 := test.Error()
	fmt.Println(res1)
}

