// The env package provides functionality to manage environment variables within a Go application. It offers methods
// to retrieve environment variable values from the operating system and seamlessly pass these values to other service components.
package env

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// The env string struct facilitates the retrieval of environment variables from the operating system and transfers
// their values to other service components.
type envVarName string

// type env string
type env struct {
	envName  envVarName
	envValue string
	required bool
}

// GetEnv retrieves the value of the specified environment variable and prints it.
// It takes an `envVarName` as input, representing the name of the environment variable,
// and utilizes its getValue method to fetch the value.
// If an error occurs during the retrieval process and the environment variable is required,
// it logs the error as a fatal error and terminates the program; otherwise, it logs the error as a warning.
func GetEnv(envVar envVarName, options ...Options) string {
	// Create new env Instance with the envVarName and with default required value
	env := &env{
		envName:  envVar,
		required: true,
	}

	// Get the environment variable value
	err := env.getValue()

	// Iterate over Options
	for _, option := range options {
		option(env)
	}

	// Exit the program if the environment variable is required to be exist with non empty value
	if err != nil && env.required {
		logrus.Fatal(err)
	} else {
		logrus.Warn(err)
	}

	return env.envValue
}

// getValue retrieves the value of the environment variable represented by 'env'.
// It converts 'env' to a string and uses it to fetch the corresponding value from the OS environment.
// If the value is empty or the environment variable does not exist, it returns an error with a descriptive message.
// Otherwise, it sets the value and returns a nil error.
func (e *env) getValue() error {
	e.envValue = os.Getenv(string(e.envName))
	if len(e.envValue) == 0 {
		return fmt.Errorf("enviornment variable '%s' is empty or not exist", e.envName)
	}
	return nil
}
