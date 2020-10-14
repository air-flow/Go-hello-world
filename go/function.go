package main

import "fmt"

func main() {
	// var a string
	// fmt.Scan(&a)
	// fmt.Println("%s", a)
	// stdin := bufio.NewScanner(os.Stdin)
	// stdin.Scan()
	// text := stdin.Text()
	// fmt.Println(text)
	sosuu()
}

// エラトステネスの篩
func sosuu() {
	num := [98]int{}
	for i := 2; i < 100; i++ {
		num[i-2] = i

	}
	fmt.Println(num)
}

// C:\Users\user\Documents\GitHub\Go-hello-world\go>go run function.go
// input:  test
// output: test
