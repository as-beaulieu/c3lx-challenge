package service

import (
	"c3lx-challenge/src/models"
	"errors"
	"go.uber.org/zap"
)

type ChallengesOperator interface {
	GetChallenges() ([]*models.Challenge, error)
	AcceptChallenge(request models.AcceptChallengeRequest) (*models.Challenge, error)
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

func (s service) AcceptChallenge(request models.AcceptChallengeRequest) (*models.Challenge, error) {
	logger := s.logger.Named("s.GetChallenges").With(zap.String("name", request.Name))
	logger.Info("Getting challenge from dao by name")

	challenge, err := s.postgres.GetChallengeByName(request.Name)
	if err != nil {
		logger.Error("error calling dao for get of challenge by name", zap.Error(err))
		return nil, err
	}

	if challenge == nil {
		logger.Error("no challenge found by request name")
		return nil, errors.New("does not exist")
	}

	logger.Info("successful get of challenge by name")
	return challenge, nil
}
