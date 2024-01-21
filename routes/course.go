package routes

import (
	"institute/config"
	"institute/features/course"
	"institute/helpers"

	m "institute/middlewares"

	"github.com/labstack/echo/v4"
)

func Courses(e *echo.Echo, handler course.Handler, jwt helpers.JWTInterface, config config.ProgramConfig) {
	courses := e.Group("/courses")

	courses.GET("", handler.GetCourses(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	courses.POST("", handler.CreateCourse(), m.AuthorizeJWT(jwt, 2, config.SECRET))
	
	courses.GET("/:id", handler.CourseDetails(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	courses.PUT("/:id", handler.UpdateCourse(), m.AuthorizeJWT(jwt, 3, config.SECRET))
	courses.DELETE("/:id", handler.DeleteCourse(), m.AuthorizeJWT(jwt, 2, config.SECRET))
}