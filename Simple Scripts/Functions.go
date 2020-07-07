package main

import "fmt"

func sum(a int, b int) int {  // in golang, the type should after the var; and beofre the brace, sould declare the type of return result
	return a+b
}


func sumTwo(a,b int) int {  // if the type of arguments is the same, could just declare once after all same-type args
	return a+b
}

func divideTwo(a, b int) (int, int) {  // using (type, type) to infer the concerned mutiple returns, but not tuples

	res1 := a / b
	res2 := a % b
	return res1, res2

}

func varyInput(nums ...int) int {

	total := 0
	// the reason for there should be a _ & num: single num could be the index
	for _, num := range nums {  // that the various number of args would be like the slice
		total += num
		// fmt.Println("the num is: ", num)
	}

	return total

}


func intSeq() func() int {  // func1 result: func2 result
    i := 0
    return func() int {  // anymouns function, without naming the function
        i++
        return i
    }
}


func main(){

	// res := sum(1,2)
	// fmt.Println(res)

	// res2 := sumTwo(2,3)
	// fmt.Println(res2)

	// ret1, ret2 := divideTwo(3,2)
	// fmt.Println(ret1)
	// fmt.Println(ret2)

	test := varyInput(1,2,3,4,5,6)
	fmt.Println(test)

	slice := []int{1, 2, 3, 4}
	test2 := varyInput(slice...)  // if input args already is a slice, should add...
	fmt.Println(test2)

	
	nextInt := intSeq()  // nextInt is the result of intSeq, which is a func. so it should be we have a function named nextInt

	fmt.Println(nextInt())
    fmt.Println(nextInt())
	fmt.Println(nextInt())  // and the state (i) of the function will updated each time, but only for the nextInt function
	
	newInts := intSeq()
    fmt.Println(newInts())
}