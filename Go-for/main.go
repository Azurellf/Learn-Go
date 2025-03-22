package main

import (
	"fmt"
	"time"
)

func close() {
	m := []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func() {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}()
	}
	time.Sleep(time.Second * 10)
}

func close1() {
	m := []int{1, 2, 3, 4, 5}
	for i, v := range m {
		go func(i, v int) {
			time.Sleep(time.Second * 5)
			fmt.Println(i, v)
		}(i, v)
	}
	time.Sleep(time.Second * 10)
}

func main() {
	close()
}
