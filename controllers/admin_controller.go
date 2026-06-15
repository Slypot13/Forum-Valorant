package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"forum-valorant/models"
	"forum-valorant/services"
)

// gère les actions d'administration.
type AdminController struct {
	adminService *services.AdminService
}

// contient les données de la page d'administration.
type AdminPage struct {
	Threads  []models.Thread
	Messages []models.Message
	Users    []models.User
}

// initialise le contrôleur d'administration.
func InitAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

// vérifie si l'utilisateur actuel est un administrateur.
func (c *AdminController) checkAdmin(w http.ResponseWriter, r *http.Request) bool {
	role, err := services.GetUserRoleFromRequest(r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
	}

	if role != "admin" {
		http.Error(w, "Accès refusé", http.StatusForbidden)
		return false
	}

	return true
}

// affiche la page du tableau de bord d'administration.
func (c *AdminController) ShowDashboard(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	threads, messages, users, err := c.adminService.GetDashboardData()

	if err != nil {
		http.Error(w, "Erreur chargement admin", http.StatusInternalServerError)
		return
	}

	page := AdminPage{
		Threads:  threads,
		Messages: messages,
		Users:    users,
	}

	tmpl := template.Must(template.ParseFiles("templates/admin.html"))
	tmpl.Execute(w, page)
}

// met à jour le statut d'un sujet (ouvert, fermé, archivé).
func (c *AdminController) UpdateThreadStatus(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	threadId, _ := strconv.Atoi(r.FormValue("thread_id"))
	status := r.FormValue("status")

	c.adminService.UpdateThreadStatus(threadId, status)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// supprime un sujet.
func (c *AdminController) DeleteThread(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	threadId, _ := strconv.Atoi(r.FormValue("thread_id"))

	c.adminService.DeleteThread(threadId)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// supprime un message.
func (c *AdminController) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	messageId, _ := strconv.Atoi(r.FormValue("message_id"))

	c.adminService.DeleteMessage(messageId)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// bannit un utilisateur.
func (c *AdminController) BanUser(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	userId, _ := strconv.Atoi(r.FormValue("user_id"))

	c.adminService.BanUser(userId)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}