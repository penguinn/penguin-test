package models

import "github.com/penguinn/penguin/component/db"

type Games struct {
	ID      int
	Gameurl string
	Imgurl  string
	Name    string
}

func (Games) TableName() string {
	return "games"
}

func (Games) ConnectionName() string {
	return "default"
}

func (p Games) GetFirstName() (string, error) {
	var game Games

	conn, err := db.ReadModel(p)
	if err != nil {
		return "", err
	}

	err = conn.Table(p.TableName()).Select("name").First(&game).Error
	if err != nil {
		return "", err
	}
	return game.Name, nil
}
