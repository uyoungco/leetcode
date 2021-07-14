package main

import "fmt"

// 001 两数之和

// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

func TwoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target - x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}


func main() {
	nums, target := []int{2,7,11,15}, 22
	fmt.Println(TwoSum(nums, target))
	//fmt.Println(target)
}
