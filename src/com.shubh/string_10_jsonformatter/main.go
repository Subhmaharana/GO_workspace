package main

import (
	"fmt"
)

func main() {
	fmt.Println(prettyJSON("[\"foo\", {\"bar\":[\"baz\",null,1.0,2]}]"))
}

func prettyJSON(A string) []string {
	var resultArr []string
	maxTabCount := 0
	for i, c := range A {
		fmt.Println(i, "==>", string(c))

		if c is openBrac
		    maxTabCount++
			resultArr = append(resultArr,string(c))
			newString := ""
			for index := 0; index < maxTabCount; index++ {
				newString := append(newString,"\t")
			}
			else if c 


	}

	// resultArr = append(resultArr, str)
	return resultArr
}
