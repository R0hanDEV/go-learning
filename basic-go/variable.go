package basicgo

import "fmt"

func Variable() {

	// declare and assign value after that
	var a int
	a = 10
	fmt.Println(a)

	// declare and assign value at same time
	var b int = 14
	fmt.Println(b)

	// const variable declaration that cannot be changed
	const c int = 10
	fmt.Println(c)

	// Short variable declarations
	d, f := "app", 10
	fmt.Println(d, f)
}
