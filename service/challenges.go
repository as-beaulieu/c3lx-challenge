package service

import (
	"c3lx-challenge/src/models"
	"go.uber.org/zap"
)

type ChallengesOperator interface {
	GetChallenges() ([]*models.Challenge, error)
	AcceptChallenge(request models.AcceptChallengeRequest) (models.Challenge, error)
}

func (s service) GetChallenges() ([]*models.Challenge, error) {
	logger := s.logger.Named("s.GetChallenges")
	logger.Info("Getting all challenges from dao")

	results, err := s.postgres.GetAllChallenges()
	if err != nil {
		logger.Error("error calling dao", zap.Error(err))
		return nil, err
	}

	logger.Info("successful return of challenges")
	return results, nil
}

func (s service) AcceptChallenge(request models.AcceptChallengeRequest) (models.Challenge, error) {
	return models.Challenge{Name: request.Name, Description: "Not Implemented"}, nil
}
