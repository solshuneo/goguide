package main

import "fmt"

func main() {
	// var name type = expression
	// var name string = "shuneo"
	// fmt.Println(name)

	// var age int = 18
	// fmt.Println("age:", age)

	// var gpa float64 = 4.0
	// fmt.Printf("%.1f\n", gpa)

	// var isMale bool = true
	// fmt.Println(isMale)

	// // Male -> Female
	// // isMale = true -> false
	// isMale = false
	// fmt.Println(isMale)

	// zerovalue
	// var zeroValueInt int
	// var zeroValueString string
	// var zeroValueBool bool
	// var zerovalueFloat64 float64

	// fmt.Println("int", zeroValueInt)
	// fmt.Printf("string \"%s\"\n", zeroValueString)
	// fmt.Println("bool:", zeroValueBool)
	// fmt.Println("float64:", zerovalueFloat64)

	// const
	var name string = "shuneo"
	name = "shuneo2"
	const NAME string = "shuneo" // cannot reassign
	fmt.Println(name, NAME)
}
