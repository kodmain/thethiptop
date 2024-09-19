package services_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/gofiber/fiber/v2"
	"github.com/kodmain/thetiptop/api/internal/application/services"
	"github.com/kodmain/thetiptop/api/internal/application/transfert"
	"github.com/kodmain/thetiptop/api/internal/domain/user/entities"
	"github.com/kodmain/thetiptop/api/internal/domain/user/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterEmployee(t *testing.T) {
	t.Run("invalid password", func(t *testing.T) {
		mockService := new(DomainUserService)
		mockService.On("RegisterEmployee", mock.AnythingOfType("*transfert.Credential"), mock.AnythingOfType("*transfert.Employee")).Return(&entities.Employee{}, nil)

		statusCode, response := services.RegisterEmployee(mockService, &transfert.Credential{
			Email:    &email,
			Password: &passwordSyntaxFail,
		}, &transfert.Employee{})
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
		assert.NotNil(t, response)
	})

	t.Run("valid password and fields", func(t *testing.T) {
		mockService := new(DomainUserService)
		mockService.On("RegisterEmployee", mock.AnythingOfType("*transfert.Credential"), mock.AnythingOfType("*transfert.Employee")).Return(&entities.Employee{}, nil)

		statusCode, response := services.RegisterEmployee(mockService, &transfert.Credential{
			Email:    &email,
			Password: &password,
		}, &transfert.Employee{})

		assert.Equal(t, fiber.StatusCreated, statusCode)
		assert.NotNil(t, response)
	})

	t.Run("employee already exists", func(t *testing.T) {
		mockService := new(DomainUserService)
		mockService.On("RegisterEmployee", mock.AnythingOfType("*transfert.Credential"), mock.AnythingOfType("*transfert.Employee")).Return(nil, fmt.Errorf(errors.ErrCredentialAlreadyExists))

		statusCode, response := services.RegisterEmployee(mockService, &transfert.Credential{
			Email:    &email,
			Password: &password,
		}, &transfert.Employee{})
		assert.Equal(t, fiber.StatusConflict, statusCode)
		assert.NotNil(t, response)
	})

	t.Run("server error during registration", func(t *testing.T) {
		mockService := new(DomainUserService)
		mockService.On("RegisterEmployee", mock.AnythingOfType("*transfert.Credential"), mock.AnythingOfType("*transfert.Employee")).Return(nil, fmt.Errorf("server error"))

		statusCode, response := services.RegisterEmployee(mockService, &transfert.Credential{
			Email:    &email,
			Password: &password,
		}, &transfert.Employee{})
		assert.Equal(t, fiber.StatusInternalServerError, statusCode)
		assert.Equal(t, "server error", response)
	})
}

func TestDeleteEmployee(t *testing.T) {
	t.Run("should return 400 if validation fails", func(t *testing.T) {
		mockService := new(DomainUserService)
		dtoEmployee := &transfert.Employee{ID: nil}

		statusCode, response := services.DeleteEmployee(mockService, dtoEmployee)

		assert.Equal(t, fiber.StatusBadRequest, statusCode)
		assert.Equal(t, "value id is required", response)
	})

	t.Run("should return 404 if employee not found", func(t *testing.T) {
		mockService := new(DomainUserService)
		employeeID := "123e4567-e89b-12d3-a456-426614174000"
		dtoEmployee := &transfert.Employee{ID: &employeeID}

		mockService.On("DeleteEmployee", dtoEmployee).Return(fmt.Errorf(errors.ErrEmployeeNotFound))

		statusCode, response := services.DeleteEmployee(mockService, dtoEmployee)

		assert.Equal(t, fiber.StatusNotFound, statusCode)
		assert.Equal(t, errors.ErrEmployeeNotFound, response)
		mockService.AssertExpectations(t)
	})

	t.Run("should return 500 if internal server error occurs", func(t *testing.T) {
		mockService := new(DomainUserService)
		employeeID := "123e4567-e89b-12d3-a456-426614174000"
		dtoEmployee := &transfert.Employee{ID: &employeeID}

		mockService.On("DeleteEmployee", dtoEmployee).Return(fmt.Errorf("internal error"))

		statusCode, response := services.DeleteEmployee(mockService, dtoEmployee)

		assert.Equal(t, fiber.StatusInternalServerError, statusCode)
		assert.Equal(t, "internal error", response)
		mockService.AssertExpectations(t)
	})

	t.Run("should return 204 if employee is deleted successfully", func(t *testing.T) {
		mockService := new(DomainUserService)
		employeeID := "123e4567-e89b-12d3-a456-426614174000"
		dtoEmployee := &transfert.Employee{ID: &employeeID}

		mockService.On("DeleteEmployee", dtoEmployee).Return(nil)

		statusCode, response := services.DeleteEmployee(mockService, dtoEmployee)

		assert.Equal(t, fiber.StatusNoContent, statusCode)
		assert.Nil(t, response)
		mockService.AssertExpectations(t)
	})
}

func TestUpdateEmployee(t *testing.T) {
	t.Run("invalid employee data", func(t *testing.T) {
		mockService := new(DomainUserService)
		statusCode, response := services.UpdateEmployee(mockService, &transfert.Employee{
			ID: nil, // ID manquant
		})
		assert.Equal(t, fiber.StatusBadRequest, statusCode)
		assert.Equal(t, "value id is required", response)
	})

	t.Run("successful employee update", func(t *testing.T) {
		mockService := new(DomainUserService)
		mockService.On("UpdateEmployee", mock.AnythingOfType("*transfert.Employee")).Return(nil, nil)

		statusCode, response := services.UpdateEmployee(mockService, &transfert.Employee{
			ID: aws.String("123e4567-e89b-12d3-a456-426614174000"),
		})

		assert.Equal(t, fiber.StatusOK, statusCode)
		assert.Nil(t, response)
		mockService.AssertExpectations(t)
	})

	t.Run("employee update error", func(t *testing.T) {
		mockService := new(DomainUserService)
		mockService.On("UpdateEmployee", mock.AnythingOfType("*transfert.Employee")).Return(nil, fmt.Errorf("update error"))

		statusCode, response := services.UpdateEmployee(mockService, &transfert.Employee{
			ID: aws.String("123e4567-e89b-12d3-a456-426614174000"),
		})
		assert.Equal(t, fiber.StatusInternalServerError, statusCode)
		assert.Equal(t, "update error", response)
		mockService.AssertExpectations(t)
	})
}

func TestGetEmployee(t *testing.T) {
	t.Run("validation error", func(t *testing.T) {
		mockService := new(DomainUserService)
		dtoEmployee := &transfert.Employee{ID: nil}

		statusCode, response := services.GetEmployee(mockService, dtoEmployee)

		assert.Equal(t, fiber.StatusBadRequest, statusCode)
		assert.Equal(t, "value id is required", response)
	})

	t.Run("employee not found", func(t *testing.T) {
		mockService := new(DomainUserService)
		dtoEmployee := &transfert.Employee{
			ID: aws.String("42debee6-2063-4566-baf1-37a7bdd139ff"),
		}

		mockService.On("GetEmployee", dtoEmployee).Return(nil, fmt.Errorf(errors.ErrEmployeeNotFound))

		statusCode, response := services.GetEmployee(mockService, dtoEmployee)

		assert.Equal(t, fiber.StatusNotFound, statusCode)
		assert.Equal(t, errors.ErrEmployeeNotFound, response)
		mockService.AssertExpectations(t)
	})

	t.Run("employee random error", func(t *testing.T) {
		mockService := new(DomainUserService)
		dtoEmployee := &transfert.Employee{
			ID: aws.String("42debee6-2063-4566-baf1-37a7bdd139ff"),
		}

		mockService.On("GetEmployee", dtoEmployee).Return(nil, fmt.Errorf("random error"))

		statusCode, response := services.GetEmployee(mockService, dtoEmployee)

		assert.Equal(t, fiber.StatusInternalServerError, statusCode)
		assert.Equal(t, "random error", response)
		mockService.AssertExpectations(t)
	})

	t.Run("successful employee retrieval", func(t *testing.T) {
		mockService := new(DomainUserService)
		dtoEmployee := &transfert.Employee{
			ID: aws.String("42debee6-2063-4566-baf1-37a7bdd139ff"),
		}

		expectedEmployee := &entities.Employee{
			ID: "42debee6-2063-4566-baf1-37a7bdd139ff",
		}

		mockService.On("GetEmployee", dtoEmployee).Return(expectedEmployee, nil)

		statusCode, response := services.GetEmployee(mockService, dtoEmployee)

		assert.Equal(t, fiber.StatusOK, statusCode)
		assert.Equal(t, expectedEmployee, response)
		mockService.AssertExpectations(t)
	})
}