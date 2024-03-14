package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next) //csrf is: cross site request forgery token
	//it's a security measure to prevent attacks
	//it changes the value of the token every time the page is refreshed

	// csrfHandler sets a cookie to the browser
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// loads and saves session data for every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
