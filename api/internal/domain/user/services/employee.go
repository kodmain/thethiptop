package services

import (
	"github.com/kodmain/thetiptop/api/internal/application/transfert"
	"github.com/kodmain/thetiptop/api/internal/domain/user/entities"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/data"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/errors"
)

func (s *UserService) RegisterEmployee(dtoCredential *transfert.Credential, dtoEmployee *transfert.Employee) (*entities.Employee, errors.ErrorInterface) {
	if dtoCredential == nil || dtoEmployee == nil {
		return nil, errors.ErrNoDto
	}

	_, err := s.repo.ReadCredential(dtoCredential)
	if err == nil {
		return nil, errors.ErrEmployeeAlreadyExists
	}

	credential, err := s.repo.CreateCredential(dtoCredential)
	if err != nil {
		return nil, errors.FromErr(err, errors.ErrInternalServer)
	}

	dtoEmployee.CredentialID = &credential.ID

	employee, err := s.repo.CreateEmployee(dtoEmployee)
	if err != nil {
		return nil, errors.FromErr(err, errors.ErrInternalServer)
	}

	employee.Validations = append(employee.Validations, &entities.Validation{
		EmployeeID: &employee.ID,
		Type:       entities.MailValidation,
	})

	if err := s.repo.UpdateEmployee(employee); err != nil {
		return nil, errors.FromErr(err, errors.ErrInternalServer)
	}

	if err := s.repo.UpdateCredential(credential); err != nil {
		return nil, errors.FromErr(err, errors.ErrInternalServer)
	}

	go s.sendValidationMail(credential, employee.Validations[0])

	return employee, nil
}

func (s *UserService) UpdateEmployee(dtoEmployee *transfert.Employee) (*entities.Employee, errors.ErrorInterface) {
	if dtoEmployee == nil {
		return nil, errors.ErrNoDto
	}

	employee, err := s.repo.ReadEmployee(&transfert.Employee{
		ID: dtoEmployee.ID,
	})

	if err != nil {
		return nil, errors.ErrEmployeeNotFound
	}

	if !s.security.CanUpdate(employee) {
		return nil, errors.ErrUnauthorized
	}

	data.UpdateEntityWithDto(employee, dtoEmployee)

	if err := s.repo.UpdateEmployee(employee); err != nil {
		return nil, errors.FromErr(err, errors.ErrInternalServer)
	}

	return employee, nil
}

func (s *UserService) DeleteEmployee(dtoEmployee *transfert.Employee) errors.ErrorInterface {
	if dtoEmployee == nil {
		return errors.ErrNoDto
	}

	employee, err := s.repo.ReadEmployee(dtoEmployee)
	if err != nil {
		return errors.ErrEmployeeNotFound
	}

	if !s.security.CanDelete(employee) {
		return errors.ErrUnauthorized
	}

	if err := s.repo.DeleteEmployee(dtoEmployee); err != nil {
		return errors.FromErr(err, errors.ErrInternalServer)
	}

	return nil
}

func (s *UserService) GetEmployee(dtoEmployee *transfert.Employee) (*entities.Employee, errors.ErrorInterface) {
	if dtoEmployee == nil {
		return nil, errors.ErrNoDto
	}

	employee, err := s.repo.ReadEmployee(dtoEmployee)
	if err != nil {
		return nil, errors.ErrEmployeeNotFound
	}

	if !s.security.CanRead(employee) {
		return nil, errors.ErrUnauthorized
	}

	return employee, nil
}
