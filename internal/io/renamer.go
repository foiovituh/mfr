package io

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/foiovituh/mfr/internal/util"
)

func Rename(flags *Flags) {
	files, err := os.ReadDir(*flags.DirectoryPath)
	util.LogFatalIfErrorIsNotNull(err)

	oldPathPerExtension := filterFilesByExtension(files, flags)

	index := 1

	for oldPath, extension := range oldPathPerExtension {
		newPath := filepath.Join(*flags.DirectoryPath,
			*flags.PatternToApply+strconv.Itoa(index)+extension)

		err := os.Rename(oldPath, newPath)
		util.LogFatalIfErrorIsNotNull(err)

		fmt.Printf("%s -> %s\n", oldPath, newPath)

		index++
	}
}

func filterFilesByExtension(files []os.DirEntry,
	flags *Flags) map[string]string {
	filteredFilesByExtension := make(map[string]string)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		oldPath := filepath.Join(*flags.DirectoryPath, file.Name())
		extension := filepath.Ext(oldPath)

		if flags.SkipFileExtension(extension) {
			continue
		}

		filteredFilesByExtension[oldPath] = extension
	}

	return filteredFilesByExtension
}
