package main

import "fmt"

func main() {
	var x int = 1

	for i:=1; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("loop")
	}

	for {
		x++
		if x > 10 {
			break
		}
	}

	fmt.Println(x)


}