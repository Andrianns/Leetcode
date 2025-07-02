func romanToInt(s string) int {
    romanMap := map[byte]int{
        'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
    }

    total:=0

    for i:=0 ; i<len(s) ; i++{
        curr := romanMap[s[i]]
        if i+1 <len(s) && curr < romanMap[s[i+1]]{
            total-=curr
        }else{
            total+=curr
        }
    }

    return total
}