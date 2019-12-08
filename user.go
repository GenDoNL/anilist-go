package anilistgo

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	About string `json:"about"`
	Avatar UserAvatar `json:"avatar"`
	BannerImage string `json:"bannerImage"`
	Favourites Favourites `json:"favourites"`
	Statistics UserStatisticsType `json:"statistics"`
	SiteUrl string `json:"siteUrl"`

	// Authenticated only
	IsFollowing bool `json:"isFollowing"`
	IsFollower bool `json:"isFollower"`
	IsBlocked bool `json:"isBlocked"`
}

type UserAvatar struct {
	Large string `json:"large"`
	Medium string `json:"medium"`
}

type Favourites struct {
	Anime MediaConnection `json:"anime"`
	Manga MediaConnection `json:"manga"`
	Characters CharacterConnection `json:"characters"`
	// TODO: Staff & Studios
}

type UserStatisticsType struct {
	Anime UserStatistics `json:"anime"`
	Manga UserStatistics `json:"manga"`
}

type UserStatistics struct {
	Count int `json:"count"`
	MeanScore float64 `json:"meanScore"`
	StandardDeviation float64 `json:"standardDeviation"`
	MinutesWatched int `json:"minutesWatched"`
	EpisodesWatched int `json:"episodesWatched"`
	ChaptersRead int `json:"chaptersRead"`
	VolumesRead int `json:"volumesRead"`
}

func (a *AniList) User(q string, variables interface{}) (m User, err error) {
	data, err := a.Data(q, variables)
	if err != nil {
		return
	}

	m = data.User
	return
}