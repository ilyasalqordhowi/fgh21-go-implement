package fazz

import "fmt"

func Matrix() {

	var matrix = [][][]int{
		[][]int{[]int{3, 0}, []int{3, 1, 3, 3}},
		[][]int{[]int{3, 0}},
		[][]int{[]int{3, 0, 3}},
		[][]int{[]int{3, 0, 10}},
	}
	fmt.Println(matrix[0][1][2])
	fmt.Println(matrix[3][0][2])
}