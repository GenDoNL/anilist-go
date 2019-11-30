package anilistgo

import "testing"

func TestAnimeFullStack(t *testing.T) {
	query := "query ($search: String, $type: MediaType) { Media (search: $search, type: $type) { id idMal title { romaji english native } type genres} }"

	variables := struct {
		Search string `json:"search"`
		Type string `json:"type"`
	}{
		"Vinland Saga",
		"ANIME",
	}

	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	m, err := a.Media(query, variables)
	if err != nil {
		t.Fatal(err)
	}

	if m.Id != 101348 {
		t.Fatalf("Wrong id was returned: %d", m.Id)
	}

	if m.IdMal != 37521 {
		t.Fatalf("Wrong idMal was returned: %d", m.IdMal)
	}

	if m.Title.Romaji != "Vinland Saga" {
		t.Fatalf("Wrong name was returned: %s", m.Title.Romaji)
	}

	if len(m.Genres) == 0 {
		t.Fatalf("No genres were returned")
	}
}