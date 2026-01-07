package usecases

import (
	"issue_tracker/internal/domain/entities"
	"issue_tracker/internal/domain/repositories"
)

type SignUpUserUC struct {
	userRepo repositories.UserRepository
}

func NewSignUpUserUC(userRepo repositories.UserRepository) *SignUpUserUC {
	return &SignUpUserUC{
		userRepo: userRepo,
	}
}

type SignUpUserInputDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpUserOutputDTO struct {
}

func (uc *SignUpUserUC) Execute(dto SignUpUserInputDTO) (*SignUpUserOutputDTO, error) {

}
