package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf add csrf protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves a session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
