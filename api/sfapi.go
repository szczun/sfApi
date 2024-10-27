package api

type SFRes struct {
    Id string `json:"id"`
	Name string `json:"name"`
	Matches string `json:"Matches"`
	Wins string `json:"wins"`
	Losses string `json:"losses"`
	Draws string `json:"draws"`
	GoalBalance string `json:"goalBalance"`
	LastFive string `json:"lastFive"`
	Points string `json:"points"`
    PhotoUrl string `json:"photoUrl"`
}
