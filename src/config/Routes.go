package config

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	admin "yardimbolgesi/admin/controllers"
	index "yardimbolgesi/index/controllers"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	// Admin routes
	r.GET("/admin", admin.Dashboard{}.Index)

	// List routes
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dashboard{}.Update)
	r.GET("/", index.Dashboard{}.Index)
	r.GET("/yeni-ekle", index.Dashboard{}.NewItem)
	r.POST("/add", index.Dashboard{}.Add)
	r.GET("/edit/:id", index.Dashboard{}.Edit)
	r.POST("/update/:id", index.Dashboard{}.UserUpdate)
	r.GET("/admin/updatereq", admin.Dashboard{}.UpdateReq)
	r.GET("/admin/detail/:id", admin.Dashboard{}.Detail)
	r.GET("/admin/reqdelete/:id", admin.Dashboard{}.ReqDelete)
	//r.GET("/admin/onay/:id", admin.Dashboard{}.Onay)
	r.GET("/admin/update-check/:id", admin.Dashboard{}.UpdateCheck)
	r.POST("/admin/updatecheck/:id", admin.Dashboard{}.UpdateCheckPost)
	r.GET("/detail/:id", index.Dashboard{}.Detail)

	// Userops routes
	r.GET("/admin/login", admin.Userops{}.Index)
	r.POST("/admin/do_login", admin.Userops{}.Login)

	// Serve static files
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets/"))
	r.ServeFiles("/index/assets/*filepath", http.Dir("index/assets/"))
	return r
}
