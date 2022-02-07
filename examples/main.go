package main

import (
	"fmt"

	"github.com/wander4747/faker-go"
)

func main() {
	f := faker.New("pt_BR")

	name := f.Person().Age()
	fmt.Println(name)
}
