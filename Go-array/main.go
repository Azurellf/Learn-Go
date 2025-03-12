package main

import (
	"fmt"
)

func main() {
	//隐式初始化
	var array0 [6]int
	fmt.Printf("%T\n", array0)
	//显式初始化
	var array1 = [3]int{1, 2, 3}
	fmt.Println(array1)
	//不带长度的显式初始化
	var array2 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", array2)
	//稀疏数组
	var array3 = [...]int{
		99: 39,
	}
	fmt.Printf("%d\n", array3[99])

	var array4 = [10]int{
		4: 23,
	}
	fmt.Printf("%T\n", array4)

	////越界
	//var array5 = [...]int{1, 2, 3, 4, 5, 6}
	//fmt.Println(array5[-1])
	//fmt.Println(array5[10])

	//切片
	slices := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(slices))
	//append
	slices = append(slices, 7)
	fmt.Println(len(slices))

	//make
	sl0 := make([]int, 6, 10)
	fmt.Println(len(sl0))
	fmt.Println(cap(sl0))

	sl1 := make([]int, 6)
	fmt.Println(len(sl1))
	fmt.Println(cap(sl1))

	//array
	array5 := [5]int{0, 1, 2, 3, 4}
	sl2 := array5[2:4:5]
	fmt.Println(len(sl2))
	fmt.Println(cap(sl2))

	sl3 := array5[0:4]
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))
	fmt.Println(sl3[3])

	//slice
	fmt.Println("slice:")
	sl4 := sl3[2:3]
	fmt.Println(len(sl4))

	//自动扩容的坑
	u := [...]int{11, 12, 13, 14, 15}
	fmt.Println("array:", u) // [11, 12, 13, 14, 15]
	s := u[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	s = append(s, 24)
	fmt.Println("after append 24, array:", u)
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s = append(s, 25)
	fmt.Println("after append 25, array:", u)
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s = append(s, 26)
	fmt.Println("after append 26, array:", u)
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s[0] = 22
	fmt.Println("after reassign 1st elem of slice, array:", u)
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
}
