package attendance

import (
	"college-app/database"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AttendaceRequest struct {
	Date    primitive.DateTime `json:"date" binding:"required"`
	Period  int                `json:"period" binding:"required"`
	Student string             `json:"student" binding:"required"`
	Present *bool              `json:"present" binding:"required"`
	Class   string             `json:"class" binding:"required"`
}

func AddAttendance(c *gin.Context) {

	var req AttendaceRequest
	var att database.AttendanceWithOutId

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
			"error":   err.Error(),
		})
		return
	}

	att.Date = req.Date
	att.Period = req.Period
	att.Student = req.Student
	att.Present = *req.Present

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	collection := database.AttendanceDatabase.Collection(req.Class)

	_, err := collection.InsertOne(ctx, att)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": http.StatusText(http.StatusCreated),
	})
}
