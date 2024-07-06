package io

import (
	"testing"
)

func TestUndeclaredDirectory(t *testing.T) {
	test := ""
	af := ArgumentFlag{DirectoryPath: &test}

	if !af.UndeclaredDirectory() {
		t.Errorf("UndeclaredDirectory() -> result: false; expected: true")
	}
}

func TestEnsureFinalSlashInDirectory(t *testing.T) {
	test := "test"
	expected := "test/"
	af := ArgumentFlag{DirectoryPath: &test}

	af.EnsureFinalSlashInDirectory()

	if *af.DirectoryPath != expected {
		t.Errorf("EnsureFinalSlashInDirectory() -> result: %s; expected: %s",
			*af.DirectoryPath, expected)
	}
}

func TestSkipFileExtensionWhenExtensionToCompareIsAny(t *testing.T) {
	extensionToCompare := ".txt"
	af := ArgumentFlag{ExtensionToFilter: &extensionToCompare}

	if af.SkipFileExtension(extensionToCompare) {
		t.Errorf("SkipFileExtension() with matching extension -> "+
			"result: %v; expected: %v", true, false)
	}
}

func TestSkipFileExtensionWhenExtensionToFilterIsEmpty(t *testing.T) {
	extensionToFilter := ""
	extensionToCompare := ".txt"
	af := ArgumentFlag{ExtensionToFilter: &extensionToFilter}

	if af.SkipFileExtension(extensionToCompare) {
		t.Errorf("SkipFileExtension() with empty filter -> "+
			"result: %v; expected: %v", true, false)
	}
}

func TestSkipFileExtensionWhenExtensionsDoNotMatch(t *testing.T) {
	extensionToFilter := ".jpg"
	extensionToCompare := ".txt"
	af := ArgumentFlag{ExtensionToFilter: &extensionToFilter}

	if !af.SkipFileExtension(extensionToCompare) {
		t.Errorf("SkipFileExtension() with non-matching extension -> "+
			"result: %v; expected: %v", false, true)
	}
}
