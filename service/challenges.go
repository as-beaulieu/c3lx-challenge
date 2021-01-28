package service

import "c3lx-challenge/src/models"

type ChallengesOperator interface {
	GetChallenges() ([]models.Challenge, error)
	AcceptChallenge(request models.AcceptChallengeRequest) (models.Challenge, error)
}

func (s service) GetChallenges() ([]models.Challenge, error) {
	return []models.Challenge{
		{
			Name:        "Not Implemented",
			Description: "Not Implemented",
		},
	}, nil
}

func (s service) AcceptChallenge(request models.AcceptChallengeRequest) (models.Challenge, error) {
	return models.Challenge{Name: "Not Implemented", Description: "Not Implemented"}, nil
}
