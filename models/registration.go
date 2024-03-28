package models

import (
	"github.com/anupamraj16/db"
)

type Registration struct {
	ID      int64
	EventID int64 `binding:"required"`
	UserID  int64 `binding:"required"`
}

func GetAllRegistrations() ([]Registration, error) {
	query := "SELECT * FROM registrations"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registrations []Registration
	// Next returns true as long as rows are there
	for rows.Next() {
		var r Registration
		err := rows.Scan(&r.ID, &r.EventID, &r.UserID)

		if err != nil {
			return nil, err
		}
		registrations = append(registrations, r)
	}

	return registrations, nil
}
