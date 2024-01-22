package userdb

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
)

// MySQLDB represents the MySQL database implementation.
type UserDB struct {
	DB *sql.DB
}

// CreateUser creates a new user in the MySQL database.
func (m *UserDB) CreateUser(name, email string) error {
	_, err := m.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", name, email)
	return err
}

// GetUser retrieves a user from the MySQL database by ID.
func (m *UserDB) GetUser(id uuid.UUID) (models.User, error) {
	var name, email string
	err := m.DB.QueryRow("SELECT name, email FROM users WHERE id = ?", id).Scan(&name, &email)
	if err != nil {
		return models.User{}, err
	}
	return models.User{Name: name}, nil
}

// UpdateUser updates a user in the MySQL database by ID.
func (m *UserDB) UpdateUser(id uuid.UUID, name, email string) error {
	_, err := m.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", name, email, id)
	return err
}

// DeleteUser deletes a user from the MySQL database by ID.
func (m *UserDB) DeleteUser(id uuid.UUID) error {
	_, err := m.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
