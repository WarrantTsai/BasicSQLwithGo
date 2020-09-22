package main

import (
	"fmt"
)

type MyInt int

func (mi MyInt) String() string {
    return fmt.Sprint("*", int(mi), "*")
}

func main() {
    var mi MyInt = 2
    fmt.Printf("%d %v\n", mi, mi)
	
	var v int = 4
	fmt.Printf("%d", v)
}