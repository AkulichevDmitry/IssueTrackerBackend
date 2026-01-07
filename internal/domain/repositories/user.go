package repositories

import (
	"issue_tracker/internal/domain/entities"
)

type UserRepository interface {
	GetByID(id entities.UserID) (*entities.User, error)
	Create(user *entities.User) error
	Update(user *entities.User) error
	Delete(id entities.UserID) error
}
