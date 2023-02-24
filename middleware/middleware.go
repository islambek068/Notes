package middleware

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// Nosurf CSRF protection to all POST request
/*
A middleware handler is simply an http.Handler
that wraps another http.Handler to do some
pre- and/or post-processing of the request.
It's called "middleware" because it sits
in the middle between the Go web server and
the actual handler.

nosurf is an HTTP package for Go that helps
you prevent Cross-Site Request Forgery attacks.
It acts like a middleware and therefore is
compatible with basically any Go HTTP application.*/
// грубо говоря нужна для защиты

func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
