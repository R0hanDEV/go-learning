package basicgo

import "fmt"

func IfElsefunction() {
	i := 1
	if i == 0 {
		fmt.Print(i)
	} else {
		fmt.Println("not zero")
	}
}

func SwitchFun() {
	i := 1
	switch i {
	case 0:
		fmt.Println("zero")
	case 2:
		fmt.Println("once")
	default:
		fmt.Println("don't know")
	}
}
