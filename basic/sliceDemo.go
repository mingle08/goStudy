package main
/*
数学上的开区间和闭区间
开区间指的是区间边界的两个值不包括在内；（a,b）
闭区间指的是区间边界的两个值包括在内。[a,b]
半开半闭区间:开区间一边的边界值不包括在内，而闭区间一边的边界值包括在内。[a,b）、（a,b]
如下：
[a,b] a<=x<=b 取值包括a、b
（a,b）a<x<b 取值不包括a、b
[a,b） a<=x<b 取值包括a，不包括b
（a,b] a<x<=b 取值不包括a，包括b

*/

import (
	"fmt"
)

func main(){
	nums := []int{0,1,2,3,4,5,6,7}
	fmt.Println("nums=", nums)
	
	fmt.Println("nums[0]=", nums[0], ", nums[3]=", nums[3])
    
    // 切片是前闭后开  [a, b)，即       a <= x < b， 取a，不取b
    fmt.Println("nums[:2]=", nums[:2])
    
    fmt.Println("nums[1:3]=", nums[1:3])
    
    fmt.Println("nums[4:]=", nums[4:])
    
    fmt.Println("len=", len(nums), ", capacity=", cap(nums))
}

