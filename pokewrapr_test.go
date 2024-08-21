package pokewrapr

import (
	"fmt"
	"testing"

	"github.com/Te4nick/pokewrapr/entity"
)

func TestResourceListByEndpoint(t *testing.T) {
	result, err := ResourceList(entity.PokemonEndpoint, "", "")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if result.Next != baseURL+"/pokemon/?offset=20&limit=20" || result.Previous != "" || len(result.Results) != 20 {
		t.Fatalf("Unexpected result: next=%s, previous=%s, count=%d", result.Next, result.Previous, result.Count)
	}
}

func TestResourceListByStringEndpointAndIntegerPaging(t *testing.T) {
	const (
		endpoint string = "location-area"
		offset   uint   = 60
		limit    uint   = 40
	)
	var (
		nextURL     string = fmt.Sprintf("%s/%s/?offset=%d&limit=%d", baseURL, endpoint, offset+limit, limit)
		previousURL string = fmt.Sprintf("%s/%s/?offset=%d&limit=%d", baseURL, endpoint, offset-limit, limit)
	)

	result, err := ResourceList(endpoint, offset, limit)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if result.Next != nextURL || result.Previous != previousURL || len(result.Results) != int(limit) {
		t.Fatalf("Unexpected result: next=%s, previous=%s, count=%d", result.Next, result.Previous, result.Count)
	}
}

func TestLocationAreaByStringName(t *testing.T) {
	const name string = "canalave-city-area"
	result, err := LocationArea(name)
	if err != nil {
		t.Fatalf("Unexpected error: '%s'", err.Error())
	}

	if result.Name != name {
		t.Fatalf("Awaited: %s, got: %s", name, result.Name)
	}
}

func TestLocationAreaById(t *testing.T) {
	const id1 int = 1
	const name1 string = "canalave-city-area"
	result, err := LocationArea(id1)
	if err != nil {
		t.Fatalf("Unexpected error: '%s'", err.Error())
	}

	if result.Name != name1 {
		t.Fatalf("Awaited: %s, got: %s", name1, result.Name)
	}

	const id2 uint = 4
	const name2 string = "sunyshore-city-area"
	result, err = LocationArea(id2)
	if err != nil {
		t.Fatalf("Unexpected error: '%s'", err.Error())
	}

	if result.Name != name2 {
		t.Fatalf("Awaited: %s, got: %s", name2, result.Name)
	}
}

func TestPokemonById(t *testing.T) {
	const id1 int = 1
	const name1 string = "bulbasaur"
	result, err := Pokemon(id1)
	if err != nil {
		t.Fatalf("Unexpected error: '%s'", err.Error())
	}

	if result.Name != name1 {
		t.Fatalf("Awaited: %s, got: %s", name1, result.Name)
	}

	const id2 uint = 4
	const name2 string = "charmander"
	result, err = Pokemon(id2)
	if err != nil {
		t.Fatalf("Unexpected error: '%s'", err.Error())
	}

	if result.Name != name2 {
		t.Fatalf("Awaited: %s, got: %s", name2, result.Name)
	}
}

func TestPokemonByStringName(t *testing.T) {
	const name string = "charmander"
	result, err := Pokemon(name)
	if err != nil {
		t.Fatalf("Unexpected error: '%s'", err.Error())
	}

	if result.Name != name {
		t.Fatalf("Awaited: %s, got: %s", name, result.Name)
	}
}
