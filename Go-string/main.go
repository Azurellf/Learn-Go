package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//下标操作
	var s = "牛的兄弟"
	fmt.Printf("0x%x\n", s[0])

	//迭代1
	for i := 0; i < len(s); i++ {
		fmt.Printf("index: %d, value: 0x%x\n", i, s[i])
	}
	//迭代2
	for i, v := range s {
		fmt.Printf("index %d, value: 0x%x\n", i, v)
	}

	//获取长度
	var t1 = "test1"
	var t2 = "测试2"
	fmt.Println(utf8.RuneCountInString(t1))
	fmt.Println(utf8.RuneCountInString(t2))
	//字符串连接1
	ss := "Rob "
	ss += "Pike"
	fmt.Println(ss)
	//字符串连接2
	var builder strings.Builder
	builder.WriteString("Robert ")
	builder.WriteString("Griesemer")
	fmt.Println(builder.String())
	//字符串连接3
	strs := []string{"Ken", "Thompson"}
	fmt.Println(strings.Join(strs, " "))
	//字符串连接4
	fmt.Println(fmt.Sprintf("%s %s", "Li", "LinFeng"))

	//比较
	c1 := "Hello World"
	c2 := "Hello" + " World"
	fmt.Println(c1 == c2)

	//转化
	t := "好累"
	//字符串->rune切片
	rt := []rune(t)
	fmt.Printf("0x%x\n", rt)
	//rune切片->字符串
	fmt.Println(string(rt))
	//字符串->byte切片
	bt := []byte(t)
	fmt.Printf("0x%x\n", bt)
	//byte切片->字符串
	fmt.Println(string(bt))
}
