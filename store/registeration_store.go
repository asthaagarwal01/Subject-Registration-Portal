package store

import (
    "subject-registration/models"
    "gorm.io/gorm"
)

type RegistrationStore interface {
    Create(registration *models.Registration) error
    GetAll() ([]models.Registration, error)
    Update(registration *models.Registration) error
    Delete(id uint) error
}

type registrationStore struct {
    db *gorm.DB
}

func NewRegistrationStore(db *gorm.DB) RegistrationStore {
    return &registrationStore{db: db}
}

func (store *registrationStore) Create(registration *models.Registration) error {
    return store.db.Create(registration).Error
}

func (store *registrationStore) GetAll() ([]models.Registration, error) {
    var registrations []models.Registration
    err := store.db.Find(&registrations).Error
    return registrations, err
}

func (store *registrationStore) Update(registration *models.Registration) error {
    return store.db.Save(registration).Error
}

func (store *registrationStore) Delete(id uint) error {
    return store.db.Delete(&models.Registration{}, id).Error
}
