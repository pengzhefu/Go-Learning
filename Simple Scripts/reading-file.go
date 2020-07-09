package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
	"os"
	"encoding/json"
)

type log struct {
	Devid   string `json:"devid"`
	TldStep int    `json:"tld_step"`
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// reading normal file

	// the following method will put the whole file into memory one time, not good
	_, err1 := ioutil.ReadFile("normal.txt")
	checkError(err1)
	// fmt.Println(string(wholeFile))

	// reading the file
	// f, err := os.Open("normal.txt")  // f is the pointer of os.File, like 
	// checkError(err)

	// defer f.Close()

	// b1 := make([]byte, 5)  // using []byte to store string
	// n1, err := f.Read(b1)  // will read the first line 
    // checkError(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	
	// n2, err := f.Read(b1)  // will read the second line 
    // checkError(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b1[:n2]))
	
	// using the line by line reader
	lineFile, err2 := os.Open("sample")
	checkError(err2)

	defer lineFile.Close()

	br := bufio.NewReader(lineFile)  // just like a generator of file in python

	for {
		a,_,c := br.ReadLine()  // a is content, c is error, 
		if c == io.EOF {
			break
		}
		str1 := string(a)
		fmt.Println(str1)

		// try to decode json array of file
		tmp := log{}
		json.Unmarshal([]byte(str1), &tmp)
		fmt.Println(tmp)
		
	}


	

}