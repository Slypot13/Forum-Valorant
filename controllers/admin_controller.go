package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"forum-valorant/models"
	"forum-valorant/services"
)

type AdminController struct {
	adminService *services.AdminService
}

type AdminPage struct {
	Threads  []models.Thread
	Messages []models.Message
	Users    []models.User
}

func InitAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{
		adminService: adminService,
	}
}

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

func (c *AdminController) UpdateThreadStatus(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	threadId, _ := strconv.Atoi(r.FormValue("thread_id"))
	status := r.FormValue("status")

	c.adminService.UpdateThreadStatus(threadId, status)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (c *AdminController) DeleteThread(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	threadId, _ := strconv.Atoi(r.FormValue("thread_id"))

	c.adminService.DeleteThread(threadId)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (c *AdminController) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	messageId, _ := strconv.Atoi(r.FormValue("message_id"))

	c.adminService.DeleteMessage(messageId)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (c *AdminController) BanUser(w http.ResponseWriter, r *http.Request) {
	if !c.checkAdmin(w, r) {
		return
	}

	userId, _ := strconv.Atoi(r.FormValue("user_id"))

	c.adminService.BanUser(userId)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}