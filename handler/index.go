package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	response := gin.H{
		"status": "OK! 🚀",
		"developer": map[string]interface{}{
			"name":     "Sahibul Nuzul Firdaus",
			"email":    "sahibulnuzulfirdaus13@gmail.com",
			"linkedIn": "https://www.linkedin.com/in/sahibul-nf/",
			"whatsapp": "https://wa.link/r3amjb",
		},
	}

	ctx.JSON(http.StatusOK, response)
}
