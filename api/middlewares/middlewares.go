package middlewares

import (
	"errors"
	"net/http"

	"github.com/srigalamilitan/fullstack_learn/api/auth"
	"github.com/srigalamilitan/fullstack_learn/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("UnAuthorized"))
			return
		}
		next(w, r)
	}
}
