package main

import "fmt"

func main() {
	arr := []int{1, 2}
	arr2 := arr[1:]

	fmt.Println(cap(arr))

	arr = append(arr, 3)

	/*
	   每次调用appendInt函数，必须先检测slice底层数组是否有足够的容量来保存新添加的元素。如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），将新添加的y元素复制到新扩展的空间，并返回slice。因此，输入的x和输出的z共享相同的底层数组。

	   如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间，然后添加y元素。结果z和输入的x引用的将是不同的底层数组。
	*/
	fmt.Println(arr, arr2)

	// pop
	arr = arr[:len(arr)-1]
}
