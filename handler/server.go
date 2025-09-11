package handler

import (
	"net/http"

	"github.com/aminerwx/password-mgr/middleware"
	"github.com/aminerwx/password-mgr/model"
	"github.com/aminerwx/password-mgr/repository"
	"github.com/aminerwx/password-mgr/view"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store repository.Repository
	port  string
}

func NewServer(store repository.Repository, port string) *Server {
	return &Server{store: store, port: port}
}

func (s *Server) Start() error {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(ctx *gin.Context) {
		view.Render(ctx, view.ViewIndexTemplate)
	})

	router.GET("/auth/signin", s.ViewSigninHandler)
	router.POST("/auth/signin", s.SubmitSigninHandler)
	router.GET("/auth/signup", s.ViewSignupHandler)
	router.POST("/auth/signup", s.SubmitSignupHandler)
	router.GET("/vault", middleware.AuthenticationMiddleware, s.GetVaultHandler)

	return router.Run(s.port)
}

type Response struct {
	Message    string     `json:"message"`
	StatusCode int        `json:"status_code"`
	Data       model.User `json:"data"`
}

var (
	StatusBadRequestJSON = Response{
		Message:    "bad request",
		StatusCode: http.StatusBadRequest,
	}
	StatusConflictJSON = Response{
		Message:    "already exist",
		StatusCode: http.StatusConflict,
	}
	StatusNotFoundJSON = Response{
		Message:    "not found",
		StatusCode: http.StatusNotFound,
	}
	StatusOkJSON = Response{
		Message:    "success",
		StatusCode: http.StatusOK,
	}
	StatusCreatedJSON = Response{
		Message:    "success",
		StatusCode: http.StatusCreated,
	}
)
