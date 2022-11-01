package main

import "fmt"

func main() {
	dict := make(map[string]string)
	dict["Jovan"] = "+38195894526"
	dict["Teodora"] = "+3818246572"
	dict["Dzem"] = "+38122365"

	keys := []string{"Jovan", "Teodora", "Dzem"}

	for i := 0; i < len(keys); i++ {
		fmt.Println("Person with name", keys[i], "has a phone number of", dict[keys[i]])
	}
}
