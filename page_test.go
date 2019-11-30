package anilistgo

import "testing"

func TestPageFullStack(t *testing.T) {
	pageSize := 3

	query := "query ($page: Int, $perPage: Int, $search: String, $type: MediaType) " +
		"{ Page (page: $page, perPage: $perPage) { pageInfo { total currentPage lastPage hasNextPage perPage }" +
		"media (search: $search, type: $type) { id idMal title { romaji english native } type genres} } }"

	variables := struct {
		Search string `json:"search"`
		Type string `json:"type"`
		Page int `json:"page"`
		PerPage int `json:"perPage"`
	}{
		"Vinland",
		"ANIME",
		1,
		pageSize,
	}

	a, err := New()
	if err != nil {
		t.Fatal(err)
	}

	p, err := a.Page(query, variables)
	if err != nil {
		t.Fatal(err)
	}
	m := p.Media[0]

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

	if len(p.Media) != pageSize {
		t.Fatalf("Page size does not match, was: %d, expected: %d", len(p.Media), pageSize)
	}

	if p.PageInfo.CurrentPage != 1 {
		t.Fatalf("Wrong page was returned: %d", p.PageInfo.CurrentPage)
	}
}