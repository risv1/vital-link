package controllers

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"time"
	"vital-link/api/database"
	"vital-link/api/models"
	"vital-link/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var clientID string
var clientSecret string
var googleOauthConfig *oauth2.Config

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	clientID = os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectUrl := "http://localhost:8000/auth/google/callback"

	googleOauthConfig =  &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func Login(c *fiber.Ctx) error {
	url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
    return c.Redirect(url)
}	

func Callback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(c.Context(), code)
	if err != nil {
		return c.Redirect("/api/login")
	}

	client := googleOauthConfig.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return c.Redirect("/api/login")
	}
	defer res.Body.Close()

	userInfo, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Redirect("/api/login")
	}

	var user map[string]interface{}
	if err := json.Unmarshal(userInfo, &user); err != nil {
		return c.Redirect("/api/login")
	}

	newUser := models.User{
		ID: uuid.New().String(),
		Name: user["given_name"].(string),
		Email: user["email"].(string),
		Role: "user",
		ProfileId: uuid.New().String(),
	}

	db := database.GetDatabase()

	collection := db.Collection("users")
	findUser := models.User{}
	err = collection.FindOne(c.Context(), models.User{Email: newUser.Email}).Decode(&findUser)
	if err == nil {
		jwtToken, err := utils.GenerateJWT(findUser.ID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to generate JWT",
			})
		}

		c.Cookie(&fiber.Cookie{
			Name: "token",
			Value: jwtToken,
			Expires: time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
			SameSite: "None",
		})

		return c.Status(200).JSON(fiber.Map{
			"message": "Login successful",
		})
	}

	_, err = collection.InsertOne(c.Context(), newUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to insert user",
		})
	}

	jwtToken, err := utils.GenerateJWT(newUser.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to generate JWT",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: jwtToken,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: "None",
	})

	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful",
	})
}

func Logout (c *fiber.Ctx) error {

	token := c.Cookies("token")
	if token == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Missing token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name: "token",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
		SameSite: "None",
	})

	return c.Status(200).JSON(fiber.Map{
		"message": "Logout successful",
	})
}