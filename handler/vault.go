package handler

import (
	"net/http"

	"github.com/aminerwx/password-mgr/view"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetVaultHandler(c *gin.Context) {
	component := view.ViewVaultTemplate(val)
	component.Render(c.Request.Context(), c.Writer)
	// TODO:
	// Fetch
	// Decrypt
	// Render
}

func (s *Server) CreateVaultHandler(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// Create Vault on signup

func (s *Server) UpdateVaultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// Update
	// Encrypt
	// Save
	w.Header().Set("Content-Type", "application/json")
}

func (s *Server) RemoveVaultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// Maybe not needed???
	w.Header().Set("Content-Type", "application/json")
}
