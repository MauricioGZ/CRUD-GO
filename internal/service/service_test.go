package service

import (
	"os"
	"testing"

	"github.com/MauricioGZ/CRUD-GO/encryption"
	"github.com/MauricioGZ/CRUD-GO/internal/entity"
	"github.com/MauricioGZ/CRUD-GO/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validpassword"))
	encryptedPassword := encryption.ToBase64(validPassword)

	repo = &repository.MockRepository{}
	repo.On("GetUserByEmail", mock.Anything, "email@newuser.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "email@existinguser.com").Return(&entity.User{
		Email:    "email@existinguser.com",
		Password: encryptedPassword,
	}, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}
