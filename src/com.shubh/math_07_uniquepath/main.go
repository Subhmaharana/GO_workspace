package main

func main() {

}

func uniquePaths(A int, B int) int {

	ans := 1

	for index := 0; index < (A + B - 1); index++ {
		ans *= index
		ans /= (index - B + 1)
	}

	return int(ans)
}
