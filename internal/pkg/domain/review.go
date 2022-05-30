package domain

type UserReview struct {
	Type         string `json:"type"` // 1 - rating, 2 - comment, 3 - rating + comment
	Rating       string `json:"rating"`
	Date         string `json:"date,omitempty"`
	MovieId      string `json:"movieID"`
	MovieTitle   string `json:"movieTitle,omitempty"`
	MoviePoster  string `json:"moviePoster,omitempty"`
	Text         string `json:"text,omitempty"`
	FeedbackType string `json:"feedbacktype,omitempty"`
}

type UserReviewResp struct {
	Id      uint64       `json:"ID"`
	Reviews []UserReview `json:"reviewsList"`
}
