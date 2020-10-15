package main

import "fmt"

// エラトステネスの篩

func main() {
	Sosuu()
}

func Sosuu() {
	num := []int{}
	for i := 2; i < 121; i++ {
		num = append(num, i)
	}
	fmt.Println(num)

}
