package userdb

import (
	"database/sql"

	_ "github.com/lib/pq"
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
func (m *UserDB) GetUser(id int) (string, string, error) {
	var name, email string
	err := m.DB.QueryRow("SELECT name, email FROM users WHERE id = ?", id).Scan(&name, &email)
	if err != nil {
		return "", "", err
	}
	return name, email, nil
}

// UpdateUser updates a user in the MySQL database by ID.
func (m *UserDB) UpdateUser(id int, name, email string) error {
	_, err := m.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", name, email, id)
	return err
}

// DeleteUser deletes a user from the MySQL database by ID.
func (m *UserDB) DeleteUser(id int) error {
	_, err := m.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
