package helpers

import "net/http"

func SetUser(w http.ResponseWriter, r *http.Request, username, password string) {
	cookie := http.Cookie{
		Name:  "user",
		Value: username + ":" + password,
	}
	http.SetCookie(w, &cookie)
}
