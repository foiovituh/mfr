package io

import (
	"flag"
	"fmt"

	"github.com/foiovituh/mfr/internal/util"
)

type ArgumentFlag struct {
	DirectoryPath     *string
	PatternToApply    *string
	ExtensionToFilter *string
}

func (af *ArgumentFlag) Get() {
	af.DirectoryPath = flag.String("d", "",
		"[REQUIRED] Target directory path to rename all files")
	af.PatternToApply = flag.String("p", "file-",
		"Pattern apply on file renaming process")
	af.ExtensionToFilter = flag.String("e", "",
		"Only rename files that contain the extension")
	flag.Parse()
}

func (af ArgumentFlag) Help() {
	flag.Usage()
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
