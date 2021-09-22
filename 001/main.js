/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
    const hashTalb = new Map()
    for (let i = 0; i < nums.length; i++) {
        const num1 = nums[i];
        const num2 = target - nums[i];
        if (hashTalb.has(num2)) {
            return [hashTalb.get(num2), i]
        } else {
            hashTalb.set(num1, i)
        }
    }
};