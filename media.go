package anilist

import (
	"encoding/json"
	"errors"
	"fmt"
)

// MediaQuery struct that contains the API query.
type MediaQuery struct {
	Query     string         `json:"query"`
	Variables MediaVariables `json:"variables"`
}

// MediaVariables for the query type.
type MediaVariables struct {
	// Page number you want to acces
	Page int `json:"page"`

	// PerPage the amount of results shown.
	PerPage int `json:"perPage"`

	// The ID as defined by AniList
	ID int `json:"id"`

	// The ID as found on MyAnimeList
	IDMal int `json:"idMal"`

	// The SearchQuery to find anime as used by the search bar on AniList
	// Use this to search for anime by name.
	SearchQuery string `json:"search"`

	// Specify the type of the Anime ("MANGA"/"ANIME")
	Type string `json:"type"`
}

// Media is a struct that will contain the data for a Media object.
// In general, this is either a Manga or an Anime.
type Media struct {
	Id              int        `json:"id"`
	IdMal           int        `json:"idMal"`
	Title           MediaTitle `json:"title"`
	Type            string     `json:"type"`
	Format          string     `json:"mediaFormat"`
	Status          string     `json:"mediaStatus"`
	Description     string     `json:"description"`
	StartDate       FuzzyDate  `json:"startDate"`
	EndDate         FuzzyDate  `json:"endDate"`
	Season          string     `json:"mediaSeason"`
	SeasonInt       int        `json:"seasonInt"`
	Episodes        int        `json:"episodes"`
	Duration        int        `json:"duration"`
	Chapters        int        `json:"chapters"`
	Volumes         int        `json:"volumes"`
	CountryOfOrigin string     `json:"countryOfOrigin"`
	IsLicensed      bool       `json:"isLicensed"`
	Source          string     `json:"source"`

	UpdatedAt   int        `json:"updatedAt"`
	CoverImage  CoverImage `json:"coverImage"`
	BannerImage string     `json:"bannerImage"`
	Genres      []string   `json:"genres"`
	Synonyms    []string   `json:"synonyms"`

	AverageScore int `json:"averageScore"`
	MeanScore    int `json:"meanScore"`
	Popularity   int `json:"popularity"`

	Favourites int `json:"favourites"`

	// Character and staff to be added ?

	// The media's next episode airing schedule
	// Can be null if it has finished airing or no date is known.
	NextAiringEpisode AiringSchedule `json:"nextAiringEpisode"`

	SiteUrl string `json:"siteUrl"`
}

type CoverImage struct {
	// The cover image url of the media at its largest size.
	// If unavailable, Large will be provided instead.
	ExtraLarge string `json:"extraLarge"`

	// The cover image url of the media at a large size.
	Large string `json:"large"`

	// The cover image url of the media at medium size.
	Medium string `json:"medium"`

	// Average #hex color of cover image
	Color string `json:"color"`
}

type MediaTitle struct {
	Romaji  string `json:"romaji"`
	English string `json:"english"`
	Native  string `json:"native"`
}

type FuzzyDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type AiringSchedule struct {
	TimeUntilAiring int `json:"timeUntilAiring"`
	AiringAt        int `json:"airingAt"`
}

const (
	pageFmtString  = "query (%s) {Page (page: $page, perPage: $perPage) { pageInfo {total currentPage lastPage hasNextPage perPage} %s }  }"
	mediaFmtString = "media (%s) { id idMal title { romaji english native } " +
		"type format status description startDate { year month day } endDate { year month day } " +
		"season seasonInt episodes duration chapters volumes countryOfOrigin isLicensed " +
		"source updatedAt coverImage {extraLarge large medium color} bannerImage genres" +
		" synonyms averageScore meanScore popularity favourites " +
		"nextAiringEpisode { timeUntilAiring airingAt } siteUrl } "
)

// Media returns the first result that is found by the query given the variables.
func (a *AniList) Media(s MediaVariables) (m Media, err error) {
	mPage, err := a.MediaPage(s)
	if err != nil {
		return
	}
	m = mPage.Media[0]
	return
}

// MediaPage retuns a page of the results found by the query given the variables.
func (a *AniList) MediaPage(s MediaVariables) (p Page, err error) {
	query, err := createMediaQuery(s)

	res, err := a.Query(query)

	if err != nil {
		return
	}

	var rest Result
	err = json.Unmarshal(res, &rest)

	p = rest.Data.Page
	return
}

func createMediaQuery(s MediaVariables) (q MediaQuery, err error) {
	var initVars string
	var vars string

	// Set-up variables for page query
	page := s.Page
	if s.Page == 0 {
		page = 1
	}

	perPage := s.PerPage
	if s.PerPage == 0 {
		perPage = 1
	}
	initVars = "$page: Int, $perPage: Int,"

	// Set-up values for media query
	if s.SearchQuery != "" {
		initVars += "$search: String,"
		vars += "search: $search,"
	}

	if s.Type != "" {
		initVars += "$type: MediaType,"
		vars += "type: $type,"
	}

	if s.ID != 0 {
		initVars += "$id: Int,"
		vars += "id: $id,"
	}

	if s.IDMal != 0 {
		initVars += "$idmal: Int,"
		vars += "idMal: $idmal,"
	}

	if initVars == "" {
		return q, errors.New("cannot query with empty search")
	}

	mediaQuery := fmt.Sprintf(mediaFmtString, vars)
	variables := MediaVariables{
		Page:    page,
		PerPage: perPage,

		SearchQuery: s.SearchQuery,
		Type:        s.Type,
		ID:          s.ID,
		IDMal:       s.IDMal,
	}

	query := fmt.Sprintf(pageFmtString, initVars, mediaQuery)

	q = MediaQuery{query, variables}
	return
}
