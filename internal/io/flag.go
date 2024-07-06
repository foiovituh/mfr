package io

import (
	"flag"
	"fmt"

	"github.com/foiovituh/mfr/internal/util"
)

type Flags struct {
	DirectoryPath     *string
	PatternToApply    *string
	ExtensionToFilter *string
}

func (f *Flags) Get() {
	f.DirectoryPath = flag.String("d", "",
		"[REQUIRED] Target directory path to rename all files")
	f.PatternToApply = flag.String("p", "file-",
		"Pattern apply on file renaming process")
	f.ExtensionToFilter = flag.String("e", "",
		"Only rename files that contain the extension")
	flag.Parse()
}

func (f Flags) Help() {
	flag.Usage()
}

func (f *Flags) UndeclaredDirectory() bool {
	return *f.DirectoryPath == ""
}

func (f *Flags) EnsureFinalSlashInDirectory() {
	if util.LastCharacter(*f.DirectoryPath) != "/" {
		*f.DirectoryPath = fmt.Sprint(*f.DirectoryPath, "/")
	}
}

func (f *Flags) SkipFileExtension(extensionToCompare string) bool {
	return *f.ExtensionToFilter != "" &&
		*f.ExtensionToFilter != extensionToCompare
}
