package utils

import (
	"log"
	"testing"
)

func TestNextId(t *testing.T) {
	InitNode(2)
	for {
		log.Println(NextId())
	}
}

func TestUuidText(t *testing.T) {
	for {
		log.Println(UuidText())
	}
}
