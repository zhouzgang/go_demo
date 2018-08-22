package main

func main() {
	//for true {
	//	println("fafa")
	//}
	//for _, c := range `_\|/`{
	//	fmt.Printf("\r%c", c)
	//}
	//for i, c := range "go" {
	//	fmt.Println(i, c)
	//}
	//nums := []int{1,2,3,4}
	//for i := 0; i<len(nums); i++ {
	//	fmt.Println(nums[i])
	//}
	slice1 := make([]int, 10)
	println(cap(slice1))
	slice1[0] = 7
	println(slice1[0])
	for k,v := range slice1 {
		println(k, v)
	}

}
