package main

import (
	"fmt"
)

func main(){
	/* 変数  */
	var n int = 100

	var m int
	m = 200

	// 型推論
	var l = 300

	// var 省略
	num := 350

	// まとめる
	var (
		a = 100
		b= 200
	)

	// 変数の０値
	var initNum int // 0
	var initStr string // ""
	var initBool bool // false
	var initError error // nil

	var hello string = "Hello, World."

	println(hello)
}

func secound(){
	/* 定数  */
	// 数値りてらる
	num := 100
	num2 := 1.5
	num3 := 1+4i

	// 文字列りてらる
	str := "hoge"

	// ルーンリテラル
	charStr := 'A'

	// 真偽値リテラル
	boolTrue := true
	boolFalse := false

	// 表記
	binNum := 0b_1100_0100
	octNum := 0o_144
	hexNum := 0x1.5p+03

	// 名前付き定数
	const n int = 100
	const m = 100
	const s = "hello, " + "world"
	const (
		x = 100
		y = 200
	)
	// 定数のデフォルトの型
	const z = 100
	// 定数は指定しない限り型を持たない。
	// 変数に代入するとデフォルトの型を持つようになる
	var n = z
	// 定数は型を持たないので無限の精度を持つ
	const infinit = 100000000 / 10000000

	println(infinit)

	// 右辺の省略ができる
	const (
		o = 1 + 2
		p
		q
	)
	// 連続した定数を作る
	const (
		x1 = iota
		x2
	)
	const (
		y1 = 1 << iota
		y2
		y3
	)
	println(x1, x2, y1, y2, y3)
}

func third(){
	/* 演算子 */
	n := 100 + 300
	m := n + 200
	msg := "top" + "of"
}

func forth(){
	a := 3

	// 改行があると自動でセミコロンを付与する
	// 改行のあとに｛があるとコンパイルエラー
	if a == 0 {
	}
	// a := f(); a > 0
	// 代入分を書くこともできる

	switch a {
	case 1, 2:
		fmt.Println("a is 1 or 2")
	default:
		fmt.Println("default")
	}

	switch {
	case b == 1:
		fmt.Println("b is 1")
	}

	// 繰り返し
	for i := 0; i <= 100; i = i + 1 {

	}
	// while文の代わりになる
	for i <= 100 {

	}
	// rangeを使った繰り返し
	for i, v := range []int{1, 2, 3}{

	}

	// breakによるループの抜け出し
	var i int
	for {
		if i % 2 == 1 {
			break
		}
		i++
	}

	// ラベル指定のbreak(goto)
	var j int
	LOOP:
	for {
		switch {
		case j % 2 == 1:
			break LOOP
		}
		j++
	}
}
