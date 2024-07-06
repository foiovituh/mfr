package io

import (
	"flag"
	"fmt"
	"os"

	"github.com/foiovituh/mfr/internal/util"
)

type ArgumentFlag struct {
	DirectoryPath     *string
	PatternToApply    *string
	ExtensionToFilter *string
}

func (af *ArgumentFlag) Get() {
	af.DirectoryPath = flag.String("d", "",
		"[REQUIRED] Full directory path to rename all files")
	af.PatternToApply = flag.String("p", "file-",
		"[OPTIONAL] Pattern to apply in the file renaming process")
	af.ExtensionToFilter = flag.String("e", "",
		"[OPTIONAL] Filter files containing the extension")
	flag.Parse()
}

func (af ArgumentFlag) Help() {
	flag.Usage()
	fmt.Println("\nSee the examples in the README.md file or in the official " +
		"repository: \n  https://github.com/foiovituh/mfr")
	os.Exit(0)
}

func (af *ArgumentFlag) UndeclaredDirectory() bool {
	return *af.DirectoryPath == ""
}

func (af *ArgumentFlag) EnsureFinalSlashInDirectory() {
	if util.LastCharacter(*af.DirectoryPath) != "/" {
		*af.DirectoryPath = fmt.Sprint(*af.DirectoryPath, "/")
	}
}

func (af ArgumentFlag) SkipFileExtension(extensionToCompare string) bool {
	return *af.ExtensionToFilter != "" &&
		*af.ExtensionToFilter != extensionToCompare
}
