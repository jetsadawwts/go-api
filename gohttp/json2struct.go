package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {

	//Req.Body
	data := []byte(`{
		"id": 2,
		"name" : "Jetsadawwts",
		"age" : 19
		}`)

	
	// u := &User{}
	// json.Unmarshal(data, u)
	var u User
	err := json.Unmarshal(data, &u)
	fmt.Printf("% #v\n", u)
	fmt.Println(err)

}