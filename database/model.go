package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Type     string             `json:"type" bson:"type"`
}

type Student struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Class string             `json:"class" bson:"class"`
}

type Attendance struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Date    primitive.DateTime `json:"date" bson:"date"`
	Period  int                `json:"period" bson:"period"`
	Student string             `json:"student" bson:"student"`
	Present bool               `json:"present" bson:"present"`
}

type Teacher struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Class string             `json:"ct" bson:"ct"`
}

type Class struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	ClassTeacher string             `jsom:"classteacher" bson:"classteacher"`
}

// ################################# without id #####################################
type UserWithOutId struct {
	Name     string `json:"name" bson:"name" binding:"required"`
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Type     string `json:"type" bson:"type" binding:"required"`
}

type StudentWithOutId struct {
	Name  string `json:"name" bson:"name" binding:"required"`
	Class string `json:"class" bson:"class" binding:"required"`
}

type AttendanceWithOutId struct {
	Date    primitive.DateTime `json:"date" bson:"date" binding:"required"`
	Period  int                `json:"period" bson:"period" binding:"required"`
	Student string             `json:"student" bson:"student" binding:"required"`
	Present bool               `json:"present" bson:"present"`
}

type TeacherWithOutId struct {
	Name  string `json:"name" bson:"name" binding:"required"`
	Class string `json:"ct" bson:"ct" binding:"required"`
}

type ClassWithOutId struct {
	Name         string `json:"name" bson:"name" binding:"required"`
	ClassTeacher string `jsom:"classteacher" bson:"classteacher" binding:"required"`
}
