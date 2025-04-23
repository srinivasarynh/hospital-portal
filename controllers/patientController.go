package controllers

import (
	"hospital-portal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient models.Patient
		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&patient).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register patient"})
			return
		}

		c.JSON(http.StatusCreated, patient)
	}
}

func GetPatients(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var patients []models.Patient

		if err := db.Find(&patients).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
			return
		}

		c.JSON(http.StatusOK, patients)
	}
}

func GetPatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var patient models.Patient
		if err := db.First(&patient, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		c.JSON(http.StatusOK, patient)
	}
}

func UpdatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var patient models.Patient
		if err := db.First(&patient, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		if err := c.BindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&patient)
		c.JSON(http.StatusOK, patient)
	}
}

func DeletePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var patient models.Patient
		if err := db.First(&patient, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}

		db.Delete(&patient)
		c.JSON(http.StatusOK, gin.H{"message": "Patient deleted successfully"})
	}
}
