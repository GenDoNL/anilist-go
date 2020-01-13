package anilistgo

type Activity struct {
	Id int `json:"id"`
	UserID int `json:"userId"`

	Status string `json:"status"`
	Progress string `json:"progress"`
	SiteUrl string `json:"siteUrl"`
	CreatedAt int `json:"createdAt"`

	User User `json:"user"`
	Media Media `json:"media"`
}

func (a *AniList) Activity(q string, variables interface{}) (m Activity, err error) {
	data, err := a.Data(q, variables)
	if err != nil {
		return
	}

	m = data.Activity
	return
}