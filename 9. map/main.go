package main

import "fmt"

func main() {
	var mymap = make(map[string]string, 100)
	fmt.Printf("mymap: %v\n", mymap)
	mymap["vi"] = "Vietnam"
	mymap["cn"] = "China"
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}

}
