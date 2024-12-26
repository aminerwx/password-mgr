package handler

import (
	"net/http"

	"github.com/aminerwx/password-mgr/cmd/vault/model"
	"github.com/aminerwx/password-mgr/cmd/vault/repository"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store repository.Repository
	port  string
}

func NewServer(store repository.Repository, port string) *Server {
	return &Server{store: store, port: port}
}

//func (s *Server) Start() error {
//	mux := http.NewServeMux()
//	mux.HandleFunc("POST /users", s.CreateUserHandler)
//	mux.HandleFunc("GET /users/{id}", s.GetUserHandler)
//	mux.HandleFunc("PUT /users/{id}", s.UpdateUserHandler)
//	mux.HandleFunc("DELETE /users/{id}", s.RemoveUserHandler)
//	mux.HandleFunc("POST /vaults", s.CreateVaultHandler)
//	mux.HandleFunc("GET /vaults/{id}", s.GetVaultHandler)
//	mux.HandleFunc("PUT /vaults/{id}", s.UpdateVaultHandler)
//	mux.HandleFunc("DELETE /vaults/{id}", s.RemoveVaultHandler)
//	wrappedMux := middleware.NewLogger(mux)
//	return http.ListenAndServe(s.port, wrappedMux)
//}

func (s *Server) Start() error {
	router := gin.Default()
	router.GET("/auth/signin", s.ViewSigninHandler)
	router.POST("/auth/signin", s.SubmitSigninHandler)
	router.GET("/auth/signup", s.ViewSignupHandler)
	router.POST("/auth/signup", s.SubmitSignupHandler)
	router.GET("/vault", s.GetVaultHandler)
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
