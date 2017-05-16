package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("getenv", "foo")
	cmd.Env = append(os.Environ(), "foo=newbar") // HL
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
