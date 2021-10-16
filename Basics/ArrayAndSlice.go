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

	for

	// 二次元配列の初期化(スライス) n * m
	graph := make([][]string, 3)

	for i:=0; i<3; i++{
		graph[i] = make([]string, 5)
	}
	fmt.Println("graph[][] ", graph)

	/* スライス  */
	var slice1 []int

	// 配列の型、長さ、容量
	slice2 = make([]int, 5, 10)

	var slice3 = []int{1, 2, 3, 4, 5}

	// 5番目が９９、７番目が１３０、その他が０の要素数８個のスライス
	slice4 := []int{5: 99, 7: 130}

	println(slice2[3])

	// 長さ
	println(len(slice2))
	// 容量
	println(cap(slice2))

	slice3 = append(slice3, 6, 7)

	arraysF := []string{}
	arraysF = append(arraysF, "Red")
	arraysF = append(arraysF, "Green")
	arraysF = append(arraysF, "Blue")
}
