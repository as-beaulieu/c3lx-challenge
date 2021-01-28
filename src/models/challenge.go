package models

type Challenge struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AcceptChallengeRequest struct {
	Name string `json:"name"`
}
