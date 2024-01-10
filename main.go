package main

import (
	"fmt"

	"github.com/yangju1999/bc_project/person"
)

func main() {
	nico := person.Person{}
	nico.SetDetails("nico", 12)
	fmt.Println(nico)
}
