package anilist

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	urlAniList = "https://graphql.anilist.co"
)

type AniList struct {
}

// Result struct that contains the data returned by a query.
type Result struct {
	Data Data `json:"data"`
}

// Data struct that wraps the data type.
type Data struct {
	Page Page `json:"Page"`

	Media Media `json:"Media"`
}

type Page struct {
	PageInfo PageInfo `json:"pageInfo"`

	Media []Media `json:"media"`
}

type PageInfo struct {
	Total       int  `json:"total"`
	PerPage     int  `json:"perPage"`
	CurrentPage int  `json:"currentPage"`
	LastPage    int  `json:"lastPage"`
	HasNextPage bool `json:"hasNextPage"`
}

// Create a new AniList struct.
// Currently this function directly returns the struct.
// However, in the future more set-up might be required.
func New() (a *AniList, err error) {
	a = &AniList{}

	return
}

// Query the AniList api given a MediaQuery struct.
func (a *AniList) Query(q interface{}) (body []byte, err error) {
	jsonQ, err := json.Marshal(q)

	resp, err := http.Post(urlAniList, "application/json", bytes.NewBuffer(jsonQ))

	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	return
}
