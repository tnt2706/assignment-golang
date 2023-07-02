package main

import (
	"assignment/src/controllers"
	"fmt"
)

func main() {
	area := controllers.Area(10.8, 5.6)
	fmt.Printf("Area of rectangle is %f", area)
}
