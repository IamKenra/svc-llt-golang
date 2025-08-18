package qrcode

import (
	"fmt"
	"os"
	"path/filepath"
	"svc-llt-golang/utils/config"

	"github.com/skip2/go-qrcode"
)

// GenerateQRCode generates QR code from UUID and saves it to file
func GenerateQRCode(uuid string) (string, error) {
	// Get storage configuration from environment
	storageConfig := config.GetStorageConfig()
	
	// Create QR codes directory if not exists (using config path)
	qrDir := storageConfig.GetQRStoragePath()
	if err := os.MkdirAll(qrDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create QR codes directory: %v", err)
	}

	// Generate filename
	filename := fmt.Sprintf("%s.png", uuid)
	filepath := filepath.Join(qrDir, filename)

	// Generate QR code with UUID content
	err := qrcode.WriteFile(uuid, qrcode.Medium, 256, filepath)
	if err != nil {
		return "", fmt.Errorf("failed to generate QR code: %v", err)
	}

	return filepath, nil
}

// DeleteQRCode deletes QR code file
func DeleteQRCode(filepath string) error {
	if filepath == "" {
		return nil // No file to delete
	}
	
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return nil // File doesn't exist, nothing to delete
	}
	
	return os.Remove(filepath)
}