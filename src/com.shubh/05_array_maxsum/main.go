package main

import "fmt"

func main() {

    intArr := []int{3, 2, 1, 4, 5, 6, 7}
    fmt.Println(maxset(intArr))
}

func maxset(A []int )  ([]int) {    
    result := make([]int,0)
    maxResult := make([]int,0)
    sum,maxSum := 0,0
    for _,element := range A {
        if element >= 0 {
            sum = sum + element;
            result = append(result,element)
        }else{
            sum = 0
            result = make([]int,0)
        }
        
        if (maxSum < sum) || (maxSum == sum && len(result)>len(maxResult)){
            maxSum = sum
            maxResult = result
        }
    }
    return maxResult
}