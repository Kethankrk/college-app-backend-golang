package attendance

import (
	"college-app/database"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type request struct {
	Starting primitive.DateTime `json:"start" binding:"required"`
	Ending   primitive.DateTime `json:"end" binding:"required"`
	Class    string             `json:"class" binding:"required"`
	Name     string             `json:"name"`
}

func GetAttendance(c *gin.Context) {
	var req request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
			"error":   err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	defer cancel()

	collection := database.AttendanceDatabase.Collection(req.Class)

	var cursor *mongo.Cursor
	var db_err error

	if req.Name != "" {
		cursor, db_err = collection.Find(ctx, bson.M{
			"$and": []bson.M{
				{"date": bson.M{"$gte": req.Starting}},
				{"date": bson.M{"$lt": req.Ending}},
				{"student": req.Name},
			},
		})

	} else {
		cursor, db_err = collection.Find(ctx, bson.M{
			"date": bson.M{
				"$gte": req.Starting,
				"$lt":  req.Ending,
			},
		})
	}

	if db_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	var result []database.Attendance

	if err := cursor.All(ctx, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": http.StatusText(http.StatusInternalServerError),
			"error":   err.Error(),
		})
		return
	}

	if result == nil {
		c.JSON(http.StatusNoContent, gin.H{
			"code":    http.StatusNoContent,
			"message": http.StatusText(http.StatusNoContent),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"message":    http.StatusText(http.StatusOK),
		"attendance": result,
	})

}
