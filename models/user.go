package models

import (
	"github.com/anupamraj16/db"
	"github.com/anupamraj16/utils"
)

type User struct {
	ID int64
	// bonding: gin advantage over native http package
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, hashedPassword)

	return err
}