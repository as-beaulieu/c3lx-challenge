package service

import "c3lx-challenge/src/models"

type ChallengesOperator interface {
	GetChallenges() ([]models.Challenge, error)
	AcceptChallenge(request models.AcceptChallengeRequest) (models.Challenge, error)
}
