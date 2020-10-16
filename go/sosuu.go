package main

import (
	"fmt"
	"math"
)

// エラトステネスの篩

func main() {
	searchlist := Step1()
	sosuulist := []int{}
	maxvalue := searchlist[len(searchlist)-1]
	// Step3
	for searchlist[0] < Step3(maxvalue) {
		searchlist, sosuulist = Step2(searchlist, sosuulist)
	}
	// Step4
	if len(searchlist) > 0 {
		sosuulist = append(sosuulist, searchlist...)
	}
	// fmt.Print(searchlist)
	fmt.Print(sosuulist)
}

// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97 101 103 107 109 113]
// [Done] exited with code=0 in 0.919 seconds

// Step1 探索リストに2からxまでの整数を昇順で入れる。
func Step1() []int {
	num := []int{}
	for i := 2; i < 121; i++ {
		num = append(num, i)
	}
	// fmt.Println(num)
	return num
}

// [2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108 109 110 111 112 113 114 115 116 117 118 119 120]
// [Done] exited with code=0 in 1.745 seconds

// Step2 探索リストの先頭の数を素数リストに移動し、その倍数を探索リストから篩い落とす。
func Step2(searchlist []int, sosuulist []int) ([]int, []int) {
	sosuulist = append(sosuulist, searchlist[0])
	temp := []int{}
	for i := 0; i < len(searchlist); i++ {
		if searchlist[i]%sosuulist[len(sosuulist)-1] != 0 {
			temp = append(temp, searchlist[i])
		}
	}
	return temp, sosuulist
}

// [3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59 61 63 65 67 69 71 73 75 77 79 81 83 85 87 89 91 93 95 97 99 101 103 105 107 109 111 113 115 117 119][2]
// [Done] exited with code=0 in 4.426 seconds

// Step3 上記の篩い落とし操作を探索リストの先頭値がxの平方根に達するまで行う。
func Step3(x int) int {
	return int(math.Sqrt(float64(x)))
}
