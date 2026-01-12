package repositories

import (
	"issue_tracker/internal/domain/entities/user"
)

type UserRepository interface {
	GetByID(id user.ID) (*user.User, error)
	Create(user user.User) error
	Update(user user.User) error
	Delete(id user.ID) error
}
