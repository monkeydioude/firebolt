package main

import (
	"errors"
	"log"
	"os"
)

// Checks that source file exists
func handleSourceFile(p string) error {
	f, _ := os.Stat(p)

	if f == nil {
		return errors.New("[ERR ] Source file should exists")
	}

	if f.Mode()&0400 == 0 {
		return errors.New("[ERR ] Source file can not be read")
	}

	return nil
}

// Verify a file exists and is readable
func handleDestinationFile(p string) error {
	f, _ := os.Stat(p)

	if f == nil {
		return nil
	}

	if f.Mode()&0200 == 0 {
		return errors.New("[ERR ] Database file exists but can not be written")
	}

	return nil
}

func handleArgs() (string, string, error) {
	if len(os.Args) <= 2 {
		if len(os.Args) == 2 && os.Args[1] == "help" {
			showHelp()
			os.Exit(1)
		}
		log.Fatal("[ERR ] Firebolt takes 2 arguments. See 'firebolt help'")
	}

	if err := handleSourceFile(os.Args[1]); err != nil {
		return "", "", err
	}

	return os.Args[1], os.Args[2], handleDestinationFile(os.Args[2])
}
