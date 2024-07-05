package controllers

import (
	"github.com/gofiber/fiber/v2"
	"vital-link/utils"
)

func Login(c *fiber.Ctx) error {
	return utils.GoogleLogin(c)
}

func Callback(c *fiber.Ctx) error {
	return utils.GoogleCallback(c)
}

func Logout (c *fiber.Ctx) error {
	return c.Redirect("/api/login")
}