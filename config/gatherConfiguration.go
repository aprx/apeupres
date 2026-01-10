// Package config store all configuration related utilities
package config

import (
	"os"
	"path/filepath"
)

func GatherConfiguration(configPath string) []string {
	configurationPaths := []string{}
	info, err := os.Stat(configPath)
	if err != nil {
		panic(err)
	}
	if info.IsDir() {
		directory, err := os.ReadDir(configPath)
		if err != nil {
			panic(err)
		}
		for _, file := range directory {
			if file.Type().IsRegular() {
				configurationPaths = append(configurationPaths, filepath.Join(configPath, file.Name()))
			}
		}
	} else {
		configurationPaths = append(configurationPaths, configPath)
	}
	return configurationPaths
}
