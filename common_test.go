package pokewrapr

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Te4nick/pokewrapr/cache"
	"github.com/Te4nick/pokewrapr/entity"
)

func TestConvertIdOrName(t *testing.T) {
	const n int8 = 1

	if s := convertIdOrName(n); s != "1" {
		t.Fatalf("Could not convert int8 id")
	}

	if s := convertIdOrName(int16(n)); s != "1" {
		t.Fatalf("Could not convert int16 id")
	}

	if s := convertIdOrName(int32(n)); s != "1" {
		t.Fatalf("Could not convert int32 id")
	}

	if s := convertIdOrName(int64(n)); s != "1" {
		t.Fatalf("Could not convert int64 id")
	}

	if s := convertIdOrName(int(n)); s != "1" {
		t.Fatalf("Could not convert int id")
	}

	if s := convertIdOrName(uint8(n)); s != "1" {
		t.Fatalf("Could not convert uint8 id")
	}

	if s := convertIdOrName(uint16(n)); s != "1" {
		t.Fatalf("Could not convert uint16 id")
	}

	if s := convertIdOrName(uint32(n)); s != "1" {
		t.Fatalf("Could not convert uint32 id")
	}

	if s := convertIdOrName(uint64(n)); s != "1" {
		t.Fatalf("Could not convert uint64 id")
	}

	if s := convertIdOrName(uint(n)); s != "1" {
		t.Fatalf("Could not convert uint id")
	}

	if s := convertIdOrName("name"); s != "name" {
		t.Fatalf("Could not convert string name")
	}
}

func TestHTTPStatusErrorOk(t *testing.T) {
	const code int = 500
	const url string = "nonexist"
	errMsg := fmt.Sprintf("HTTP %s Status Code %d", url, code)

	err := HTTPStatusError{
		StatusCode: code,
		URL:        url,
	}

	if err.Error() != errMsg {
		t.Errorf("Expexted: '%s', got: '%s'", errMsg, err.Error())
	}
}

func TestFetchByIdWithDefaultCacheOk(t *testing.T) {
	const name = "bulbasaur"
	var result entity.Pokemon

	err := fetch("/pokemon/1/", &result)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error())
	}

	if result.Name != name {
		t.Fatalf("Awaited: %s, got: %s", name, result.Name)
	}
}

func TestFetchByIdWithNoTTLCacheOk(t *testing.T) {
	const name = "bulbasaur"
	var result entity.Pokemon

	CacheTTL = 0
	defer func() { CacheTTL = cache.CacheDefaultTTL }()

	err := fetch("/pokemon/1/", &result)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error())
	}

	if result.Name != name {
		t.Fatalf("Awaited: %s, got: %s", name, result.Name)
	}
}

func TestFetchByIdNoCacheOk(t *testing.T) {
	const name = "bulbasaur"
	var result entity.Pokemon

	UseCache = false
	defer func() { UseCache = true }()

	err := fetch("/pokemon/1/", &result)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error())
	}

	if result.Name != name {
		t.Fatalf("Awaited: %s, got: %s", name, result.Name)
	}
}

func TestFetchByIdFullUrlOk(t *testing.T) {
	const name = "bulbasaur"
	var result entity.Pokemon

	err := fetch("https://pokeapi.co/api/v2/pokemon/1/", &result)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error())
	}

	if result.Name != name {
		t.Fatalf("Awaited: %s, got: %s", name, result.Name)
	}
}

func TestFetchByIdNoSlashEndpointOk(t *testing.T) {
	const name = "bulbasaur"
	var result entity.Pokemon

	err := fetch("pokemon/1/", &result)
	if err != nil {
		t.Fatalf("Unexpected error: " + err.Error())
	}

	if result.Name != name {
		t.Fatalf("Awaited: %s, got: %s", name, result.Name)
	}
}

func TestFetchWrongEndpointError(t *testing.T) {
	var result entity.Pokemon

	err := fetch("/nonexist/", &result)
	if httpErr, ok := err.(HTTPStatusError); !ok || httpErr.StatusCode != 404 {
		t.Fatalf("Expected http.Get() error " + err.Error())
	}
}

func TestFetchJsonMarshalError(t *testing.T) {
	err := fetch("/pokemon/1000/", nil)
	if err == nil {
		t.Fatalf("Expected 'json: Unmarshal(nil)'")
	}

	switch err.(type) {
	case *json.InvalidUnmarshalError:
		return
	default:
		t.Fatalf("Expected 'json: Unmarshal(nil)'")
	}
}
