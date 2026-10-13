// 5-16
package main

import (
	"fmt"
)

func hello() {
	fmt.Println("Hello World!")
}

func world() {

	fmt.Println("world")
}

func main() {
	defer world()
	hello()
	hello()
	hello()

}
