package handler

import (
	"log"

	"github.com/aminerwx/password-mgr/view"
	"github.com/gin-gonic/gin"
)

func (s *Server) ViewSigninHandler(c *gin.Context) {
	v := view.ViewSigninTemplate()
	v.Render(c.Request.Context(), c.Writer)
}

func (s *Server) SubmitSigninHandler(c *gin.Context) {
	username := c.PostForm("username")
	masterPassword := c.PostForm("master")
	if username == "dorakyura" {
		log.Println(username)
		log.Println(masterPassword)
	}
}

func (s *Server) ViewSignupHandler(c *gin.Context) {
}

func (s *Server) SubmitSignupHandler(c *gin.Context) {
}
