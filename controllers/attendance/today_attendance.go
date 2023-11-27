package attendance

import (
	"college-app/database"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TodayAttendance(c *gin.Context) {
	class := c.Param("class")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	collection := database.AttendanceDatabase.Collection(class)

	todayStart := time.Now().UTC().Truncate(24 * time.Hour)
	todayEnd := todayStart.Add(24 * time.Hour)

	cursour, db_err := collection.Find(ctx, bson.M{
		"date": bson.M{
			"$gte": todayStart,
			"$lt":  todayEnd,
		},
	})

	if db_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	var result []database.Attendance

	if decode_err := cursour.All(ctx, &result); decode_err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": http.StatusText(http.StatusInternalServerError),
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
