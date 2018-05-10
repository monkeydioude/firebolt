package main

import (
	"testing"
)

func TestIFailWhenFileIsNotReadable(t *testing.T) {
	_, err := retrieveBuckets("testdata/unknown")
	if err == nil {
		t.Error("Should have failed on ReadFile")
	}
}

func TestIFailWhenUnmarshalingWrongJson(t *testing.T) {
	buckets, err := retrieveBuckets("testdata/1-buckets-wrong.json")

	if !(err != nil && buckets != nil) {
		t.Error("Should have failed on Unmarshal")
	}
}

func TestUnmarshalingCorrectlyFormattedJson(t *testing.T) {
	buckets, err := retrieveBuckets("testdata/2-buckets-ok.json")

	if !(err == nil && buckets != nil) {
		t.Error("Should have not failed")
	}

}
