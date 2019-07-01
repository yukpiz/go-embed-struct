package main

import "log"

type app struct {
	Args
}

type Args struct {
	ID   int
	Name string
}

func main() {
	args := Args{
		ID:   1,
		Name: "yukpiz",
	}

	a := app{args}
	log.Printf("%+v\n", a)
	log.Println(a.ID)
	log.Println(a.Name)
}
