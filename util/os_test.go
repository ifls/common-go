package util

import (
	"log"
	"os"
	"testing"
)

func TestENV(t *testing.T) {
	envs := os.Environ()
	for _, env := range envs {
		log.Println(env)
	}

	err := os.Setenv("qqnumber", "password")
	if err != nil {
		t.Fatal(err)
	}
	envs = os.Environ()
	for _, env := range envs {
		log.Println(env)
	}
	err = os.Unsetenv("qqnumber")
	if err != nil {
		t.Fatal(err)
	}
	envs = os.Environ()
	for _, env := range envs {
		log.Println(env)
	}
}

func TestGetTid(t *testing.T) {
	GetTid("ss")
}

func TestPrintOS(t *testing.T) {
	PrintOS()
}

func TestPrintOsInfo(t *testing.T) {
	PrintOsInfo()
}
