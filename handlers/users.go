package handlers

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/markbates/goth/gothic"
)

func Users(db gorm.DB, ren Renderer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := getPageData(db, r)

		if needsLogin(r) {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			ren.Render(w, http.StatusOK, "users", data, "layout")
		}
	})
}

func needsLogin(r *http.Request) bool {

	session, err := gothic.Store.Get(r, gothic.SessionName)
	if err != nil {
		fmt.Println("Sesson Error: ", err)
	}
	return session.IsNew
}