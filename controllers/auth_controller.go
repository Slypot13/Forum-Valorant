package controllers

import (
	"html/template"
	"net/http"

	"forum-valorant/models"
	"forum-valorant/services"
)

// gère l'inscription et la connexion.
type AuthController struct {
	userService *services.UserService
}

// initialise le contrôleur.
func InitAuthController(userService *services.UserService) *AuthController {
	return &AuthController{
		userService: userService,
	}
}

// affiche la page d'inscription.
func (c *AuthController) ShowRegister(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, nil)
}

// traite le formulaire d'inscription.
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

//  affiche la page de connexion.
func (c *AuthController) ShowLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}

// Login connecte l'utilisateur et crée le cookie.
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	identifier := r.FormValue("identifier")
	password := r.FormValue("password")
	token, err := c.userService.Login(identifier, password)

	if err != nil {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, err.Error())
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}