// +build OMIT

package main

import (
	"log"
	"os/exec"
)

func main() {
	err := exec.Command("mkdir", "/tmp/foo").Run()
	if err != nil {
		log.Fatal(err)
	}

	err = exec.Command("rm", "-rf", "/tmp/foo").Run()
	if err != nil {
		log.Fatal(err)
	}
}
