package plusone

func plusOne(digits []int) []int {
    carry := 1
    n := len(digits)
    
    for i := n-1; i>=0; i-- {
        temp := digits[i] + carry
        if temp > 9 {
            digits[i] = 0
            
            if i == 0 {
                digits = append([]int{1}, digits...)
                return digits
            }
            
            continue
        }
        digits[i] = temp
        break
    }
    
    return digits
}