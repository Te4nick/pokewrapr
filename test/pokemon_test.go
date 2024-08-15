package test

import (
	"fmt"
	"testing"

	"github.com/Te4nick/pokewrapr"
)

func TestPokemonById(t *testing.T) {
	const id1 int = 1
	const name1 string = "bulbasaur"
	result, err := pokewrapr.Pokemon(id1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if result.Name != name1 {
		t.Fatalf(fmt.Sprintf("Awaited: %s, got: %s", name1, result.Name))
	}

	const id2 uint = 4
	const name2 string = "charmander"
	result, err = pokewrapr.Pokemon(id2)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if result.Name != name2 {
		t.Fatalf(fmt.Sprintf("Awaited: %s, got: %s", name2, result.Name))
	}
}

func TestPokemonByStringName(t *testing.T) {
	const name string = "charmander"
	result, err := pokewrapr.Pokemon(name)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if result.Name != name {
		t.Fatalf(fmt.Sprintf("Awaited: %s, got: %s", name, result.Name))
	}
}
