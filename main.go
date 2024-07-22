package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/uploads", "./uploads")
	router.LoadHTMLGlob("templates/*")

	// Home route
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Upload route
	router.POST("/upload", uploadHandler)

	// Run the server
	router.Run(":8080")
}

func uploadHandler(c *gin.Context) {
	log.Println("Upload handler called")

	// Ensure the `uploads` directory exists
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		err = os.Mkdir("uploads", os.ModePerm)
		if err != nil {
			log.Printf("Failed to create uploads directory: %v", err)
			c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": "Failed to create uploads directory"})
			return
		}
	}

	// Get the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Failed to get uploaded file: %v", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": "Failed to get uploaded file"})
		return
	}
	filePath := "uploads/" + file.Filename
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		log.Printf("Failed to save uploaded file: %v", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": "Failed to save uploaded file"})
		return
	}
	log.Println("File uploaded and saved successfully")

	// Get environment variables
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	apiSecret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		log.Printf("Failed to initialize Cloudinary: %v", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": "Failed to initialize Cloudinary"})
		return
	}

	// Generate current Unix timestamp
	timestamp := time.Now().Unix()

	// Upload file to Cloudinary
	resp, err := cld.Upload.Upload(context.TODO(), filePath, uploader.UploadParams{
		Timestamp:      timestamp,
		Categorization: "google_tagging",
		AutoTagging:    0.5, // Lower confidence threshold for tagging (0 to 1.0)
	})
	if err != nil {
		log.Printf("Failed to upload to Cloudinary: %v", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": "Failed to upload to Cloudinary"})
		return
	}
	log.Println("File uploaded to Cloudinary successfully")
	log.Printf("Cloudinary Response: %+v", resp)

	// Ensure tags are properly formatted as an array
	if len(resp.Tags) == 0 {
		log.Printf("No tags returned from Cloudinary.")
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"error": "No tags returned from Cloudinary"})
		return
	}

	log.Printf("Tags returned from Cloudinary: %+v", resp.Tags)

	// Render the results page
	c.HTML(http.StatusOK, "results.html", gin.H{
		"url":  resp.SecureURL,
		"tags": resp.Tags,
	})
}
