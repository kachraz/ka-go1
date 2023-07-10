// Now to try git with this

package main

import (
	"os"
	"os/exec"
	"syscall"
)

// Context from Function try1/main.go:main.main

// This function is to run exa via the command line using syscall.Exec

func main() {
	
	var address = "https://github.com/liqowomo/lq-wr1"
	
	binary, lookErr := exec.LookPath("git")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"git", "clone", address}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
