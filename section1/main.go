package main

import (
	"fmt"
)

func main() {
	chap1()
	chap1Try()
	chap1Try2()

}

func chap1(){
	/* 型のキャスト  */
	var f float64 = 10
	var n int = int(f)
	println(n)
}

func chap1Try(){
	var sum int
	sum = 5 + 6 + 3
	avg := sum / 3
	if avg > 4.5 {
		println("good")
	}
}

func chap1Try2(){
	var a, b, c bool
	if a && b || !c {
		println("True")
	} else {
		println("False")
	}
}

func chap2(){
	/* コンポジット型 */
	// 複数のデータ型が集まって一つのデータ型になったもの

	// 構造体
	// 異なるデータ型を集めたデータ型
	// ゼロ値はフィールドすべてがゼロ値
	var p = struct {
		// 型リテラル
		name string
		age int
	}{
		name: "Gopher",
		age: 10,
	}
	p.age++
	println(p.name, p.age)

}

func chap03() {
	/* 配列 */
	var ns1 [5]int
	var ns2 = [5]int{10, 20, 30, 40, 50}
	ns3 := [...]int{10, 20, 30}
	ns4 := [...]int{5: 50, 10: 100}

	ns := [...]int{1, 2, 3, 4, 5}
	println(ns[3])
	println(len(ns))
	fmt.Println(ns[1:2])
}
