package basicgo

import "fmt"

func TypesofFor() {

	// While loop
	i := 1
	for i <= 4 {
		// fmt.Println(i)
		i = i + 1
	}

	// Three-component loop
	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}

	// For-each range loop
	strings := []string{"hello", "hi"}
	for index, s := range strings {
		fmt.Println(index, s)
	}

	// Exit a loop
	sum := 0
	for ix := 1; ix < 5; ix++ {
		if ix%2 != 0 { // skip odd numbers
			fmt.Println(i, "sf")
			continue
		}
		fmt.Println(i, "sim")
		sum = sum + i
	}
	fmt.Println(sum)
}
