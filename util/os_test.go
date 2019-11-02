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

	os.Setenv("qqnumber", "password")

	envs = os.Environ()
	for _, env := range envs {
		log.Println(env)
	}
	os.Unsetenv("qqnumber")
	envs = os.Environ()
	for _, env := range envs {
		log.Println(env)
	}
}

func TestPrintOS(t *testing.T) {
	PrintOS()
}

func TestPrintOsInfo(t *testing.T) {
	PrintOsInfo()
}
