package controllers

import (
	"ascii_store/backend/internal/domain/entities"
	"ascii_store/backend/internal/domain/interfaces"
	"ascii_store/backend/internal/logger"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type AsciiHandler struct {
	PORT    string
	service interfaces.AsciiService
}

func NewAsciiHandler(service interfaces.AsciiService) *AsciiHandler {
	PORT := os.Getenv("PORT")

	return &AsciiHandler{PORT: PORT, service: service}
}

func (p *AsciiHandler) StartApi() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"},
		AllowMethods:     []string{"GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/pet", p.getAscii)

	r.DELETE("/pet", p.deleteAscii)

	r.PUT("/pet", p.uploadAscii)

	r.Run(":" + p.PORT)
}

func (p *AsciiHandler) getAscii(ctx *gin.Context) {
	ascii, err := p.service.GetAscii()

	if err != nil {
		logger.Log.Info("Something went wrong", "ERROR", err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ascii":       "",
			"description": "",
		})

		return
	}

	logger.Log.Info("Successfully sent ascii", "DESCRIPTION", ascii.Description)

	ctx.JSON(http.StatusOK, gin.H{
		"ascii":       ascii.Ascii,
		"description": ascii.Description,
	})
}

func (p *AsciiHandler) deleteAscii(ctx *gin.Context) {
	delete_status := p.service.DeleteAscii()

	if delete_status != nil {
		logger.Log.Info("Something went wrong during deleting all ASCIIs", "ERROR", delete_status)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": delete_status.Error(),
		})

		return
	}

	logger.Log.Info("Successfully deleted all ASCIIs")

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Successfully deleted",
	})
}

func (p *AsciiHandler) uploadAscii(ctx *gin.Context) {
	var ascii entities.Ascii
	if err := ctx.ShouldBindJSON(&ascii); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "Give correct body format",
		})

		return
	}

	if ascii.Ascii == "" || ascii.Description == "" {
		logger.Log.Info("Invalid data. Something is not passed")

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Fill all the fields",
		})

		return
	}

	logger.Log.Info("Received", "DESCRIPTION", ascii.Description)

	err := p.service.UploadAscii(ascii)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Something went wrong",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Ascii uploaded successfully",
	})
}
