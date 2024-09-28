package main

import (
	"fmt"
	basicgo "go-learning/basic-go"
)

func main() {
	// basicgo.Variable()
	// basicgo.TypesofFor()
	// basicgo.IfElsefunction()
	// basicgo.SwitchFun()
	// fmt.Println(basicgo.ArgumentFunction("fjs"))

	var p int = 10
	var q int = 11
	fmt.Println(p, q, "before passing argument")
	basicgo.CallByvalue(p, q)

}
