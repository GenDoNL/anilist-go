package anilistgo

import "testing"

func TestCharacterFullStack(t *testing.T) {
	query := "query ($search: String) { Character (search: $search) { id name {first last full} } }"

	variables := struct {
		Search string `json:"search"`
	}{
		"Kazuto",
	}

	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	c, err := a.Character(query, variables)
	if err != nil {
		t.Fatal(err)
	}

	if c.Id != 36765 {
		t.Fatalf("Wrong id returned: %d", c.Id)
	}

	if c.Name.Full != "Kazuto Kirigaya" {
		t.Fatalf("Wrong name returned: %s", c.Name.Full)
	}
}