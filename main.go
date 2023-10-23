package main

import (
	"net/http"
	admin_models "yardimbolgesi/admin/models"
	"yardimbolgesi/config"
)

func main() {
	admin_models.List{}.Migrate()
	admin_models.User{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
