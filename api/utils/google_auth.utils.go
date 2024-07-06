package utils 

import (
    "github.com/gofiber/fiber/v2"
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
	"fmt"
	"github.com/joho/godotenv"
	"os"
) 

var clientID string
var clientSecret string
var googleOauthConfig *oauth2.Config

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	clientID = os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

	googleOauthConfig =  &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8000/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleLogin(c *fiber.Ctx) error {
    url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
    return c.Redirect(url)
}

func GoogleCallback(c *fiber.Ctx) error {
    code := c.Query("code")
	token, err := googleOauthConfig.Exchange(c.Context(), code)
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/api/login")
	}
	return c.JSON(token)
}
