package io

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/foiovituh/mfr/internal/util"
)

func Rename(argumentFlag *ArgumentFlag) {
	files, err := os.ReadDir(*argumentFlag.DirectoryPath)
	util.LogFatalIfError(err)

	oldPathPerExtension := filterFilesByExtension(files, argumentFlag)

	index := 1

	for oldPath, extension := range oldPathPerExtension {
		newPath := filepath.Join(*argumentFlag.DirectoryPath,
			*argumentFlag.PatternToApply+strconv.Itoa(index)+extension)

		err := os.Rename(oldPath, newPath)
		util.LogFatalIfError(err)

		fmt.Printf("%s -> %s\n", oldPath, newPath)

		index++
	}
}

func filterFilesByExtension(files []os.DirEntry,
	argumentFlag *ArgumentFlag) map[string]string {
	filteredFilesByExtension := make(map[string]string)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		oldPath := filepath.Join(*argumentFlag.DirectoryPath, file.Name())
		extension := filepath.Ext(oldPath)

		if argumentFlag.SkipFileExtension(extension) {
			continue
		}

		filteredFilesByExtension[oldPath] = extension
	}

	return filteredFilesByExtension
}
