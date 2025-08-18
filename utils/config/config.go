package config

import (
	"os"
	"path/filepath"
	"strconv"
)

// StorageConfig holds configuration for file storage (images and QR codes)
type StorageConfig struct {
	StoragePath string // Base storage path from env
	MaxSizeMB   int    // Max file size in MB
	MaxWidth    int    // Max image width in pixels
	MaxHeight   int    // Max image height in pixels
	Quality     int    // JPEG quality percentage
}

// GetStorageConfig returns storage configuration from environment variables
func GetStorageConfig() StorageConfig {
	config := StorageConfig{
		// Default values - sama dengan current development setup
		StoragePath: "storage", // Default ke ./storage (current working directory)
		MaxSizeMB:   2,
		MaxWidth:    600,
		MaxHeight:   400,
		Quality:     70,
	}

	// Get storage path from env
	if storagePath := os.Getenv("STORAGE_PATH"); storagePath != "" {
		config.StoragePath = storagePath
	}

	// Get max file size from env
	if maxSizeStr := os.Getenv("IMAGE_MAX_SIZE_MB"); maxSizeStr != "" {
		if maxSize, err := strconv.Atoi(maxSizeStr); err == nil && maxSize > 0 {
			config.MaxSizeMB = maxSize
		}
	}

	// Get max width from env
	if maxWidthStr := os.Getenv("IMAGE_MAX_WIDTH"); maxWidthStr != "" {
		if maxWidth, err := strconv.Atoi(maxWidthStr); err == nil && maxWidth > 0 {
			config.MaxWidth = maxWidth
		}
	}

	// Get max height from env
	if maxHeightStr := os.Getenv("IMAGE_MAX_HEIGHT"); maxHeightStr != "" {
		if maxHeight, err := strconv.Atoi(maxHeightStr); err == nil && maxHeight > 0 {
			config.MaxHeight = maxHeight
		}
	}

	// Get quality from env
	if qualityStr := os.Getenv("IMAGE_QUALITY"); qualityStr != "" {
		if quality, err := strconv.Atoi(qualityStr); err == nil && quality > 0 && quality <= 100 {
			config.Quality = quality
		}
	}

	return config
}

// GetImageStoragePath returns the full path for image storage
func (c *StorageConfig) GetImageStoragePath() string {
	return filepath.Join(c.StoragePath, "images")
}

// GetQRStoragePath returns the full path for QR code storage
func (c *StorageConfig) GetQRStoragePath() string {
	return filepath.Join(c.StoragePath, "qrcodes")
}

// GetMaxSizeBytes returns max file size in bytes
func (c *StorageConfig) GetMaxSizeBytes() int64 {
	return int64(c.MaxSizeMB * 1024 * 1024)
}