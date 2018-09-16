//快速排序

package main

import (
	"fmt"
)

var a []int //需要排序的数组

func main() {
	a = []int{14, 1, 42, 18, 9, 23, 24, 5, 10, 18}
	quicksort(0, len(a)-1)
	fmt.Println(a)
}

//left 需要排序的数组起始位置 right 最右边的那个数的位置
func quicksort(left int, right int) {

	if left >= right { //需要排序的起始位置大于或等于终止位置，就表明不再需要排序
		return
	}

	i := left       //左边的游标
	j := right      //右边的游标
	temp := a[left] //基数

	for i != j {
		for a[j] >= temp && i < j { //直到遇到了 <temp 的数就停下来
			j--
		}

		for a[i] <= temp && i < j { //直到遇到了 >temp 的数就停下来
			i++
		}

		//交换这两个数
		if i < j {
			a[i], a[j] = a[j], a[i]
			fmt.Println(a)
		}

	}

	//将基数归位
	a[i], a[left] = temp, a[i]
	fmt.Println("----")
	//递归处理基数左边未处理的
	quicksort(left, i-1)
	fmt.Println(a)
	//递归处理基数右边未处理的
	if i != len(a)-1 {
		quicksort(i+1, right)
	}

}
