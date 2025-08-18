package image

import (
	"fmt"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"svc-llt-golang/utils/config"

	"github.com/nfnt/resize"
)

// ProcessImageUpload handles image upload with compression and saves with lansia UUID
func ProcessImageUpload(file *multipart.FileHeader, uuid string) (string, error) {
	// Get storage configuration from environment
	storageConfig := config.GetStorageConfig()
	
	// Create images directory if not exists (using config path)
	imageDir := storageConfig.GetImageStoragePath()
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create images directory: %v", err)
	}

	// Validate image file with configurable size limit
	if err := validateImageFile(file, storageConfig); err != nil {
		return "", err
	}

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// Generate filename using lansia UUID (same pattern as QR code)
	filename := fmt.Sprintf("%s.jpg", uuid)
	filepath := filepath.Join(imageDir, filename)

	// Decode image
	img, _, err := image.Decode(src)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %v", err)
	}

	// Resize image using configurable dimensions
	resizedImg := resizeImage(img, uint(storageConfig.MaxWidth), uint(storageConfig.MaxHeight))

	// Create output file
	dst, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create output file: %v", err)
	}
	defer dst.Close()

	// Save as JPEG with configurable quality
	if err := jpeg.Encode(dst, resizedImg, &jpeg.Options{Quality: storageConfig.Quality}); err != nil {
		return "", fmt.Errorf("failed to encode image: %v", err)
	}

	return filepath, nil
}

// validateImageFile validates uploaded image file
func validateImageFile(file *multipart.FileHeader, storageConfig config.StorageConfig) error {
	// Check file size using configurable limit
	maxSize := storageConfig.GetMaxSizeBytes()
	if file.Size > maxSize {
		return fmt.Errorf("image file too large: %d bytes (max: %d bytes)", file.Size, maxSize)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	validExts := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp"}
	
	isValidExt := false
	for _, validExt := range validExts {
		if ext == validExt {
			isValidExt = true
			break
		}
	}
	
	if !isValidExt {
		return fmt.Errorf("invalid image type: %s. Allowed types: jpg, jpeg, png, gif, bmp", ext)
	}

	// Validate if file is actually an image
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	_, format, err := image.DecodeConfig(src)
	if err != nil {
		return fmt.Errorf("invalid image file: %v", err)
	}

	// Check supported formats
	supportedFormats := []string{"jpeg", "jpg", "png", "gif", "bmp"}
	isSupported := false
	for _, supportedFormat := range supportedFormats {
		if format == supportedFormat {
			isSupported = true
			break
		}
	}

	if !isSupported {
		return fmt.Errorf("unsupported image format: %s", format)
	}

	return nil
}

// resizeImage resizes image if it exceeds maximum dimensions
func resizeImage(img image.Image, maxWidth, maxHeight uint) image.Image {
	bounds := img.Bounds()
	width := uint(bounds.Dx())
	height := uint(bounds.Dy())

	// Check if resize is needed
	if width <= maxWidth && height <= maxHeight {
		return img
	}

	// Calculate new dimensions maintaining aspect ratio
	var newWidth, newHeight uint

	if width > height {
		// Landscape orientation
		newWidth = maxWidth
		newHeight = 0 // Let resize calculate to maintain aspect ratio
	} else {
		// Portrait orientation
		newWidth = 0 // Let resize calculate to maintain aspect ratio
		newHeight = maxHeight
	}

	// Resize using Lanczos resampling for best quality
	return resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
}

// DeleteImage removes image file from storage
func DeleteImage(filepath string) error {
	if filepath == "" {
		return nil // No file to delete
	}
	
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil // File doesn't exist, nothing to delete
	}
	
	return os.Remove(filepath)
}