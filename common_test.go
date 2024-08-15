package pokewrapr

import "testing"

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
