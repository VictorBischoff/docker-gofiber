package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/victorbischoff/structs-api/pkg/apistructs"
	"github.com/victorbischoff/structs-api/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []apistructs.User
	database.DB.Find(&users)
	return c.Status(200).JSON(users)
}

func NewUser(c *fiber.Ctx) error {
	user := new(apistructs.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	match := CheckPasswordHash(user.Password, hash)
	fmt.Println("Match:   ", match)

	user.Password = hash
	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}
	newUser := apistructs.User{
		UserName: user.UserName,
		Password: user.Password,
		Age:      user.Age,
		Email:    user.Email,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}
