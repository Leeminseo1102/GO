// 6-7
package main

import (
	"fmt"
)

type Rectangle struct {
	width  int
	height int
}

func RectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func main() {
	rect := Rectangle{30, 30}

	area := RectangleArea(&rect)

	fmt.Println(area)
}
