package service

import (
	"os"
	"testing"

	"github.com/MauricioGZ/CRUD-GO/encryption"
	"github.com/MauricioGZ/CRUD-GO/internal/entity"
	"github.com/MauricioGZ/CRUD-GO/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

type MockCache struct{}

var repo *repository.MockRepository
var s Service
var mc MockCache

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validpassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	var validID int64 = 1
	var noParentID int64 = 0
	var parentIDDoesNotExist int64 = 100

	repo = &repository.MockRepository{}
	//user repo mocks
	repo.On("GetUserByEmail", mock.Anything, "email@doesntexist.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "email@existinguser.com").Return(&entity.User{
		ID:       validID,
		Email:    "email@existinguser.com",
		Password: encryptedPassword,
		RoleID:   3,
	}, nil)
	repo.On("InsertUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	//address repo mocks
	repo.On("SaveAddress", mock.Anything, validID, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetAddressesByUserId", mock.Anything, validID).Return([]entity.Address{{Address: mock.Anything}, {Address: mock.Anything}}, nil)
	//categories repo mocks
	repo.On("GetCategoryByName", mock.Anything, "New Category").Return(nil, nil)
	repo.On("GetCategoryByName", mock.Anything, "Existing Category").Return(&entity.Categories{Name: "Existing Category"}, nil)
	repo.On("InsertCategory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetCategoryByID", mock.Anything, noParentID).Return(nil, nil)
	repo.On("GetCategoryByID", mock.Anything, parentIDDoesNotExist).Return(nil, nil)
	repo.On("GetCategoryByID", mock.Anything, mock.Anything).Return(&entity.Categories{Name: "Existing Category"}, nil)

	s = New(repo)
	mc.GetAllPermissionsRoles()
	mc.GetAllRoles()

	code := m.Run()
	os.Exit(code)
}

func (mc *MockCache) GetAllPermissionsRoles() {
	rolesPermissions["Customer"] = append(rolesPermissions["Customer"], "Read")
}

func (mc *MockCache) GetAllRoles() {
	roleIDs["Customer"] = 3
}
