func removeElement(nums []int, val int) int {
    k := 0

    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[k] = nums[i]
            k++
        }
    }
//    3 != 3  =false -> skip [3,2,2,3]
//    2 != 3 = true -> nums[k(0)] = nums[i(1)] = [2,2,2,3], k = 1
//    2 != 3 = true -> nums[k(1)] = nums[i(2)] = [2,2,2,3], k = 2
//    3 != 3 = false -> skip [2,2,2,3] , k = 2
    return k
}