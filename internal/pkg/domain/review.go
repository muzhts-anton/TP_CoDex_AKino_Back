package domain

type UserReview struct {
	MovieId string `json:"number"`
	Type    string `json:"type"`

	Rating string `json:"text,omitempty"`

	Date         string `json:"date,omitempty"`
	FeedbackType string `json:"feedbacktype,omitempty"`
	MovieTitle   string `json:"movieTitle,omitempty"`
}

type UserReviewResp struct {
	Id      uint64       `json:"ID"`
	Reviews []UserReview `json:"reviewsList"`
}
