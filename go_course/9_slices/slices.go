package main

import (
	"fmt"
	"slices"
)

func main() {

	// unintialized slice in nil

	//var nums []int
	//fmt.Println(nums == nil)
	//fmt.Println(len(nums))

	// var list = make([]int, 0, 5)

	// // cap ->  capacity ->  maxm number of element can fit
	// fmt.Println(list)
	// fmt.Println(len(list))
	// fmt.Println(cap(list))
	// list = append(list, 1)
	// fmt.Println(list)
	// fmt.Println(len(list))
	// fmt.Println(cap(list))
	// list = append(list, 2)
	// list = append(list, 2)
	// list = append(list, 2)
	// fmt.Println(list)
	// fmt.Println(len(list))
	// fmt.Println(cap(list))
	// list = append(list, 2)
	// list = append(list, 2)
	// list = append(list, 2)
	// list = append(list, 2)
	// list = append(list, 2)
	// list = append(list, 2)
	// list = append(list, 2)
	// fmt.Println(list)
	// fmt.Println(len(list))
	// fmt.Println(cap(list))

	// copy(destination,source)

	// slice operator

	// var nums = []int{1, 2, 3}

	// fmt.Println(nums[1:])

	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}

	fmt.Println(slices.Equal(nums1, nums2))

}
