package main

import (
    "fmt"
    "os"
)

func main() {

    f := createFile("./defer.txt")
    defer closeFile(f)  //Suppose we wanted to create a file, write to it, and then close when we’re done. Here’s how we could do that with defer.
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
	f, err := os.Create(p)
	fmt.Println("error is: ", err)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}