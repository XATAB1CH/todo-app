package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	fmt.Fprintf(c.Writer, "signUp")
}

func (h *Handler) signIn(c *gin.Context) {

}
