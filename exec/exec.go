package exec

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func GetEnvVarOrDefault(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

func RunCommandWithEnv(cmdString string, logger *log.Logger) error {
	command := exec.Command("bash", "-c", cmdString)

	command.Env = os.Environ()

	stdout, err := command.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %v", err)
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to get stderr pipe: %v", err)
	}

	if err := command.Start(); err != nil {
		return fmt.Errorf("failed to start command: %v", err)
	}

	go func() {
		_, err := io.Copy(os.Stdout, stdout)
		if err != nil {
			logger.Printf("Failed to copy stdout: %v\n", err)
		}
	}()

	go func() {
		_, err := io.Copy(os.Stderr, stderr)
		if err != nil {
			logger.Printf("Failed to copy stderr: %v\n", err)
		}
	}()

	logger.Printf("Executing: %s\n", cmdString)

	if err := command.Wait(); err != nil {
		return fmt.Errorf("command execution failed: %v", err)
	}

	logger.Printf("Finished executing: %s\n", cmdString)

	return nil
}
