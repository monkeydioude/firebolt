package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	bolt "github.com/coreos/bbolt"
)

type entry map[string]string
type bucket map[string]entry

func retrieveBuckets(p string) (*bucket, error) {
	sourceFile, err := ioutil.ReadFile(p)
	var buckets *bucket

	if err != nil {
		return nil, fmt.Errorf("[ERR ] %s", err)
	}

	if err := json.Unmarshal(sourceFile, &buckets); err != nil {
		return buckets, fmt.Errorf("[ERR ] %s", err)
	}

	return buckets, nil
}

func fillBucket(db *bolt.DB, b string, c entry) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(b))

		if err != nil {
			return fmt.Errorf("[WARN] %s", err)
		}

		for k, v := range c {
			b.Put([]byte(k), []byte(v))
		}
		return nil
	})

	return err
}

func burnBuckets(b *bucket, p string) error {
	db, err := bolt.Open(p, 0666, nil)
	defer db.Close()

	if err != nil {
		return fmt.Errorf("[ERR ] %s", err)
	}

	for k, v := range *b {
		if err := fillBucket(db, k, v); err != nil {
			log.Printf("[WARN] On filling bucket '%s': %s", k, err)
		}
		log.Printf("[INFO] Successfuly filled bucket '%s'", k)
	}

	return nil
}
