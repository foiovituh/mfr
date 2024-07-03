package main

import (
	"os"

	"github.com/foiovituh/mfr/internal/io"
	"github.com/foiovituh/mfr/internal/static"
	"github.com/foiovituh/mfr/internal/util"
)

func main() {
	missingArguments := len(os.Args) < 3

	util.LogFatalIfTrue(&missingArguments, static.SpecifyArguments)

	directoryPath := &os.Args[1]
	newFilePath := &os.Args[2]

	if util.LastCharacter(*directoryPath) != static.Slash {
		*directoryPath = *directoryPath + static.Slash

	}

	io.Rename(directoryPath, newFilePath)
}
