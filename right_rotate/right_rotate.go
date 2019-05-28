package rightrotate

func rotate(nums []int, k int)  {
    n := len(nums)
    if n == 0 {
        return
    }
    for j:=1; j<=k; j++ {
        last := nums[n-1]
        previous := nums[0]
        for i:= 1; i<n; i++ {
            temp := nums[i]
            nums[i] = previous
            previous = temp 
        }

        nums[0] = last
    }
}

func rotate2(nums []int, k int)  {
    n := len(nums)
    arr := make([]int, len(nums)) 
    for i, _ := range nums {
        arr[(i+k) % n] = nums[i]
    }
    
    for i,_:= range nums {
        nums[i] = arr[i]
    }
}

