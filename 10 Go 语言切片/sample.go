package main

import "fmt"

func main() {
	sample1()
	fmt.Println("=====================")
	sample2()
	fmt.Println("=====================")
	sample3()
	fmt.Println("=====================")
	sample4()
	fmt.Println("=====================")
}

func sample1() {
	var numbers = make([]int, 3, 5)

	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func sample2() {
	var numbers []int

	printSlice2(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

func printSlice2(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func sample3() {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice3(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice3(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice3(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice3(number3)

}

func printSlice3(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func sample4() {
	var numbers []int
	printSlice4(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice4(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice4(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice4(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*20)
	// numbers1 := make([]int, len(numbers)*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice4(numbers1)
}

func printSlice4(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
