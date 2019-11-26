package anilistgo

import "testing"

func TestFullStack(t *testing.T) {
	n, err := New()

	if err != nil {
		t.Fatal(err)
	}

	res, err := n.Media(MediaVariables{SearchQuery: "Vineland Saga", Type: "ANIME", Page: 1, PerPage: 5})
	
	if err != nil {
		t.Fatal(err)
	}

	if res.Id != 101348 {
		t.Fatal("Wrong id was returned.")
	}

	if res.IdMal != 37521 {
		t.Fatal("Wrong idMal was returned")
	}

	if res.Title.Romaji != "Vinland Saga" {
		t.Fatal("Wrong name was returned.")
	}

	if len(res.Genres) == 0 {
		t.Fatal("No genres were returned.")
	}
}

func TestQueryByID(t *testing.T) {
	n, err := New()

	if err != nil {
		t.Fatal(err)
	}

	res, err := n.Media(MediaVariables{ID: 101348})

	if err != nil {
		t.Fatal(err)
	}

	if res.Title.Romaji != "Vinland Saga" {
		t.Fatal("Wrong name was returned when querying by ID.")
	}
}

//TODO: DEBUG THIS TEST
func TestQueryBIDMal(t *testing.T) {
	n, err := New()

	if err != nil {
		t.Fatal(err)
	}

	res, err :=  n.Media(MediaVariables{IDMal: 37521, Type: "ANIME"})

	if err != nil {
		t.Fatal(err)
	}

	if res.Title.Romaji != "Vinland Saga" {
		t.Fatalf("Wrong name was returned when querying by IDMal: %s", res.Title.Romaji)
	}
}

