package services

import "github.com/penguinn/penguin-test/models"

func GetHelloWorld() (string, error) {
	str, err := models.Games{}.GetFirstName()
	if err != nil {
		return "", err
	}
	return "hello world! " + str, nil
}
