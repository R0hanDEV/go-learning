package basicgo

import "fmt"

// argument function that take value and return value if you want
func ArgumentFunction(s string) string {

	// Anomynous function

	// Anomynous function always exceute inside the local function not in global function.
	func() {
		fmt.Println("Welcome! to GeeksforGeeks")
	}()
	return s
}

func CallByvalue(p, q int) {
	var swp int
	swp = p
	p = q
	q = swp
	fmt.Println(p, q, "after passing argument")
}
