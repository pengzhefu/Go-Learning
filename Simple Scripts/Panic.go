package main

import (
  "os"
  "fmt"
)

func main() {

    panic("a problem")  // custom own error info, so here means that this program will return an error info when execution done

    _, err := os.Create("./file.txt")
    fmt.Println("error is: ", err)
    if err != nil {
		panic(err)  // if some error happen here, the output will like the above line: panic: a problem
					// or, if not mentioned in the above panic, it will display the system defines
    }
}