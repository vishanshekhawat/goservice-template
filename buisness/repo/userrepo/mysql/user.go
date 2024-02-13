package userdb

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
)

// MySQLDB represents the MySQL database implementation.
type UserDB struct {
	DB *sql.DB
}

// CreateUser creates a new user in the MySQL database.
func (m *UserDB) CreateUser(ctx context.Context, user models.User) (int, error) {
	res, err := m.DB.ExecContext(ctx, "INSERT INTO users (name, email,city) VALUES (?, ?,?)", user.Name, user.Email, user.City)

	if err != nil {
		return 0, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(userID), nil
}

// GetUser retrieves a user from the MySQL database by ID.
func (m *UserDB) GetUsers() ([]models.User, error) {

	rows, err := m.DB.Query("SELECT id,name, email,city FROM users")
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	var results []models.User
	for rows.Next() {
		var result models.User
		err := rows.Scan(&result.ID, &result.Name, &result.Email, &result.City)
		if err != nil {
			return results, err
		}

		results = append(results, result)
	}

	return results, nil
}
