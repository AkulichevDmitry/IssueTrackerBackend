package repositories

import (
	"issue_tracker/internal/domain/entities"
)

type UserRepository interface {
	GetByID(id entities.UserID) (*entities.UserEntity, error)
	Create(user *entities.UserEntity) error
	Update(user *entities.UserEntity) error
	Delete(id entities.UserID) error
}
