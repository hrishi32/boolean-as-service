package controller

import (
	"github.com/gin-gonic/gin"
)

// Handle400 handles bad request error
func Handle400(c *gin.Context, err error) {
	c.Writer.WriteHeader(400)
}

// Handle404 Handles content not found error
func Handle404(c *gin.Context, err error) {
	c.Writer.WriteHeader(404)
}

// Handle500 handles internal server error
func Handle500(c *gin.Context, err error) {
	c.Writer.WriteHeader(500)
}

// DatabaseError handles all errors from database, including connection errors and record not found errors.
// func DatabaseError(c *gin.Context, err error) {
// 	// fmt.Println("Error from database", err)
// 	// log.Fatal(err)
// 	c.JSON(http.StatusInternalServerError, gin.H{
// 		"message": "Error from database",
// 		"error":   err.Error(),
// 	})
// }

// // ParseError handles errors regarding parsing uuid from requests.
// func ParseError(c *gin.Context, err error) {
// 	fmt.Println("Error while parsing uuid from request", err)
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"message": "Error while parsing uuid from request",
// 		"error":   err.Error(),
// 	})
// }

// // BindError handles JSON binding errors.
// func BindError(c *gin.Context, err error) {
// 	fmt.Println("Error while binding JSON into struct", err)
// 	c.JSON(http.StatusBadRequest, gin.H{
// 		"message": "Error while binding JSON into struct",
// 		"error":   err.Error(),
// 	})
// }
