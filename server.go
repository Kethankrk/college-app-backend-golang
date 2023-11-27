package main

import (
	"college-app/controllers/attendance"
	"college-app/controllers/students"
	"college-app/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("./.env")
	database.InitMongo()
}

func main() {
	app := gin.Default()

	app.GET("/", students.StudentLogin)
	app.POST("/add-attendance", attendance.AddAttendance)
	app.POST("/get-attendance", attendance.GetAttendance)
	app.GET("/today-attendance/:class", attendance.TodayAttendance)

	app.Run()
}
