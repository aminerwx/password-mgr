package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetVaultHandler(c *gin.Context) {
}

func (s *Server) CreateVaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (s *Server) UpdateVaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func (s *Server) RemoveVaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
