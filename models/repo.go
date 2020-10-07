package models

import "github.com/google/uuid"

// Repo is an interface which will help in mock
type Repo interface {
	Get(uuid.UUID) (Boolean, error)
	Create(Boolean) (uuid.UUID, error)
	Update(uuid.UUID, Boolean) error
	Delete(uuid.UUID) error
}

var repo Repo

// GetRepo is a function to access instance of Repo
func GetRepo() Repo {
	return repo
}

// SetRepo is a function to set repo instance from outside
func SetRepo(r Repo) {
	repo = r
}
