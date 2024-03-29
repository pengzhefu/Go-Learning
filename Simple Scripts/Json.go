package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type response1 struct {
    Page   int
    Fruits []string
}

//Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.
type response2 struct {
    Page   int      `json:"page"`  // the string means the key name in json array
    Fruits []string `json:"fruits"`
}

type log struct {
	Devid   string `json:"devid"`
	TldStep int    `json:"tld_step"`
}

func main() {


	// json.Marshal is for encoded, convert data structure 2 json array
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

	// so the input data should be []byte?
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	//We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{} will hold a map of strings to arbitrary data types.
    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

	// the way convert data type in golang
    num := dat["num"].(float64)
    fmt.Println(num)

    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    str := `{"page": 1, "fruits": ["apple", "peach"]}`  // string in byte type
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

	// easy way convert map to json array
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
	
	// single json array, that only has one level
	fmt.Println("test by zfpeng")
	singleStr := `{"key1": "value1", "key2": 1234}`
	var singleMap map[string]interface{}
	if err := json.Unmarshal([]byte(singleStr), &singleMap); err != nil {
        panic(err)
    }
	fmt.Println(singleMap)
	fmt.Println(singleMap["key1"].(string))  // this converting method only used for []byte, `123` this one
    fmt.Println(singleMap["key2"].(float64))  // it should be float64 when using interface
    
    fmt.Println("test for custom json")
    tmp := `{"devid":"7734990C","tld_step":2}`
    ret := log{}
    json.Unmarshal([]byte(tmp), &ret)
    fmt.Println(ret)
    fmt.Println(ret.Devid)
}