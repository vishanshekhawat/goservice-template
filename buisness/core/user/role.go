package user

import "errors"

// Set of possible roles for a user.
var (
	RoleAdmin = Role{"ADMIN"}
	RoleUser  = Role{"USER"}
)

// Set of known roles.
var roles = map[string]Role{
	RoleAdmin.Name: RoleAdmin,
	RoleUser.Name:  RoleUser,
}

// Role represents a role in the system.
type Role struct {
	Name string
}

// ParseRole parses the string value and returns a role if one exists.
func ParseRole(value string) (Role, error) {
	role, exists := roles[value]
	if !exists {
		return Role{}, errors.New("invalid role")
	}

	return role, nil
}
