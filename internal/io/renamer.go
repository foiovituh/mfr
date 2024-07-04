package io

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/foiovituh/mfr/internal/util"
)

func Rename(directoryPath, newFilesPrefix string) {
	files, err := os.ReadDir(directoryPath)

	util.LogFatalIfErrorIsNotNull(err)

	for index, file := range files {
		if file.IsDir() {
			continue
		}

		oldPath := filepath.Join(directoryPath, file.Name())
		newPath := directoryPath + newFilesPrefix + strconv.Itoa(index) +
			filepath.Ext(oldPath)

		os.Rename(oldPath, newPath)

		fmt.Printf("%s -> %s\n", oldPath, newPath)
	}
}
