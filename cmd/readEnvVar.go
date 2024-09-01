package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func readEnvVar(varName string) string {
	return os.Getenv(varName)
}

func readAndCheckEnvVar(varName string) (string, error) {
	varVal := readEnvVar(varName)
	if varVal == "" {
		return "", errors.New(fmt.Sprintf("%s environment variable not set properly", varName))
	}

	return varVal, nil
}

func readAndCheckIntEnvVar(varName string) (string, error) {
	varVal, err := readAndCheckEnvVar(varName)
	if err != nil {
		return "", err
	}

	if _, err = strconv.Atoi(varVal); err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("value of %s environment variable has to be an integer", varName))
	}

	return varVal, nil
}
