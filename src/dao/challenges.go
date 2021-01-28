package dao

import (
	"c3lx-challenge/src/models"
	"fmt"
	"log"
)

type ChallengeStore interface {
	GetAllChallenges() ([]*models.Challenge, error)
	GetChallengeByName(name string) (*models.Challenge, error)
}

func (d dao) GetAllChallenges() ([]*models.Challenge, error) {
	results := make([]*models.Challenge, 0)
	queryString := `SELECT * FROM ablue.challenges;`

	rows, err := d.connection.Query(queryString)
	if err != nil {
		return nil, fmt.Errorf(
			"error with select all statement of challenges in database: %+v", err)
	}

	for rows.Next() {
		var challenge models.Challenge

		if err := rows.Scan(
			&challenge.Name, &challenge.Description, &challenge.FullName); err != nil {
			log.Println(err)
			return nil, err
		}

		results = append(results, &challenge)
	}

	return results, nil
}

func (d dao) GetChallengeByName(name string) (*models.Challenge, error) {
	var challenge models.Challenge

	queryString := `SELECT * FROM ablue.challenges WHERE name = $1`

	row, err := d.connection.Query(queryString, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for row.Next() {
		if err := row.Scan(
			&challenge.Name, &challenge.Description, &challenge.FullName); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &challenge, nil
}
