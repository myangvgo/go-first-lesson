package main

import (
	"fmt"
	_ "initorder/pkg1"
	_ "initorder/pkg2"
)

var (
	_ = constInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)


const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 has been initialized!")
	}
	
	if c2 != "" {
		fmt.Println("main: const c2 has been initialized!")
	}

	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s has been initialized\n", name)
	return name
}


func init() {
	fmt.Println("main: first init function invoked")
}

func init() {
	fmt.Println("main: second init function invoked")
}


func main() {
	//
}