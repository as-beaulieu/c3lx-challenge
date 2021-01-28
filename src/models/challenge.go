package models

type Challenge struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	FullName    string `json:"full_name"`
	Accepted    bool   `json:"accepted"`
}

type AcceptChallengeRequest struct {
	Name string `json:"name"`
}
