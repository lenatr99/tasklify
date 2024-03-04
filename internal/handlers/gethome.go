package handlers

import (
	"net/http"
	"tasklify/internal/templates"

	"github.com/go-chi/jwtauth/v5"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_, claims, _ := jwtauth.FromContext(r.Context())

	email, ok := claims["email"].(string)

	if !ok {
		c := templates.GuestIndex()
		err := templates.Layout(c, "Tasklify").Render(r.Context(), w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

		return
	}

	c := templates.Index(email)
	err := templates.Layout(c, "Tasklify").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
