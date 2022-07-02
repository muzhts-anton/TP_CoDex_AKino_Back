package ratdelivery

type RatingResp struct {
	NewMovieRating string `json:"newrating"`
}

type RatingReq struct {
	MovieId string `json:"movieId"`
	UserId  string `json:"userId"`
	Rating  string `json:"rating"`
}
