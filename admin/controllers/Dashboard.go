package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"yardimbolgesi/admin/helpers"
	"yardimbolgesi/admin/models"
	index "yardimbolgesi/index/models"
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

func (dashboard Dashboard) UpdateReq(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/update")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Update"] = index.UserUpdate{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) Detail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/detail")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Detail"] = models.List{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) ReqDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	update := index.UserUpdate{}.Get(params.ByName("id"))
	update.Delete()
	http.Redirect(w, r, "/admin/updatereq", http.StatusSeeOther)
}

func (dashboard Dashboard) UpdateCheck(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/update_check")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Update"] = index.UserUpdate{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}

func (dashboard Dashboard) UpdateCheckPost(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
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

	update := index.UserUpdate{}.Get(params.ByName("update_id"))
	update.Delete()

	http.Redirect(w, r, "/admin/updatereq/", http.StatusSeeOther)
}
