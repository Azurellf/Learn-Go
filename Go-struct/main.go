package main

import (
	"fmt"
)

type T1 int
type T2 T1
type T3 string

type Person struct {
	Name  string
	Phone string
	Addr  string
}

type Book struct {
	BookName string
	Author   Person
}

// 嵌入式字段
type Print struct {
	Name string
	Person
}

func main() {
	var n1 T1
	var n2 T2 = 5
	n1 = T1(n2)
	println(n1)
	//底层类型不同无法转换
	//var s T3 = "string"
	//n1 = T1(s)

	var man Person
	println(man.Name)
	var man1 = Person{"li", "123", "am"}
	fmt.Println(man1)

	var book1 = Book{"Gone with wind", Person{"lin", "234", "cn"}}
	println(book1.Author.Phone)

	//使用嵌入字段时，field:value初始化
	paint := Print{
		Person: Person{
			Name:  "132",
			Phone: "123",
			Addr:  "213",
		},
	}
	fmt.Println(paint)
}
