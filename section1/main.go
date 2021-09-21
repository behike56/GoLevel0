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
