package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var a string
	// fmt.Scan(&a)
	// fmt.Println("%s", a)
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	fmt.Println(text)
}

// C:\Users\user\Documents\GitHub\Go-hello-world\go>go run function.go
// input:  test
// output: test
