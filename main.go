package main

import (
	"net/http"
	admin_models "yardimbolgesi/admin/models"
	"yardimbolgesi/config"
	index_models "yardimbolgesi/index/models"
)

func main() {
	admin_models.List{}.Migrate()
	admin_models.User{}.Migrate()
	index_models.UserUpdate{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
