package main

import "fmt"

// エラトステネスの篩

func main() {
	search_list := Step1()
	sosuu_list := []int{}
	search_list, sosuu_list = Step2(search_list, sosuu_list)
	fmt.Print(search_list)
	fmt.Print(sosuu_list)
}

// CreateValueList 1から120までの配列を生成する
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

func Step2(search_list []int, sosuu_list []int) ([]int, []int) {
	sosuu_list = append(sosuu_list, search_list[0])
	temp := []int{}
	for i := 0; i < len(search_list); i++ {
		if search_list[i]%sosuu_list[len(sosuu_list)-1] != 0 {
			temp = append(temp, search_list[i])
		}
	}
	return temp, sosuu_list
}

// [3 5 7 9 11 13 15 17 19 21 23 25 27 29 31 33 35 37 39 41 43 45 47 49 51 53 55 57 59 61 63 65 67 69 71 73 75 77 79 81 83 85 87 89 91 93 95 97 99 101 103 105 107 109 111 113 115 117 119][2]
// [Done] exited with code=0 in 4.426 seconds
