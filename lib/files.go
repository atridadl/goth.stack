package lib

import (
	"os"
	"path/filepath"
)

// ListFiles returns a slice of file paths in the given directory
func ListFiles(dir string) ([]string, error) {
	var filePaths []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			// Recursively list files in subdirectories
			subDir := filepath.Join(dir, entry.Name())
			subFiles, err := ListFiles(subDir)
			if err != nil {
				return nil, err
			}
			filePaths = append(filePaths, subFiles...)
		} else {
			// Add file path to the slice
			filePath := filepath.Join(dir, entry.Name())
			filePaths = append(filePaths, filePath)
		}
	}
	return filePaths, nil
}
