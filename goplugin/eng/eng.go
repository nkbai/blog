package main

import "fmt"
import "github.com/nkbai/blog/goplugin/anotherlib"

type greeting string

func (g greeting) Greet() {
	fmt.Println("Hello Universe")
}
func (g greeting) GetShareVariable() int{
	return anotherlib.ShareVariable
}
// exported
var Greeter greeting