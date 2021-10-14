package main

import (
	"fmt"
)

func main () {
	/* 配列　*/
	// 宣言と初期化
	var arraysA [5]string
	arraysA[0] = "Apple"
	arraysA[1] = "Banana"
	arraysA[2] = "Cactus"
	arraysA[3] = "Deets"
	arraysA[4] = "Eagle"

	fmt.Println("arraysA: ", arraysA)
	fmt.Println("length: ", len(arraysA), "capacity: ", cap(arraysA))

	var arraysB [3]string = [3]string {"Apple", "Banana", "Cactus"}
	fmt.Println("arraysB: ", arraysB)

	arraysC := [...] string{"Apple", "Banana", "Cactus"}
	fmt.Println("arraysC: ", arraysC)

	arraysD := [3]string{}
	fmt.Println("arraysD: ", arraysD)

	arraysE := [...]int{3:15, 7:21}
	fmt.Println("arraysE", arraysE)

	// 配列の繰り返し処理
	for idc, val := range arraysA {
		fmt.Println("index is ", idc)
		fmt.Println("value is ", val)
	}

	// 二次元配列の初期化(スライス) n * m
	graph := make([][]string, 3)

	for i:=0; i<3; i++{
		graph[i] = make([]string, 5)
	}
	fmt.Println("graph[][] ", graph)

	/* スライス  */
	arraysF := []string{}
	arraysF = append(arraysF, "Red")
	arraysF = append(arraysF, "Green")
	arraysF = append(arraysF, "Blue")
}
