package items

import (
	"api/app/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetItem ...
func GetItem(c *gin.Context) {
	itemID := strings.TrimSpace(c.Param("id"))
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_error"})
		return
	}

	item, err := Is.Item(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_error", "description": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
	return
}

// PostItem ...
func PostItem(c *gin.Context) {
	i := &models.Item{}
	if err := c.BindJSON(i); c.Request.ContentLength == 0 || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bind_error", "description": err.Error()})
		return
	}
	err := Is.CreateItem(i)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "save_error", "description": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, i)
}

// GetItems ...
func GetItems(c *gin.Context) {
	items, err := Is.Items()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find_error", "description": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
	return
}

// DeleteItem ...
func DeleteItem(c *gin.Context) {
	itemID := strings.TrimSpace(c.Param("id"))
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_error"})
		return
	}
	err := Is.DeleteItem(itemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete_error", "description": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item eliminado satisfactoriamente"})
	return
}
