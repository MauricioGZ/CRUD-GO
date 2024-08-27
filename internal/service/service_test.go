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
	var validId int64 = 1

	repo = &repository.MockRepository{}
	// user repo mocks
	repo.On("GetUserByEmail", mock.Anything, "email@doesntexist.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "email@existinguser.com").Return(&entity.User{
		ID:       validId,
		Email:    "email@existinguser.com",
		Password: encryptedPassword,
	}, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	// address repo mocks
	repo.On("SaveAddress", mock.Anything, validId, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetAddressesByUserId", mock.Anything, validId).Return([]entity.Address{{Address: mock.Anything}, {Address: mock.Anything}}, nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}
