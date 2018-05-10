package main

import (
	"os"
	"testing"
)

func TestDestFileExistsAndIsWritable(t *testing.T) {
	err := handleDestinationFile("testdata/readable-writable-file.json")
	if err != nil {
		t.Error("File should exist and be readable")
	}
}

func TestIFailOnNotWritableDestFile(t *testing.T) {
	file := "testdata/not-writable-file.json"
	// Read only
	os.Chmod(file, 0466)

	err := handleDestinationFile(file)
	if err == nil {
		t.Error("File should exist and should not be writable")
	}

	os.Chmod(file, 0666)
}

func TestIJustReturnWhenDestFileDoesNotExist(t *testing.T) {
	err := handleDestinationFile("testdata/unknownsource")
	if err != nil {
		t.Error("File should not exist")
	}
}

func TestIFailOnNotExistingSourceFile(t *testing.T) {
	err := handleSourceFile("testdata/unknownsource")
	if err == nil {
		t.Error("File should not exist")
	}
}

func TestSourceFileExistsAndIsReadable(t *testing.T) {
	err := handleSourceFile("testdata/source-1.json")
	if err != nil {
		t.Error("File should exist")
	}
}

func TestIFailOnNotReadableFile(t *testing.T) {
	file := "testdata/not-readable-file.json"
	// Write only
	os.Chmod(file, 0266)

	err := handleSourceFile(file)
	if err == nil {
		t.Error("File should exist and should not be writable")
	}

	os.Chmod(file, 0666)
}
