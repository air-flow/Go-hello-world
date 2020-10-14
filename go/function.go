package main

import (
	"fmt"
)

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

func sosuu() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// C:\Users\user\Documents\GitHub\Go-hello-world\go>go run function.go
// input:  test
// output: test
