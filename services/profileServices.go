package services

import (
	"errors"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
	"gorm.io/gorm"
)

type ProfileService interface {
	GetProfile(userID int) (*entities.Profil, error)
	UpdateProfile(userID int, updatedProfile *entities.Profil) error
}

type profileService struct {
	repo repositories.ProfileRepository
}

func NewProfileService(repo repositories.ProfileRepository) ProfileService {
	return &profileService{repo}
}

func (s *profileService) GetProfile(userID int) (*entities.Profil, error) {
	var user entities.User
	if err := s.repo.GetDB().First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	profile, err := s.repo.GetProfileByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var id = 0
			id += 1
			return &entities.Profil{
				ID:     id,
				Avatar: "https://cu-sehati.com/wp-content/uploads/2020/04/avatar.png.jpg",
				UserID: userID,
				User:   user,
			}, nil
		}
		return nil, err
	}

	return profile, nil
}

func (s *profileService) UpdateProfile(userID int, updatedProfile *entities.Profil) error {
	profile, err := s.repo.GetProfileByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("profile not found")
		}
		return err
	}
	if updatedProfile.Avatar != "" {
		profile.Avatar = updatedProfile.Avatar
	}
	profile.User.NamaLengkap = updatedProfile.User.NamaLengkap
	profile.User.Alamat = updatedProfile.User.Alamat
	profile.User.NoTelpon = updatedProfile.User.NoTelpon
	return s.repo.UpdateProfile(profile)
}
