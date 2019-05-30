package twosum

func twoSum(nums []int, target int) []int {
    l := len(nums)
    var result []int
    m := make(map[int]int, len(nums))
    for i:=0; i<l; i++ {
        
        if v, ok := m[nums[i]]; ok && nums[i]*2==target  {
           result := append(result, v, i)
            return result 
        }
        
        m[nums[i]] = i
        
        d := target - nums[i]
        
        if v, ok := m[d]; ok && i != v {
            result := append(result, v, i)
            return result
        }
        
       
    }
    
    return result
}
