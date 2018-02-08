package main

import "fmt"

const p = "My Name is Subhransu"

func main() {
	fmt.Println("Hello World By Subhransu")
	const q = 100
	fmt.Println("Printing constants ", p)
	fmt.Println("Printing constants ", q)

	//scope 
	x := 0
	fmt.Println(x)
	{
		x := 43
		fmt.Println(x)
		y := "Inside scope.."
		fmt.Println(y)
	}
	fmt.Println(x)



	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

}

// func getFunc(){
// 	res := http.Get("http://www.mcleods.com/")
// 	page := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	fmt.Println("%s",page)
// }