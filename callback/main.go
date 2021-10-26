package main

import "fmt"

func callback(addAsArgument func(x, y int64) int64, name string) {
	fmt.Println("Before")
	fmt.Println(name)
	fmt.Println(addAsArgument(5, 6))

}
func main() {
	myAdd := func(x, y int64) int64 {
		fmt.Println("Inside")
		return x + y
	}

	callback(myAdd, "Fer")
}
