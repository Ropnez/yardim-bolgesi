package controllers

import (
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"yardimbolgesi/admin/helpers"
	"yardimbolgesi/admin/models"
)

type Userops struct{}

func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}

func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password"))))

	user := models.User{}.Get("username = ? AND password = ?", username, password)
	if (user.Username == username) && (user.Password == password) {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}

}
