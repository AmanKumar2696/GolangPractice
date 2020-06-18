package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Integer Array:")
	i := 0
	length := len(arr)
	for ; i < length; i++ {
		fmt.Println("Index: ", i, " Value: ", arr[i])
	}
	sum := 0
	for i = 0; i < length; i++ {
		sum += arr[i]
	}
	fmt.Println("Sum of all Elements: ", sum)

	slc := []int{10, 20, 30}
	fmt.Println("Slice", slc)
	slc = append(slc, 40)
	fmt.Println("Slice appended: ", slc)
	fmt.Println("Slicing:\nslc[2:]: ", slc[2:])
}
