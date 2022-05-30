package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

}
