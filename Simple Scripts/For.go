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

	channel := make(chan string, 2)

	// channel <- "one"
	// channel <- "two"

	// close(channel)

	// two ways to implement for over channel
	// for {
	// 	elem, more := <-channel  // using this way need to assure that the channel is closed when stop inputting the data into channel
	// 	if more {
	// 		fmt.Println("here is the first way: ", elem, more)
	// 	} else {
	// 		fmt.Println("first way done")
	// 		break
	// 	}
	// }

	channel <- "one"
	channel <- "two"
	close(channel) // using this way need to assure that the channel is closed when stop inputting the data into channel
	for elem := range channel {  // it would not have index or boolean, just the element, unlike the for range loop for slice
		fmt.Println("the second way: ",elem)
	}

	// both two for range ways need to close the channel if channel stops receiving the data
	// and, both ways just use normal funcs, not goroutine


}