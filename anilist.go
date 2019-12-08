package anilistgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	urlAniList = "https://graphql.anilist.co"
)

type AniList struct {
}

// Query struct that contains the API query.
type Query struct {
	Query     string         `json:"query"`
	Variables interface{}    `json:"variables"`
}

// Result struct that contains the data returned by a query.
type Result struct {
	Errors []Error `json:"errors"`
	Data   Data    `json:"data"`
}

type Error struct {
	Message string `json:"message"`
}
// Data struct that wraps the data type.
type Data struct {
	Page Page `json:"Page"`

	Media Media `json:"Media"`
	User User `json:"User"`

	Character Character `json:"character"`
}

// Create a new AniList struct.
// Currently this function directly returns the struct.
// However, in the future more set-up might be required.
func New() (a *AniList, err error) {
	a = &AniList{}

	return
}

func (a *AniList) Data(q string, variables interface{}) (d Data, err error) {
	qrs := Query{q, variables}
	d, err = a.structQuery(qrs)
	return
}

func (a *AniList) structQuery(q interface{}) (d Data, err error) {
	jsonQ, err := json.Marshal(q)

	if err != nil {
		return
	}

	return a.Query(jsonQ)
}

// Query accepts a json encoding of the query, which we then return the result of.
func (a *AniList) Query(jsonData []byte) (d Data, err error) {
	resp, err := http.Post(urlAniList, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var res Result
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}

	if len(res.Errors) != 0 {
		err = errors.New(res.Errors[0].Message)
		return
	}

	d = res.Data
	return
}
