package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "-偶数")
		}

		if i % 2 != 0 {
			fmt.Println(i, "-奇数")
		}
	}

	for i := 1; i <= 100; i++ {
		switch i % 2 == 0{
			case true:
			fmt.Println(i, "-偶数")
			case false:
			fmt.Println(i, "-奇数")
		}
	}
}
