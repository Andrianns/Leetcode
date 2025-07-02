func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int) // map[num] = index

    for i, num := range nums {
        complement := target - num
        if index, found := numMap[complement]; found {
            return []int{index, i}
        }
        numMap[num] = i
    }

    return []int{} // jika tidak ada pasangan yang cocok
}