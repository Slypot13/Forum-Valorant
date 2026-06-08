package controllers

import (
	"html/template"
	"net/http"

	"forum-valorant/models"
	"forum-valorant/services"
)

type AuthController struct {
	userService *services.UserService
}

func InitAuthController(userService *services.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

func (c *AuthController) ShowRegister(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, nil)
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := models.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := c.userService.Register(user)

	if err != nil {
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, err.Error())
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}