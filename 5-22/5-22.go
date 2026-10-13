// 5-22
package main

import (
	"fmt"
)

func f() {
	defer func() {
		s := recover()

		fmt.Println(s)
	}()
	panic("ERROR !!!")
}

func main() {
	f()
	fmt.Println("Hello World!")
}
