package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) balance(c *gin.Context) {
	uid := c.Param("uid")
	balance, err := h.services.Balance.Balance(uid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "wallet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Balance": balance})
}
