package models

type Config struct {
	Type         string
	User         string
	Password     string
	HostPort     string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
}
