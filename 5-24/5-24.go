// 5-24
package main

import (
	"fmt"
)

func f() {
	defer func(){
		s := recover()
		fmt.Println(s)
	}
	
	a := [...]int{1,2}
	
	for i := 0; i < ; i++{
		fmt.Println()
	}
}
func main() {
	f()
	
	fmt.Println("Hello World!")
}
