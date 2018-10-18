package test

import (
	"log"
	"os"
	"testing"
)

func baseDir(t *testing.T) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
