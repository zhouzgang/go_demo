package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var v1 uint = 2
	var v2 int8 = -128
	var v3 float32 = 1.2
	var v4 = true
	v5 := 1
	v6 := "abcde"
	//var v7, v8, v9 = 1, "a", true
	v7, v8, v9 := 1, "a", true	// 这种格式只能在函数体中出现
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

}

