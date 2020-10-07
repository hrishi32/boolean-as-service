package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hrishi32/boolean-as-service/database"
	"gorm.io/gorm"
)

// RepoImplement is a struct for implementation of Repo interface
type RepoImplement struct{}

// Boolean is a struct to define basic structure of boolean object
type Boolean struct {
	ID    uuid.UUID `gorm:"primaryKey;column:id"`
	Value bool
	Key   string
}

// Migrate is a custom function for AutoMigration
func Migrate() {
	db, connectionError := database.GetConnection()
	if connectionError == nil {
		b := Boolean{ID: uuid.New(), Value: true, Key: "SomeKey"}
		db.AutoMigrate(&b)
	}

}

// Get receives a boolean object from database using id.
func (*RepoImplement) Get(id uuid.UUID) (Boolean, error) {
	db, connectionError := database.GetConnection()

	if connectionError != nil {
		return Boolean{}, connectionError
	}
	var boolean Boolean
	if err := db.First(&boolean, id).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		notFoundError := errors.New("Record not found")
		return Boolean{}, notFoundError
	}

	return boolean, nil
}

// Create inserts a new boolean object in the database
func (*RepoImplement) Create(b Boolean) (uuid.UUID, error) {
	id := uuid.New()
	b.ID = id
	db, connectionError := database.GetConnection()
	if connectionError != nil {
		return uuid.UUID{}, connectionError
	}
	result := db.Create(&b)

	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return id, nil
}

// Update modifies the existing boolean in the database.
func (r *RepoImplement) Update(id uuid.UUID, newBoolean Boolean) error {
	db, connectionError := database.GetConnection()

	if connectionError != nil {
		return connectionError
	}

	_, err := r.Get(id)
	if err != nil {
		return err
	}
	newBoolean.ID = id

	db.Save(&newBoolean)

	return nil

}

// Delete removes the boolean from database using id.
func (r *RepoImplement) Delete(id uuid.UUID) error {
	db, connectionError := database.GetConnection()
	if connectionError != nil {
		return connectionError
	}
	b, err := r.Get(id)

	if err != nil {
		return err
	}

	db.Delete(&b)

	return nil
}
