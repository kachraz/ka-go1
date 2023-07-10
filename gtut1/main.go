// This file is for learning how to run a system proces with go 
// Now you can try other stuff in other files 





package main

import (
	"os"
	"os/exec"
	"syscall"
)

// Context from Function try1/main.go:main.main

// This function is to run exa via the command line using syscall.Exec

func main() {
	
	
	binary, lookErr := exec.LookPath("exa")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"exa", "-l", "-a", "--icons"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}
