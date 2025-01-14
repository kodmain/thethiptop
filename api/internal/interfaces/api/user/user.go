package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kodmain/thetiptop/api/config"
	"github.com/kodmain/thetiptop/api/internal/application/security"
	services "github.com/kodmain/thetiptop/api/internal/application/services/user"
	transfert "github.com/kodmain/thetiptop/api/internal/application/transfert/user"
	gameRepository "github.com/kodmain/thetiptop/api/internal/domain/game/repositories"
	"github.com/kodmain/thetiptop/api/internal/domain/user/repositories"
	domain "github.com/kodmain/thetiptop/api/internal/domain/user/services"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/errors"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/providers/database"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/providers/mail"
	"github.com/kodmain/thetiptop/api/internal/infrastructure/serializers/jwt"
)

// @Tags		User
// @Summary		Authenticate a client/employees.
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		email		formData	string	true	"Email address" format(email) default(user-thetiptop@yopmail.com)
// @Param		password	formData	string	true	"Password" default(Aa1@azetyuiop)
// @Success		200	{object}	nil "Client signed in"
// @Failure		400	{object}	nil "Invalid email or password"
// @Failure		500	{object}	nil "Internal server error"
// @Router		/user/auth [post]
// @Id			user.UserAuth
func UserAuth(ctx *fiber.Ctx) error {
	dto := &transfert.Credential{}
	if err := ctx.BodyParser(dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	status, response := services.UserAuth(
		domain.User(
			security.NewUserAccess(ctx.Locals("token")),
			repositories.NewUserRepository(database.Get(config.GetString("services.client.database", config.DEFAULT))),
			gameRepository.NewGameRepository(database.Get(config.GetString("services.game.database", config.DEFAULT))),
			mail.Get(config.GetString("services.client.mail", config.DEFAULT)),
		), dto,
	)

	return ctx.Status(status).JSON(response)
}

// @Tags		User
// @Summary		Renew JWT for a client/employees.
// @Accept		*/*
// @Accept		multipart/form-data
// @Produce		application/json
// @Success		200	{object}	nil "JWT token renewed"
// @Failure		400	{object}	nil "Invalid token"
// @Failure		401	{object}	nil "Token expired"
// @Failure		500	{object}	nil "Internal server error"
// @Param 		Authorization header string true "With the bearer started"
// @Router		/user/auth/renew [get]
// @Id			user.UserAuthRenew
func UserAuthRenew(ctx *fiber.Ctx) error {
	token := ctx.Locals("token")
	if token == nil {
		return ctx.Status(errors.ErrAuthNoToken.Code()).JSON(errors.ErrAuthNoToken)
	}

	status, response := services.UserAuthRenew(
		token.(*jwt.Token),
	)

	return ctx.Status(status).JSON(response)
}

// @Tags		User
// @Summary		Update a client/employees password.
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		email		formData	string	true	"Email address" format(email) default(user-thetiptop@yopmail.com)
// @Param		password	formData	string	true	"Password" default(Aa1@azetyuiop)
// @Param		token		formData	string	true	"Token"
// @Success		204	{object}	nil "Password updated"
// @Failure		400	{object}	nil "Invalid email, password or token"
// @Failure		404	{object}	nil "Client not found"
// @Failure		409	{object}	nil "Client already validated"
// @Failure		410	{object}	nil "Token expired"
// @Failure		500	{object}	nil "Internal server error"
// @Router		/user/password [put]
// @Id			user.CredentialUpdate
// @Security 	Bearer
func CredentialUpdate(ctx *fiber.Ctx) error {
	dtoCredential := &transfert.Credential{}
	if err := ctx.BodyParser(dtoCredential); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	dtoValidation := &transfert.Validation{}
	if err := ctx.BodyParser(dtoValidation); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	status, response := services.CredentialUpdate(
		domain.User(
			security.NewUserAccess(ctx.Locals("token")),
			repositories.NewUserRepository(database.Get(config.GetString("services.client.database", config.DEFAULT))),
			gameRepository.NewGameRepository(database.Get(config.GetString("services.game.database", config.DEFAULT))),
			mail.Get(config.GetString("services.client.mail", config.DEFAULT)),
		), dtoValidation, dtoCredential,
	)

	return ctx.Status(status).JSON(response)
}

// @Tags		User
// @Summary		Validate a client/employees email.
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		token	formData	string	true	"Token"
// @Param		email	formData	string	true	"Email address" format(email) default(user-thetiptop@yopmail.com)
// @Success		204	{object}	nil "Client email validate"
// @Failure		400	{object}	nil "Invalid email or token"
// @Failure		404	{object}	nil "Client not found"
// @Failure		409	{object}	nil "Client already validated"
// @Failure		410 {object}	nil "Token expired"
// @Failure		500	{object}	nil "Internal server error"
// @Router		/user/register/validation [put]
// @Id			user.MailValidation
func MailValidation(ctx *fiber.Ctx) error {
	dtoCredential := &transfert.Credential{}
	if err := ctx.BodyParser(dtoCredential); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	dtoValidation := &transfert.Validation{}
	if err := ctx.BodyParser(dtoValidation); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	status, response := services.MailValidation(
		domain.User(
			security.NewUserAccess(ctx.Locals("token")),
			repositories.NewUserRepository(database.Get(config.GetString("services.client.database", config.DEFAULT))),
			gameRepository.NewGameRepository(database.Get(config.GetString("services.game.database", config.DEFAULT))),
			mail.Get(config.GetString("services.client.mail", config.DEFAULT)),
		), dtoValidation, dtoCredential,
	)

	return ctx.Status(status).JSON(response)
}

// @Tags		User
// @Summary		Recover a client/employees validation type.
// @Accept		multipart/form-data
// @Produce		application/json
// @Param		email		formData	string	true	"Email address" format(email) default(user-thetiptop@yopmail.com)
// @Param		type		formData	string	true	"Type of validation" enums(mail, password, phone)
// @Router		/user/validation/renew [post]
// @Id			user.ValidationRecover
func ValidationRecover(ctx *fiber.Ctx) error {
	dtoCredential := &transfert.Credential{}
	if err := ctx.BodyParser(dtoCredential); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	dtoValidation := &transfert.Validation{}
	if err := ctx.BodyParser(dtoValidation); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	status, response := services.ValidationRecover(
		domain.User(
			security.NewUserAccess(ctx.Locals("token")),
			repositories.NewUserRepository(database.Get(config.GetString("services.client.database", config.DEFAULT))),
			gameRepository.NewGameRepository(database.Get(config.GetString("services.game.database", config.DEFAULT))),
			mail.Get(config.GetString("services.client.mail", config.DEFAULT)),
		), dtoCredential, dtoValidation,
	)

	return ctx.Status(status).JSON(response)
}
