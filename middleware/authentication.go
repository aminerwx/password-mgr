package middleware

import (
	"fmt"
	"net/http"

	"github.com/aminerwx/password-mgr/internal"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(c *gin.Context) {
	tokenStr, err := c.Cookie("token")
	if err != nil {
		fmt.Println("Token missing in cookie")
		c.Redirect(http.StatusSeeOther, "/auth/signin")
		c.Abort()
		return
	}
	_, err = internal.VerifyTokenJWT(tokenStr)
	if err != nil {
		fmt.Printf("Token verification failed: %v\\n", err)
		c.Redirect(http.StatusSeeOther, "/auth/signin")
		c.Abort()
		return
	}
	c.Next()
}
