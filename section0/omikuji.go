package main

import (
	"fmt"
	"math/rand"
	"time"

)

func main(){
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(5)

	switch s {
	case 0:
		fmt.Println(s+1, ":凶")
	case 1, 2:
		fmt.Println(s+1, ":吉")
	case 3, 4:
		fmt.Println(s+1, ":中吉")
	case 5:
		fmt.Println(s+1, ":大吉")
	}
}
