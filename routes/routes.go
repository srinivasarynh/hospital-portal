package routes

import (
	"hospital-portal/controllers"
	"hospital-portal/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(incommingRoutes *gin.Engine, db *gorm.DB) {
	auth := incommingRoutes.Group("/auth")
	auth.POST("/login", controllers.Login(db))

	reception := incommingRoutes.Group("/reception")
	reception.Use(middlewares.AuthMiddleware("receptionist"))
	{
		reception.POST("/patients", controllers.CreatePatient(db))
		reception.GET("/patients", controllers.GetPatients(db))
		reception.GET("/patients/:id", controllers.GetPatient(db))
		reception.PUT("patients/:id", controllers.UpdatePatient(db))
		reception.DELETE("/patients/:id", controllers.DeletePatient(db))
	}

	doctor := incommingRoutes.Group("/doctor")
	doctor.Use(middlewares.AuthMiddleware("doctor"))
	{
		doctor.GET("/patients", controllers.GetPatients(db))
		doctor.GET("/patients/:id", controllers.GetPatient(db))
		doctor.PUT("/patients/:id", controllers.UpdatePatient(db))
		doctor.PATCH("/patients/:id/notes", controllers.DoctorUpdatePatientNote(db))
	}
}
