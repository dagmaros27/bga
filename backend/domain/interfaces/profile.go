package interfaces

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
)

type ProfileRepository interface {
	GetUserProfile(userId string) (*entities.Profile, error)
	UpdateUserProfile(profile *entities.Profile) (*entities.Profile, error)
	CreateUserProfile(profile *entities.Profile) (*entities.Profile, error)
	DeleteUserProfile(user_id string) error
}

type ProfileService interface {
	GetUserProfile(userId string) (*entities.Profile, error)
	UpdateUserProfile(profile *dto.UpdateProfileDto) (*entities.Profile, error)
	CreateUserProfile(profile *dto.CreateProfileDto) (*entities.Profile, error)
	DeleteUserProfile(user_id string) error
}
