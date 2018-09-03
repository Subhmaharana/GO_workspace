package main

import (
	"fmt"
)

func main() {
	A := " ef8938 9203 3"
	result := atoi(A)
	fmt.Println(result)
}

const INT_MAX = int(^uint32(0) >> 1)
const INT_MIN = -INT_MAX - 1

func atoi(A string) int {
	result := 0
	start, end := 0, len(A)
	isNegative := false

	for start < end && A[start] == ' ' {
		start++
	}

	for i := start; i < end; i++ {
		if A[i] == '-' {
			isNegative = true
			continue
		}
		if A[i] == '+' {
			continue
		}

		if A[i] == ' ' {
			break
		}

		if A[i] >= '0' && A[i] <= '9' {
			result = (result * 10) + int(A[i]-'0')
			if result > INT_MAX {
				if isNegative {
					return INT_MIN
				}
				return INT_MAX
			}
		} else {
			break
		}
	}

	if isNegative {
		return -result
	}
	return result
}
