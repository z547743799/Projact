package main

import "fmt"

var a []int

func main() {
	a = []int{12, 44, 3, 245, 2, 3, 12}
	qsort(0, len(a)-1)
	fmt.Println(a)
}
func qsort(left int, rigth int) {
	if left >= rigth {
		return
	}
	i := left
	j := rigth
	temp := a[left]
	for i != j {
		for a[j] >= temp && j > i {
			j--
		}
		for a[i] <= temp && j > i {
			i++
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
			fmt.Println(a)
		}
	}

	a[i], a[left] = temp, a[i]
	qsort(left, i-1)
	qsort(i+1, rigth)

}
