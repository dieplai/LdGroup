package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"perfume-quiz-backend/models"
	"perfume-quiz-backend/utils"
)

// Get all questions
func GetQuestions(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set UTF-8 encoding header
		c.Header("Content-Type", "application/json; charset=utf-8")

		var questions []models.Question
		if err := db.Find(&questions).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch questions"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"questions": questions,
		})
	}
}

// Submit quiz and calculate result
func SubmitQuiz(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set UTF-8 encoding header
		c.Header("Content-Type", "application/json; charset=utf-8")

		var submission models.QuizSubmission
		if err := c.ShouldBindJSON(&submission); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validate answers length
		if len(submission.Answers) != 5 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Must provide exactly 5 answers"})
			return
		}

		// Calculate result
		result, scores, description := utils.CalculateResult(submission.Answers, submission.Gender)

		// Save to database
		customerResult := models.CustomerResult{
			Name:        submission.Name,
			Phone:       submission.Phone,
			Gender:      submission.Gender,
			Answers:     strings.Join(submission.Answers, ""),
			FinalResult: result,
			TraitScores: utils.ScoresToJSON(scores),
		}

		if err := db.Create(&customerResult).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save result"})
			return
		}

		// Return result
		c.JSON(http.StatusOK, gin.H{
			"result":      result,
			"description": description,
			"scores":      scores,
			"id":          customerResult.ID,
		})
	}
}

// Admin: Get all results
func GetAllResults(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var results []models.CustomerResult
		if err := db.Order("created_at DESC").Find(&results).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch results"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"results": results,
		})
	}
}

// Admin: Get statistics
func GetStats(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var stats struct {
			Total      int64 `json:"total"`
			Maestro    int64 `json:"maestro"`
			Sculptor   int64 `json:"sculptor"`
			Poet       int64 `json:"poet"`
			Ballerina  int64 `json:"ballerina"`
			Artist     int64 `json:"artist"`
		}

		db.Model(&models.CustomerResult{}).Count(&stats.Total)
		db.Model(&models.CustomerResult{}).Where("final_result = ?", "The Maestro").Count(&stats.Maestro)
		db.Model(&models.CustomerResult{}).Where("final_result = ?", "The Sculptor").Count(&stats.Sculptor)
		db.Model(&models.CustomerResult{}).Where("final_result = ?", "The Poet").Count(&stats.Poet)
		db.Model(&models.CustomerResult{}).Where("final_result = ?", "The Ballerina").Count(&stats.Ballerina)
		db.Model(&models.CustomerResult{}).Where("final_result = ?", "The Artist").Count(&stats.Artist)

		c.JSON(http.StatusOK, stats)
	}
}

// Admin: Delete result
func DeleteResult(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		if err := db.Delete(&models.CustomerResult{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete result"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Result deleted successfully"})
	}
}