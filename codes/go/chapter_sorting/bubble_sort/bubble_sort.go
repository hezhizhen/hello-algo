// File: bubble_sort.go
// Created Time: 2022-12-06
// Author: Slone123c (274325721@qq.com)

package bubble_sort

/* 冒泡排序 */
func bubbleSort(nums []int) {
	// 外循环：待排序元素数量为 n-1, n-2, ..., 1
	for i := len(nums) - 1; i > 0; i-- {
		// 内循环：冒泡操作
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

/* 冒泡排序（标志优化）*/
func bubbleSortWithFlag(nums []int) {
	// 外循环：待排序元素数量为 n-1, n-2, ..., 1
	for i := len(nums) - 1; i > 0; i-- {
		flag := false // 初始化标志位
		// 内循环：冒泡操作
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				// 交换 nums[j] 与 nums[j + 1]
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true // 记录交换元素
			}
		}
		if flag == false { // 此轮冒泡未交换任何元素，直接跳出
			break
		}
	}
}
