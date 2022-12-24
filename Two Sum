func twoSum(nums []int, target int) []int {
    res := make([]int, 2)
    a := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        if val, ok := a[nums[i]]; ok {
            res[0] = i
            res[1] = val
        } else {
            a[target - nums[i]] = i
        }
    }
    return res
}
