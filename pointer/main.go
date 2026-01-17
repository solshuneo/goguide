package main

import "fmt"

func main() {
	// pointer: address of a var
	var aliasName string = "shuneo"
	// var realName string = "Thien"
	var addressName *string = &aliasName
	fmt.Println(addressName)

}
