package cachedb

import (
	"fmt"
	"sync"
)

// CacheDB represents the PostgreSQL database implementation.
type CacheDB struct {
	mu     sync.Mutex
	users  map[int]User
	nextID int
}

// User represents a user entity.
type User struct {
	ID    int
	Name  string
	Email string
}

// CreateUser creates a new user in the mock database.
func (m *CacheDB) CreateUser(name, email string) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user := User{
		ID:    m.nextID,
		Name:  name,
		Email: email,
	}

	m.users[user.ID] = user
	m.nextID++

	return user.ID, nil
}

// GetUser retrieves a user from the mock database by ID.
func (m *CacheDB) GetUser(id int) (string, string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, ok := m.users[id]
	if !ok {
		return "", "", fmt.Errorf("user not found")
	}

	return user.Name, user.Email, nil
}

// UpdateUser updates a user in the mock database by ID.
func (m *CacheDB) UpdateUser(id int, name, email string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, ok := m.users[id]
	if !ok {
		return fmt.Errorf("user not found")
	}

	user.Name = name
	user.Email = email
	m.users[id] = user

	return nil
}

// DeleteUser deletes a user from the mock database by ID.
func (m *CacheDB) DeleteUser(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.users[id]
	if !ok {
		return fmt.Errorf("user not found")
	}

	delete(m.users, id)
	return nil
}
