func twoSum(nums []int, target int) []int {
    // Create a map to store the complement of each number
    complementMap := make(map[int]int)
    
    // Loop through the array
    for i, num := range nums {
        // Calculate the complement (what we need to add to num to reach target)
        complement := target - num
        
        // Check if the complement exists in the map
        if index, found := complementMap[complement]; found {
            // If found, return the indices
            return []int{index, i}
        }
        
        // If not found, add the current number and its index to the map
        complementMap[num] = i
    }
    
    // Return empty array if no solution found
    return []int{}
}