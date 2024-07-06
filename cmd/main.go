package main

import (
	"github.com/foiovituh/mfr/internal/io"
)

func main() {
	flags := io.Flags{}
	flags.Get()

	if flags.UndeclaredDirectory() {
		flags.Help()
	}

	flags.EnsureFinalSlashInDirectory()

	io.Rename(&flags)
}
