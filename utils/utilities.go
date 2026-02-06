// Package utils provides generic utilities for the program
package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Map(vs []string, f func(string) string) []string {
	// Convenient function from https://gobyexample.com/collection-functions
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func CloseFileWithLog(o *os.File) {
	// makes the linter happy
	err := o.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteConfig(output *os.File, data string) {
	_, err := io.WriteString(output, data)
	if err != nil {
		panic(err)
	}
}

func GetOutputFile(outputPath string) (*os.File, error) {
	switch outputPath {
	case "":
		return nil, errors.New("empty output path")
	case "-":
		return os.Stdout, nil
	default:
		return os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o640)
	}
}

func ShowActiveConfig() error {
	activeConfiguration := os.Getenv("APEUPRES_NAME")
	if activeConfiguration != "" {
		fmt.Println("Active configuration : ", activeConfiguration)
		activeEnvVariables := os.Getenv("APEUPRES_TO_CLEAN_ENV")
		for v := range strings.SplitSeq(activeEnvVariables, ":") {
			fmt.Printf("%s: %s\n", v, os.Getenv(v))
		}
	} else {
		fmt.Println("No active configuration")
	}

	return nil
}
