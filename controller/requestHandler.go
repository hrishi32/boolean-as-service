package controller

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/hrishi32/boolean-as-service/models"
)

// GetHandler handles GET request of server by using model's get function.
func GetHandler(c *gin.Context) {
	id, parseError := uuid.Parse(c.Param("id"))
	if parseError != nil {
		// ParseError(c, parseError)
		Handle400(c, parseError)
		return
	}

	b, databaseError := models.GetRepo().Get(id)
	if databaseError != nil && databaseError.Error() == "Record not found" {
		Handle404(c, databaseError)
		return
	}

	if databaseError != nil {
		Handle500(c, databaseError)
		return
	}

	c.JSON(200, gin.H{
		"id":    b.ID,
		"value": b.Value,
		"key":   b.Key,
	})
}

// PostHandler handles POST request of server by usning model's Create function.
// It returns saved boolean object with uuid assigned to it, or an error in JSON format.
func PostHandler(c *gin.Context) {
	var b models.Boolean

	bindError := c.ShouldBindJSON(&b)
	if bindError != nil {
		// BindError(c, bindError)
		Handle400(c, bindError)
		return
	}

	bID, databaseError := models.GetRepo().Create(b)
	// Error must be 500 in this case, because we are not fetching anything from database.
	if databaseError != nil {
		// DatabaseError(c, databaseError)
		Handle500(c, databaseError)
		return
	}

	b.ID = bID

	c.JSON(200, gin.H{
		"id":    b.ID,
		"value": b.Value,
		"key":   b.Key,
	})
}

// PatchHandler handles PATCH request of server by using model's Update method.
func PatchHandler(c *gin.Context) {
	id, parseError := uuid.Parse(c.Param("id"))
	if parseError != nil {
		// ParseError(c, parseError)
		Handle400(c, parseError)
		return
	}

	var b models.Boolean
	bindError := c.ShouldBindJSON(&b)
	if bindError != nil {
		// BindError(c, bindError)
		Handle400(c, bindError)
		return
	}

	databaseError := models.GetRepo().Update(id, b)
	if databaseError != nil && databaseError.Error() == "Record not found" {
		// DatabaseError(c, databaseError)
		Handle404(c, databaseError)
		return
	}

	if databaseError != nil {
		Handle500(c, databaseError)
		return
	}

	c.JSON(200, gin.H{
		"id":    id,
		"value": b.Value,
		"key":   b.Key,
	})
}

// DeleteHandler handles DELETE request of server by using model's Delete method.
func DeleteHandler(c *gin.Context) {
	id, parseError := uuid.Parse(c.Param("id"))

	if parseError != nil {
		// ParseError(c, parseError)
		Handle400(c, parseError)
		return
	}

	databaseError := models.GetRepo().Delete(id)

	if databaseError != nil && databaseError.Error() == "Record not found" {
		Handle404(c, databaseError)
		return
	}

	if databaseError != nil {
		Handle500(c, databaseError)
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
