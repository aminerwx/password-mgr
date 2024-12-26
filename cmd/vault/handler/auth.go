package handler

import (
	"github.com/aminerwx/password-mgr/cmd/vault/template"
	"github.com/gin-gonic/gin"
)

func (s *Server) ViewSigninHandler(c *gin.Context) {
	view := template.ViewSigninTemplate()
	view.Render(c.Request.Context(), c.Writer)
}

func (s *Server) SubmitSigninHandler(c *gin.Context) {
}

func (s *Server) ViewSignupHandler(c *gin.Context) {
}

func (s *Server) SubmitSignupHandler(c *gin.Context) {
}
