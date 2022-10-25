package main

import "fmt"

type hero struct {
	name        string
	home_planet string
	power_level int
}

type hero_collection struct {
	collection []hero
}

func (hc *hero_collection) insert_hero(h hero) {
	hc.collection = append(hc.collection, h)
}

func (hc *hero_collection) print_heroes() {
	var i int
	for i = 0; i < len(hc.collection); i++ {
		fmt.Println("Printing hero", hc.collection[i])
	}
}

func main() {
	var heroes hero_collection
	hero1 := hero{name: "superman", home_planet: "krypton", power_level: 9000}
	hero2 := hero{name: "goku", home_planet: "vegeta", power_level: 9001}
	hero3 := hero{name: "batman", home_planet: "earth", power_level: 8500}
	heroes.insert_hero(hero1)
	heroes.insert_hero(hero2)
	heroes.insert_hero(hero3)
	heroes.print_heroes()
}
