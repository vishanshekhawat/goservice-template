package cachedb

import (
	"context"
	"sync"

	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
)

// CacheDB represents the PostgreSQL database implementation.
type CacheDB struct {
	mu     sync.Mutex
	Users  map[int]models.User
	nextID int
}

// CreateUser creates a new user in the mock database.
func (m *CacheDB) CreateUser(ctx context.Context, user models.User) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	newUser := models.User{
		ID:    len(m.Users) + 1,
		Name:  user.Name,
		Email: user.Email,
		City:  user.City,
	}

	m.Users[newUser.ID] = newUser
	m.nextID++

	return user.ID, nil
}

// DeleteUser deletes a user from the mock database by ID.
func (m *CacheDB) GetUsers() ([]models.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var users []models.User
	for _, val := range m.Users {
		users = append(users, val)
	}

	return users, nil
}
