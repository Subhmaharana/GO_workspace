package main

import "fmt"

func main() {

    intArr := []int{3, 2, 1, 4, 5, 6, 7}
    fmt.Println(maxset(intArr))
}



/**
 * @input A : Integer array
 * @input B : Integer array
 * 
 * @Output Integer
 */
 
func coverPoints(A []int , B []int )  (int) {
    result := 0
    for i := 1; i < len(A); i++ {
        result = result + max(abs(A[i] - A[i-1]),abs(B[i] - B[i-1]))
    }
        
    return result;
}

func max(a int, b int) int{
    if a > b {
        return a
    }
    return b
}

func abs(a int) int{
    if a < 0 {
        return a * -1
    }
    return a
}