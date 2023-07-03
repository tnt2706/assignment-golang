package controllers

import "fmt"

func init() {
	fmt.Println("Initializing user controller")
}

func Area(len, wid float64) float64 {
	area := len * wid
	return area
}
