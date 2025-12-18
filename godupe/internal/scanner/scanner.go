package scanner

import (
	"godupe/internal/models"
	"godupe/internal/utils"
	"io/fs"
	"path/filepath"
)

func FindDuplicates(rootPath string) (map[string][]string, error) {

	// Step 1: Traverse the directory and group files by size
	sizeMap := make(map[int64][]models.FileInfo)
	err := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			info, _ := d.Info()
			fileSize := info.Size()

			newFileInfo := models.FileInfo{
				Path: path,
				Size: fileSize,
			}

			sizeMap[fileSize] = append(sizeMap[fileSize], newFileInfo)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	
	// Step 2: For files with the same size, compute hashes and group by hash
	duplicates := make(map[string][]string)

	for _, files := range sizeMap {
		if len(files) >1 {
			for _, file := range files {
				hash, err := utils.CalculateHash(file.Path)
				if err != nil {
					return nil, err
				}
				duplicates[hash] = append(duplicates[hash], file.Path)
			}
		}
	}

	// Step 3: Filter out groups with only one file
	finalResults := make(map[string][]string)
	for hash, paths := range duplicates {
		if len(paths) > 1 {
			finalResults[hash] = paths
		}
	}

	return finalResults, nil
}
