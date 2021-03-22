package main

import (
	"errors"
)

func validateScore(score int) (int, error) {

	if score < 0 {
		return -1, errors.New("Score should be greather than zero")
	}

	return score, nil
}

func validateSubject(subjectName string) (string, error) {

	if len(subjectName) <= 0 {
		return "", errors.New("Subject name should have at least one character")
	}

	return subjectName, nil
}
