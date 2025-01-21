package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet"
)

func (h *Handler) send(c *gin.Context) {
	var input wallet.Transaction

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.From == input.To {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot transfer to the same wallet"})
		return
	}

	id, err := h.services.Send.Send(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"Transaction completed, transaction identification number": id})
}
