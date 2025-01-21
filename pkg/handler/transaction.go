package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) transaction(c *gin.Context) {
	count := c.DefaultQuery("count", "1")

	countInt, err := strconv.Atoi(count)
	if err != nil {
		fmt.Println("Invalid 'count' parameter:", count)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'count' parameter"})
		return
	}

	transactions, err := h.services.Transactions(countInt)
	if err != nil {
		fmt.Println("Error fetching transactions:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Transactions fetched",
		"count":        len(transactions),
		"transactions": transactions,
	})
}
