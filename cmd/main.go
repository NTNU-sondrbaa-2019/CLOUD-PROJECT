package main

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
)

func main() {

	type Test struct {
		Name string `json:"name"`
		Author string `json:"author"`
	}

	test := Test {
		"This is a test JSON",
		"Sondre Benjamin Aasen",
	}

	CO1Cache.Initialize()
	CO1Cache.WriteJSON("test", test)

	fmt.Println("Hello World!")

}
