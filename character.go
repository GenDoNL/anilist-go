package anilistgo

type Character struct {
	Id int `json:"id"`
	Name CharacterName `json:"name"`
	Image CharacterImage `json:"image"`
	Description string `json:"description"`
	SiteUrl string `json:"siteUrl"`
	Media Media `json:"Media"`
	Favourites int `json:"favourites"`

	// Authenticated only
	IsFavourite bool `json:"isFavourite"`
}

type CharacterConnection struct {
	Edges []CharacterEdge `json:"edges"`
	Nodes []Character `json:"nodes"`
	PageInfo PageInfo `json:"pageInfo"`
}

type CharacterEdge struct {
	Node Character `json:"node"`
	Id int `json:"id"`
	Role string `json:"role"`
	// VoiceActors []Staff `json:"voiceActors"`
	Media []Media `json:"Media"`

	// Authenticated only
	FavouriteOrder int `json:"favouriteOrder"`
}

type CharacterName struct {
	First string `json:"first"`
	Last string `json:"last"`
	Full string `json:"full"`
	Native string `json:"native"`
	Alternative []string `json:"alternative"`
}

type CharacterImage struct {
	Large string `json:"large"`
	Medium string `json:"medium"`
}

const (
	// This string is everything you can query for a Media object.
	CharacterQueryAll = ""
)

// Media returns the Media result that is found by the query and variables.
func (a *AniList) Character(q string, variables interface{}) (m Character, err error) {
	data, err := a.Data(q, variables)
	if err != nil {
		return
	}

	m = data.Character
	return
}


