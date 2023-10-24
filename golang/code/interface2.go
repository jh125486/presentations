package main

import (
	"fmt"
)

// START OMIT
type Pokemon interface {
	Attack() string
	SayName() string
	Type() string
}

func DescribePokemon(p Pokemon) {
	fmt.Printf("%v [%v] attacks with %v\n", p.SayName(), p.Type(), p.Attack())
}

type Pikachu struct{}             // Pikachu Implementation
func (p Pikachu) Attack() string  { return "Thunderbolt" }
func (p Pikachu) SayName() string { return "Pikachu" }
func (p Pikachu) Type() string    { return "Electric" }

type Bulbasaur struct{}             // Bulbasaur Implementation
func (b Bulbasaur) Attack() string  { return "Vine Whip" }
func (b Bulbasaur) SayName() string { return "Bulbasaur" }
func (b Bulbasaur) Type() string    { return "Grass/Poison" }

func main() {
	DescribePokemon(Pikachu{})
	DescribePokemon(Bulbasaur{})
}

// END OMIT
