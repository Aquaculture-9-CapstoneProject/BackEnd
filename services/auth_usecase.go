package services

import (
	"errors"

	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/entities"
	"github.com/Aquaculture-9-CapstoneProject/BackEnd.git/repositories"
)

type AuthUseCase interface {
	DaftarUser(namaLengkap, alamat, noTelpon, email, password, konfirmasiPass string) (*entities.User, error)
	LoginUser(email, password string) (*entities.User, error)
	LoginAdmin(email, password string) (*entities.Admin, error)
}

type authUseCase struct {
	repo repositories.RepoAuth
}

func NewAuthUseCase(repo repositories.RepoAuth) AuthUseCase {
	return &authUseCase{repo: repo}
}

func (uc *authUseCase) DaftarUser(namaLengkap, alamat, noTelpon, email, password, konfirmasiPass string) (*entities.User, error) {
	if password != konfirmasiPass {
		return nil, errors.New("password dan Confirmasi Password Tidak sesuai")
	}

	searchEmail, err := uc.repo.UserEmail(email)
	if err == nil && searchEmail != nil {
		return nil, errors.New("email sudah digunakan")
	}

	user := &entities.User{
		NamaLengkap:    namaLengkap,
		Alamat:         alamat,
		NoTelpon:       noTelpon,
		Email:          email,
		Password:       password,
		KonfirmasiPass: konfirmasiPass,
	}

	if err := uc.repo.DaftarAuth(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *authUseCase) LoginUser(email, password string) (*entities.User, error) {
	user, err := uc.repo.UserEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, errors.New("password salah")
	}
	return user, nil
}

func (uc *authUseCase) LoginAdmin(email, password string) (*entities.Admin, error) {
	admin, err := uc.repo.AdminEmail(email)
	if err != nil {
		return nil, err
	}
	if admin.Password != password {
		return nil, errors.New("password salah")
	}
	return admin, nil
}
