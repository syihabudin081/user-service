package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"user-service/models"
	"user-service/services"
	"user-service/utils"
)

type UserController struct {
	Service   services.UserService
	Validator *validator.Validate
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		Service:   service,
		Validator: validator.New(),
	}
}
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	userModel := new(models.User)
	if err := c.BodyParser(userModel); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := uc.Validator.Struct(userModel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})

	}
	userModel.Password = string(hashedPassword)

	if err := uc.Service.CreateUser(userModel); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to create user",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(&utils.Response{
		Status:  "success",
		Message: "User created successfully",
		Data:    userModel,
	})
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	users, err := uc.Service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to fetch users",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "Users fetched successfully",
		Data:    users,
	})
}

func (uc *UserController) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := uc.Service.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&utils.Response{
			Status:  "error",
			Message: "User not found",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "User fetched successfully",
		Data:    user,
	})
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userUpdate := new(models.User)
	if err := c.BodyParser(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to parse request body",
			Error:   err.Error(),
		})
	}
	if err := uc.Validator.StructPartial(userUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}

	userModelID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
			Status:  "error",
			Message: "Invalid user ID",
			Error:   err.Error(),
		})
	}

	userUpdate.ID = userModelID
	if err := uc.Service.UpdateUser(userUpdate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to update user",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "User updated successfully",
	})
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uc.Service.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&utils.Response{
			Status:  "error",
			Message: "Failed to delete user",
			Error:   err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(&utils.Response{
		Status:  "success",
		Message: "User deleted successfully",
	})
}
