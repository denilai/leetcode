package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	//mask := make([]bool, len(nums))
	var find int
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			find++
			if i == len(nums)-find {
				continue
			} else {
				nums[i], nums[len(nums)-find] = nums[len(nums)-find], nums[i]
			}
		}
	}
	return len(nums) - find
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	arr1 := make([]int, m)
	copy(nums1, arr1)
	fmt.Println(arr1)
	for i, j := 0, 0; i+j < m+n; {
		fmt.Printf("i = %v, j = %v\n", i, j)
		if i >= m {
			nums1[i+j] = nums2[j]
			fmt.Println(nums2[j])
			j++
		} else if j >= n {
			nums1[i+j] = arr1[i]
			fmt.Println(arr1[i])
			i++
		} else {
			nums1[i+j] = min(arr1[i], nums2[j])
			fmt.Printf("%v, %v, min = %v\n", arr1[i], nums2[j], min(arr1[i], nums2[j]))
			if arr1[i] < nums2[j] {
				i++
			} else {
				j++
			}
		}
		//fmt.Printf("i = %v, j = %v\n", i, j)
	}
}

func main1() {
	//nums1 := []int{1, 2, 3, 0, 0, 0}
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Println(nums1, nums2)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1, nums2)
}

func main() {
	//nums := []int{3, 2, 2, 3}
	nums1 := []int{2, 0, 1, 6, 8, 2, 2, 3, 2, 0, 4, 2}
	val1 := 2
	//val := 3
	fmt.Println(nums1)
	k := removeElement(nums1, val1)
	fmt.Println(nums1, k)
}
