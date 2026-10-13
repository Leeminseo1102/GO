// 3-13
package main

import (
	"fmt"
)

func main() {
	s := "world"

	switch s {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("일치하지 않습니다")
	}
}
