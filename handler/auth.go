package handler

import (
	"log"
	"net/http"

	"github.com/aminerwx/password-mgr/internal"
	"github.com/aminerwx/password-mgr/model"
	"github.com/aminerwx/password-mgr/view"
	"github.com/gin-gonic/gin"
)

func (s *Server) ViewSigninHandler(c *gin.Context) {
	view.Render(c, view.ViewSigninTemplate)
}

func (s *Server) SubmitSigninHandler(c *gin.Context) {
	username := c.PostForm("username")
	masterPassword := c.PostForm("master")
	// TODO:
	// Check user/master match db
	user, err := s.store.GetUser(username)
	if err != nil {
		log.Println(err)
		return
	}
	isValid, _, _ := internal.Compare(masterPassword, user.MasterPassword)
	if isValid {
		tokenStr, err := internal.CreateTokenJWT(username)
		if err != nil {
			log.Println(err)
			return
		}
		c.SetCookie("token", tokenStr, 3600, "/", "localhost", false, true)
		c.Set("userID", user.ID)
		c.Redirect(http.StatusSeeOther, "/vault")
	} else {
		log.Println("Invalid password")
		c.Redirect(http.StatusSeeOther, "/auth/signin")
	}
	// TODO:
	// handle user not found
}

func (s *Server) ViewSignupHandler(c *gin.Context) {
	view.Render(c, view.ViewSignupTemplate)
}

func (s *Server) SubmitSignupHandler(c *gin.Context) {
	username := c.PostForm("username")
	master := c.PostForm("master")
	confirm := c.PostForm("confirm")
	// TODO:
	// Check user/master match db
	if master == confirm {
		user, _ := model.NewUser(username, master)
		err := s.store.CreateUser(user)
		if err != nil {
			log.Println(err)
			return
		}

		err = s.store.CreateVault(user)
		if err != nil {
			log.Println(err)
			return
		}

		tokenStr, err := internal.CreateTokenJWT(username)
		if err != nil {
			log.Println(err)
			return
		}

		c.SetCookie("token", tokenStr, 3600, "/", "localhost", false, true)
		c.Set("userID", user.ID)
		c.Redirect(http.StatusSeeOther, "/vault")
	} else {
		c.Redirect(http.StatusSeeOther, "/auth/signup")
	}
	// handle user not found
}
