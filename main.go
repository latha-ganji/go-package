package gopackage

import (
	"fmt"
	//a "github.com/latha-ganji/go-package/testPackage"
)

func main() {
	fmt.Println("hello")
	fmt.Println(Add(10, 20))

	//var message = a.TestPackageLib("Hello Latha")
	//fmt.Println(message)
}

func Add(a, b int) int {
	return a + b
}
