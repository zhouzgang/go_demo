package main

import (
	"fmt"
	//"github.com/golang/lint/testdata"
)

func init() {
	println("----init-----")
}

func main() {
	fmt.Println("hello world")

	var v1 uint = 2
	var v2 int8 = -128
	var v3 float32 = 1.2
	var v4 = true
	v5 := 1
	v6 := "abcde"
	//var v7, v8, v9 = 1, "a", true
	v7, v8, v9 := 1, "a", true // 这种格式只能在函数体中出现
	var (
		v10 int
		v11 string
	)
	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11)
	fmt.Println(&v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, &v9, &v10, &v11)
	var v12 string = "a"
	var v13 string
	v13 = v12
	fmt.Println(v12, v13, &v12, &v13)

	v14, v15 := 1, 2
	fmt.Println(maxMin(v14, v15, 5, "hahaha"))

	v16, v17 := 2, 4
	swp(&v16, &v17)
	fmt.Println(v16, v17)

	nextNum := getSquence()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(getSquence()())
}

func maxMin(t1, t2, t3 int, t4 string) (int, int) {
	fmt.Println(t4)
	return t3, t1
}

func swp(x, y *int) {
	fmt.Println(x, y)
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func getSquence() func() int {
	i  := 0
	return func() int{
		i += 1
		return i
	}
}
