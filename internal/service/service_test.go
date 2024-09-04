package service

import (
	"os"
	"testing"

	"github.com/MauricioGZ/CRUD-GO/encryption"
	"github.com/MauricioGZ/CRUD-GO/internal/entity"
	"github.com/MauricioGZ/CRUD-GO/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

type mockCache struct{}

var repo *repository.MockRepository
var s Service
var mc mockCache

const (
	validUserID                       int64  = 1
	validUserIDWithoutOrders          int64  = 2
	validUserIDWithoutAddresses       int64  = 3
	validAddresID                     int64  = 1
	noParentID                        int64  = 0
	invalidCategoryID                 int64  = 100
	validCategoryID                   int64  = 1
	validProductID                    int64  = 1
	invalidProductID                  int64  = 100
	validOrderItemID                  int64  = 1
	existingUserEmail                 string = "email@existinguser.com"
	unexistingUserEmail               string = "email@unexistinguser.com"
	existingUserEmailWithoutOrders    string = "noordersemail@existinguser.com"
	existingUserEmailWithoutAddresses string = "noaddressesemail@existinguser.com"
	validPassword                     string = "validpassword"
	invalidPassword                   string = "invalidpassword"
	newCategory                       string = "New Category"
	existingCategory                  string = "Existing Category"
	unexistingCategory                string = "Unexisting Category"
)

var (
	validOrderID   int64 = 1
	invalidOrderID int64 = 100
)

func TestMain(m *testing.M) {
	encryptedValidPassword, _ := encryption.Encrypt([]byte(validPassword))
	encryptedValidPassword64 := encryption.ToBase64(encryptedValidPassword)

	//cache mocks
	mc.GetAllPermissionsRoles()
	mc.GetAllRoles()
	repo = &repository.MockRepository{}
	//user repo mocks
	repo.On("GetUserByEmail", mock.Anything, unexistingUserEmail).Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, existingUserEmail).Return(&entity.User{
		ID:       validUserID,
		Email:    existingUserEmail,
		Password: encryptedValidPassword64,
		RoleID:   roleIDs["Customer"],
	}, nil)
	repo.On("GetUserByEmail", mock.Anything, existingUserEmailWithoutOrders).Return(&entity.User{
		ID:       validUserIDWithoutOrders,
		Email:    existingUserEmailWithoutOrders,
		Password: encryptedValidPassword64,
		RoleID:   roleIDs["Customer"],
	}, nil)
	repo.On("GetUserByEmail", mock.Anything, existingUserEmailWithoutAddresses).Return(&entity.User{
		ID:       validUserIDWithoutAddresses,
		Email:    existingUserEmailWithoutOrders,
		Password: encryptedValidPassword64,
		RoleID:   roleIDs["Customer"],
	}, nil)
	repo.On("InsertUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	//address repo mocks
	repo.On("SaveAddress", mock.Anything, validUserID, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetAddressesByUserId", mock.Anything, validUserID).Return([]entity.Address{{Address: mock.Anything}, {Address: mock.Anything}}, nil)
	repo.On("GetAddressesByUserId", mock.Anything, validUserIDWithoutAddresses).Return(nil, nil)
	repo.On("GetAddressByID", mock.Anything, validAddresID, validUserID).Return(&entity.Address{ID: validAddresID}, nil)
	repo.On("GetAddressByID", mock.Anything, mock.Anything, validUserIDWithoutAddresses).Return(nil, nil)
	repo.On("UpdateAddressByID", mock.Anything, validAddresID, validUserID, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("DeleteAddressByID", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	//categories repo mocks
	repo.On("GetCategoryByName", mock.Anything, newCategory).Return(nil, nil)
	repo.On("GetCategoryByName", mock.Anything, unexistingCategory).Return(nil, nil)
	repo.On("GetCategoryByName", mock.Anything, existingCategory).Return(&entity.Categories{Name: existingCategory}, nil)
	repo.On("InsertCategory", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetCategoryByID", mock.Anything, noParentID).Return(nil, nil)
	repo.On("GetCategoryByID", mock.Anything, invalidCategoryID).Return(nil, nil)
	repo.On("GetCategoryByID", mock.Anything, validCategoryID).Return(&entity.Categories{Name: existingCategory}, nil)
	//products repo mocks
	repo.On("InsertProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, validCategoryID, mock.Anything).Return(nil)
	repo.On("GetProductByID", mock.Anything, validProductID).Return(&entity.Product{ID: validProductID}, nil)
	repo.On("GetProductByID", mock.Anything, invalidProductID).Return(nil, nil)
	repo.On("GetProductsByCategoryID", mock.Anything, mock.Anything).Return([]entity.Product{{ID: validProductID}}, nil)
	repo.On("UpdateProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("DeleteProductByID", mock.Anything, mock.Anything).Return(nil)
	//orders repo mocks
	repo.On("InsertOrder", mock.Anything, validUserID, mock.Anything, mock.Anything).Return(&validOrderID, nil)
	repo.On("GetOrderByID", mock.Anything, validOrderID).Return(&entity.Order{ID: validOrderID}, nil)
	repo.On("GetOrderByID", mock.Anything, invalidOrderID).Return(nil, nil)
	repo.On("GetOrdersByUserID", mock.Anything, validUserID).Return([]entity.Order{{ID: validProductID}}, nil)
	repo.On("GetOrdersByUserID", mock.Anything, validUserIDWithoutOrders).Return(nil, nil)
	//order items mocks
	repo.On("InsertOrderItem", mock.Anything, validOrderID, validProductID, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetOrderItemsByUserID", mock.Anything, validUserID).Return([]entity.OrderItemByUserID{{OrderItem: entity.OrderItem{ID: validOrderItemID}}}, nil)
	repo.On("GetOrderItemsByUserID", mock.Anything, validUserIDWithoutOrders).Return(nil, nil)

	s = New(repo)

	code := m.Run()
	os.Exit(code)
}

func (mc *mockCache) GetAllPermissionsRoles() {
	rolesPermissions["Customer"] = append(rolesPermissions["Customer"], "Read")
	rolesPermissions["Seller"] = append(rolesPermissions["Seller"], "Create", "Update", "Read")
	rolesPermissions["Admin"] = append(rolesPermissions["Admin"], "Create", "Update", "Read", "Delete")
}

func (mc *mockCache) GetAllRoles() {
	roleIDs["Customer"] = 3
	roleIDs["Seller"] = 2
	roleIDs["Admin"] = 1
}
