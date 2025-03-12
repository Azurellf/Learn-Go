package main

import (
	"fmt"
)

type Position struct {
	x float64
	y float64
}

func main() {
	var m0 map[int]string
	fmt.Println(m0 == nil)

	m1 := map[int]string{}
	fmt.Println(m1 == nil)

	//复杂类型的字面初始化
	m2 := map[int][]string{
		1: []string{"value1", "value2"},
		2: []string{"Li", "Azure"},
	}
	fmt.Printf("%d %d\n", len(m2[2]), cap(m2[2]))

	m3 := map[Position]string{
		Position{1.23, 2.34}: "home",
		Position{3.45, 4.56}: "school",
	}
	fmt.Println(m3[Position{1.23, 2.34}])
	//语法糖
	m4 := map[Position]int{
		{5.67, 6.78}: 4,
		{7.89, 8.9}:  5,
	}
	fmt.Println(m4[Position{5.67, 6.78}])

	//make
	//m5 := make(map[string]int, 8)，指定容量
	//fmt.Println(cap(m5))，没有cap操作

	//insert
	m4[Position{192.1, 192.168}] = 6
	//len
	fmt.Println(len(m4))
	//get
	value, ok := m4[Position{1.1, 2.2}]
	if !ok {
		fmt.Println(value)
	}
	//delete
	fmt.Println(m2)
	delete(m2, 2)
	fmt.Println(m2)
	//遍历
	fmt.Println("{")
	for k, v := range m3 {
		fmt.Printf("%+v %s\n", k, v)
	}
	fmt.Println("}")

}
