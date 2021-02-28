func twoSum(nums []int, target int) []int {
    length := len(nums)
    i := 0
    for i <= length/2 {
        j := i + 1
        for j < length {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            } else {
                j = j + 1
            }
        }
        i = i + 1
    }
    return []int{}
}
