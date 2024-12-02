func generateParenthesis(n int) []string {
        var result []string
    var backtrack func(current string, openCount, closeCount int)

    backtrack = func(current string, openCount, closeCount int) {

        if len(current) == 2*n {
            result = append(result, current)
            return
        }


        if openCount < n {
            backtrack(current+"(", openCount+1, closeCount)
        }


        if closeCount < openCount {
            backtrack(current+")", openCount, closeCount+1)
        }
    }


    backtrack("", 0, 0)
    
    return result
}