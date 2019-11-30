package anilistgo

type Page struct {
	PageInfo PageInfo `json:"pageInfo"`
	Media []Media `json:"media"`
}

type PageInfo struct {
	// The total amount of results found for the given query
	Total       int  `json:"total"`
	// The amount of results shown on a single page
	PerPage     int  `json:"perPage"`
	// The current page number
	CurrentPage int  `json:"currentPage"`
	// The index of the LastPage
	LastPage    int  `json:"lastPage"`
	// Whether there is a next page.
	HasNextPage bool `json:"hasNextPage"`
}

type PageVariables struct {
	// Page number you want to access
	PageNumber int `json:"page"`

	// PerPage the amount of results shown.
	PerPage int `json:"perPage"`
}

const (
	// This string is everything you can query for a Page object.
	PageQueryAll  = "Page (page: $page, perPage: $perPage) { pageInfo {total currentPage lastPage hasNextPage perPage}  "
)

func (a *AniList) Page(q string, variables interface{}) (m Page, err error) {
	data, err := a.Data(q, variables)
	if err != nil {
		return
	}

	m = data.Page
	return
}
