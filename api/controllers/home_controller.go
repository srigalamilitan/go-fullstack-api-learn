package controllers

import (
	"net/http"

	"github.com/srigalamilitan/fullstack_learn/api/responses"
)

func (server *Server) home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcode")
}
