package main

func case1() int {
	println("eval case1 expr")
	return 1
}

func case2_1() int {
	println("eval case2_1 expr")
	return 0
}

func case2_2() int {
	println("eval case2_2 expr")
	return 2
}

func case3() int {
	println("eval case3 expr")
	return 3
}

func switchexpr() int {
	println("switch expr")
	return 2
}

type person struct {
	name string
	age  int
}

func main() {
	switch switchexpr() {
	case case1():
		println("exec case1")
	//表达式列表
	case case2_1(), case2_2():
		println("exec case2_1 and case2_2")
	case case3():
		println("exec case3")
	default:
		println("exec default")
	}

	//类型不局限于整形
	p := person{"Tom", 13}
	switch p {
	case person{"Tom", 13}:
		println("it's tom")
	case person{"Lee", 30}:
		println("it's Lee")
	}

	//type switch
	var x interface{} = "hello"
	//v存储的是值信息
	switch v := x.(type) {
	case nil:
	case int:
		println("the type of x is int, value is ", v)
	case string:
		println("the type of x is string, value is ", v)
	case bool:
	default:
	}

	var nums = []int{5, 19, 37, 20, 12}
	var firstEven = -1
	for i := 0; i < len(nums); i++ {
		switch nums[i] % 2 {
		case 0:
			firstEven = nums[i]
			println("find first Even is", firstEven)
			//break跳出最内层的for，switch，select
			break
		case 1:
		}
	}
}
