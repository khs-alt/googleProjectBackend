package handler

import (
	"backend/app/models"
	"backend/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/sqltocsv"
)

func ExportImageDataHandler(c *gin.Context) {
	w := c.Writer
	var requestData models.TestCodeData

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	testcode := requestData.TestCode
	rows, err := sql.ExportImageData(testcode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sqltocsv.Write(w, rows)
}

func ExportVideoDataHandler(c *gin.Context) {
	w := c.Writer
	var requestData models.TestCodeData

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	testcode := requestData.TestCode
	rows, err := sql.ExportVideoData(testcode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sqltocsv.Write(w, rows)
}
