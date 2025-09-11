package view

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, v func() templ.Component) {
	myView := v()
	myView.Render(c.Request.Context(), c.Writer)
}
