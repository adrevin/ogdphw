package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	process := exec.Command(cmd[0], cmd[1], cmd[2], cmd[3])

	process.Stdin = os.Stdin
	process.Stdout = os.Stdout
	process.Stderr = os.Stderr

	// remove quotes
	/*for _, env := range os.Environ() {
		entry := strings.SplitN(env, "=", 2)
		os.Setenv(entry[0], strings.Trim(entry[1], "\""))
	}*/

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
		if err, ok := err.(*exec.ExitError); ok {
			return err.ExitCode()
		} else {
			return 1
		}
	}
	return 1
}
