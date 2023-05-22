package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	process := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	for key, value := range env {
		if _, ok := os.LookupEnv(key); ok {
			os.Unsetenv(key)
			if !value.NeedRemove {
				os.Setenv(key, value.Value)
			}
		} else {
			os.Setenv(key, value.Value)
		}
	}

	if err := process.Start(); err != nil {
		return 1
	}
	if err := process.Wait(); err != nil {
		if err, ok := err.(*exec.ExitError); ok { //nolint:errorlint
			return err.ExitCode()
		}
	}
	return 0
}
