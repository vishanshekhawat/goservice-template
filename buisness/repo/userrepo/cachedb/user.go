package cachedb

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
)

// CacheDB represents the PostgreSQL database implementation.
type CacheDB struct {
	mu     sync.Mutex
	Users  map[int]models.User
	nextID int
}

// CreateUser creates a new user in the mock database.
func (m *CacheDB) CreateUser(name, email string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user := models.User{
		ID:    len(m.Users) + 1,
		Name:  name,
		Email: email,
	}

	m.Users[user.ID] = user
	m.nextID++

	return nil
}

// GetUser retrieves a user from the mock database by ID.
func (m *CacheDB) GetUser(id uuid.UUID) (models.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, ok := m.Users[id]
	if !ok {
		return models.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

// UpdateUser updates a user in the mock database by ID.
func (m *CacheDB) UpdateUser(id int, name, email string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, ok := m.Users[id]
	if !ok {
		return fmt.Errorf("user not found")
	}

	user.Name = name
	user.Email = email
	m.Users[id] = user

	return nil
}

// DeleteUser deletes a user from the mock database by ID.
func (m *CacheDB) DeleteUser(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.Users[id]
	if !ok {
		return fmt.Errorf("user not found")
	}

	delete(m.Users, id)
	return nil
}

// DeleteUser deletes a user from the mock database by ID.
func (m *CacheDB) GetUsers() ([]models.User, error) {
	return []models.User{}, nil
}
