package anilistgo

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

	Relations MediaConnection `json:"relations"`
	// Character and staff to be added ?

	// The media's next episode airing schedule
	// Can be null if it has finished airing or no date is known.
	NextAiringEpisode AiringSchedule `json:"nextAiringEpisode"`

	SiteUrl string `json:"siteUrl"`
}

type MediaConnection struct {
	Edges []MediaEdge `json:"edges"`
	Nodes []Media `json:"nodes"`
	PageInfo PageInfo `json:"pageInfo"`
}

type MediaEdge struct {
	Node Media `json:"node"`
	Id int `json:"id"`
	Relation string `json:"relationType"`
	IsMainStudio bool `json:"isMainStudio"`
	Characters []Character `json:"characters"`
	CharacterRole string `json:"characterRole"`
	StaffRole string `json:"staffRole"`
	// VoiceActors []Staff `json:"voiceActors"`

	// Authenticated only
	FavouriteOrder int `json:"favouriteOrder"`
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
	// This string is everything you can query for a Media object.
	MediaQueryAll = "{ id idMal title { romaji english native } " +
		"type format status description startDate { year month day } endDate { year month day } " +
		"season seasonInt episodes duration chapters volumes countryOfOrigin isLicensed " +
		"source updatedAt coverImage {extraLarge large medium color} bannerImage genres" +
		"synonyms averageScore meanScore popularity favourites " +
		"nextAiringEpisode { timeUntilAiring airingAt } siteUrl } "
)

// Media returns the Media result that is found by the query and variables.
func (a *AniList) Media(q string, variables interface{}) (m Media, err error) {
	data, err := a.Data(q, variables)
	if err != nil {
		return
	}

	m = data.Media
	return
}


