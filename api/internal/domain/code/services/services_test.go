package services_test

import (
	"github.com/kodmain/thetiptop/api/internal/application/security"
	"github.com/kodmain/thetiptop/api/internal/domain/code/entities"
	"github.com/kodmain/thetiptop/api/internal/domain/code/services"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/errors"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/providers/database"
	"github.com/stretchr/testify/mock"
)

// CodeRepositoryMock est le mock pour CodeRepositoryInterface
type CodeRepositoryMock struct {
	mock.Mock
}

// Implémentation de la méthode ListErrors de CodeRepositoryInterface
func (m *CodeRepositoryMock) ListErrors() map[string]*entities.Code {
	args := m.Called()
	return args.Get(0).(map[string]*entities.Code)
}

// setup retourne un service mocké et ses dépendances
func setup() (*services.CodeService, *CodeServiceMock, *PermissionMock, *CodeRepositoryMock) {
	mockCodeService := new(CodeServiceMock)
	mockSecurity := new(PermissionMock)
	mockRepository := new(CodeRepositoryMock)
	service := services.Code(mockSecurity, mockRepository)

	return service, mockCodeService, mockSecurity, mockRepository
}

// CodeServiceMock est le mock pour CodeService
type CodeServiceMock struct {
	mock.Mock
}

func (m *CodeServiceMock) ListErrors() (map[string]*entities.Code, errors.ErrorInterface) {
	args := m.Called()
	if args.Get(0) != nil && args.Get(1) == nil {
		return args.Get(0).(map[string]*entities.Code), nil
	}
	return args.Get(0).(map[string]*entities.Code), args.Get(1).(errors.ErrorInterface)
}

// PermissionMock est le mock pour PermissionInterface
type PermissionMock struct {
	mock.Mock
}

func (m *PermissionMock) IsAuthenticated() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *PermissionMock) GetCredentialID() *string {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(*string)
}

func (m *PermissionMock) IsGrantedByRoles(roles ...security.Role) bool {
	args := m.Called(roles)
	return args.Bool(0)
}

func (m *PermissionMock) IsGrantedByRules(roles ...security.Rule) bool {
	args := m.Called(roles)
	return args.Bool(0)
}

func (m *PermissionMock) CanRead(ressource database.Entity, rules ...security.Rule) bool {
	args := m.Called(ressource)
	return args.Bool(0)
}

func (m *PermissionMock) CanCreate(ressource database.Entity, rules ...security.Rule) bool {
	args := m.Called(ressource)
	return args.Bool(0)
}

func (m *PermissionMock) CanUpdate(ressource database.Entity, rules ...security.Rule) bool {
	args := m.Called(ressource)
	return args.Bool(0)
}

func (m *PermissionMock) CanDelete(ressource database.Entity, rules ...security.Rule) bool {
	args := m.Called(ressource)
	return args.Bool(0)
}
