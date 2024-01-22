package main

import (
	"flag"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/skip2/go-qrcode"
	"gopkg.in/go-playground/colors.v1"
	"image/color"
	"os"
	"fmt"
	"net/url"
	"encoding/json"
	"path/filepath"
	"time"
 )

var (
	addr = flag.String("addr", ":3000", "HTTP address to listen on")
)

func main() {
	flag.Parse()

	app := fiber.New()
	app.Use(cors.New())

	// Serve static files (Svelte app)
	app.Static("/", "./frontend/public")

	// API endpoint to generate QR code
	app.Post("/api/generateQR", generateQRCode)

	// Start the server
	err := app.Listen(*addr)
	if err != nil {
		panic(err)
	}
}

func generateQRCode(c *fiber.Ctx) error {

    var requestData map[string]string
    if err := c.BodyParser(&requestData); err != nil {
        return c.Status(http.StatusBadRequest).SendString("Error parsing JSON request")
    }
    
	fileName := generateFileName(requestData)

	// Check if file already exists (this way can reduce response from 8ms to 1ms)
	if len(fileName) < 250 {
		// Check if file already exists (this way can reduce response from 8ms to 1ms)
		if fileExist, err := os.Open("./temp/" + fileName); err == nil {
			return c.SendStream(fileExist)
		}else{
			fileExist.Close()
		}
	}
	
	/*
	if err := deleteOldFiles("./temp/"); err != nil {
		fmt.Println("Error:", err)
	} 
	*/


	// Process QR code generation
	data := requestData["data"]
    fgColor := requestData["fgColor"]
    bgColor := requestData["bgColor"]

	fgColorParsed, err := colors.Parse(fgColor)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid foreground color format")
	}
	fgColorRGBA := fgColorParsed.ToRGBA()

	bgColorParsed, err := colors.Parse(bgColor)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid background color format")
	}
	bgColorRGBA := bgColorParsed.ToRGBA()

	// Convert RGBA to color.RGBA
	fgRGBA := color.RGBA{
		R: uint8(fgColorRGBA.R),
		G: uint8(fgColorRGBA.G),
		B: uint8(fgColorRGBA.B),
		A: uint8(fgColorRGBA.A * 255), // Convert float64 to uint8
	}

	bgRGBA := color.RGBA{
		R: uint8(bgColorRGBA.R),
		G: uint8(bgColorRGBA.G),
		B: uint8(bgColorRGBA.B),
		A: uint8(bgColorRGBA.A * 255), // Convert float64 to uint8
	}

 	err = qrcode.WriteColorFile(data, qrcode.Medium, 256, &fgRGBA, &bgRGBA, fileName)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error generating QR code")
	}

	// Read the file
	file, err := os.Open(fileName)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error reading QR code file")
	}
	
	return c.SendStream(file)
}


func generateFileName(data map[string]string) string {
	jsonString, _ := json.Marshal(data)
	truncatedJSON := string(jsonString)

	if len(truncatedJSON) > 250 {
		truncatedJSON = truncatedJSON[:250]
	}

	return "./temp/" + url.QueryEscape(truncatedJSON) + ".png"
}

func deleteOldFiles(directoryPath string) error {
	// Get the current time
	currentTime := time.Now()

	// Calculate the time threshold (1 hour ago)
	thresholdTime := currentTime.Add(-time.Hour)

	// Walk through the directory
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it is a regular file and older than 1 hour
		if !info.IsDir() && info.ModTime().Before(thresholdTime) {
			fmt.Printf("Deleting old file: %s\n", path)
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking the path %v: %v", directoryPath, err)
	}

	return nil
}
