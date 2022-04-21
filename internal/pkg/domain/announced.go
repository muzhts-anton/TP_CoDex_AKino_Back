package domain

type Announced struct {
	Id            string `json:""`
	Poster        string `json:""`
	Title         string `json:""`
	TitleOriginal string `json:""`
	Info          string `json:""`
	Description   string `json:""`
	Trailer       string `json:""`
	Releasedate   string `json:""`
	Country       string `json:""`
	Director      string `json:""`
}
