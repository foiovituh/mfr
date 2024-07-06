package main

import (
	"github.com/foiovituh/mfr/internal/io"
)

func main() {
	argumentFlag := io.ArgumentFlag{}
	argumentFlag.Get()

	if argumentFlag.UndeclaredDirectory() {
		argumentFlag.Help()
	}

	argumentFlag.EnsureFinalSlashInDirectory()

	io.Rename(&argumentFlag)
}
