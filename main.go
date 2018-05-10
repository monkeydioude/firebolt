package main

import (
	"fmt"
	"log"
)

func showHelp() {
	fmt.Print("Firebolt takes 2 arguments:\n")
	fmt.Print("\t1/ Path to the formatted json\n")
	fmt.Print("\t2/ Path to the DB file to put data to, created if does not exist\n")
}

func main() {
	sourceFilePath, dbFilePath, err := handleArgs()

	if err != nil {
		log.Fatalf("[ERR ] %s", err)
	}

	buckets, err := retrieveBuckets(sourceFilePath)
	if err != nil {
		log.Fatal(err)
	}

	if err := burnBuckets(buckets, dbFilePath); err != nil {
		log.Fatal(err)
	}

	log.Println("[INFO] Job done")
}
