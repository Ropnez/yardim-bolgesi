package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"yardimbolgesi/admin/helpers"
	"yardimbolgesi/admin/models"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["List"] = models.List{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	content := r.FormValue("content")
	city := r.FormValue("city")
	distict := r.FormValue("distict")
	adress := r.FormValue("adress")
	number := r.FormValue("number")
	description := r.FormValue("description")

	models.List{
		Content:     content,
		City:        city,
		Distict:     distict,
		Adress:      adress,
		Number:      number,
		Description: description,
	}.Add()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	list := models.List{}.Get(params.ByName("id"))
	list.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["List"] = models.List{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	list := models.List{}.Get(params.ByName("id"))
	content := r.FormValue("content")
	city := r.FormValue("city")
	distict := r.FormValue("distict")
	adress := r.FormValue("adress")
	number := r.FormValue("number")
	description := r.FormValue("description")

	list.Updates(models.List{
		Content:     content,
		City:        city,
		Distict:     distict,
		Adress:      adress,
		Number:      number,
		Description: description,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
