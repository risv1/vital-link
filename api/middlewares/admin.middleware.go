package middlewares

import (
	"vital-link/api/database"
	"vital-link/api/utils"
	"sync"
	"github.com/gofiber/fiber/v2"
)

type UserCache struct {
	mu sync.RWMutex
	cache map[string]fiber.Map
}

var users = UserCache{cache: make(map[string]fiber.Map)}

func CheckAdmin(c *fiber.Ctx) error {
	token := c.Cookies("jwt")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	tokenVal, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	userId := tokenVal["id"].(string)

	users.mu.RLock()
	userData, found := users.cache[userId]
	users.mu.RUnlock()

	if !found{
		db := database.GetDatabase()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Database connection failed",
			})
		}

		collection := db.Collection("users")
		user := collection.FindOne(c.Context(), fiber.Map{"_id": tokenVal})
		if user.Err() != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		userData = fiber.Map{}
		user.Decode(&userData)

		users.mu.Lock()
		users.cache[userId] = userData
		users.mu.Unlock()
	}

	if userData["role"] != "admin" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}