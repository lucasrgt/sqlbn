package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	QueryDir  string `yaml:"queryDir"`
	OutputDir string `yaml:"outputDir"`
}

func main() {
	// Get the current working directory (project directory)
	execDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		os.Exit(1)
	}

	// Define the default names for the configuration files
	configFileYAML := filepath.Join(execDir, "sqlbn.yaml")
	configFileYML := filepath.Join(execDir, "sqlbn.yml")

	// Check if both configuration files exist (yaml and yml)
	var configFile string
	var existErr error
	if _, existErr = os.Stat(configFileYAML); existErr == nil {
		if _, existErr = os.Stat(configFileYML); existErr == nil {
			fmt.Println("Both 'sqlbn.yaml' and 'sqlbn.yml' found in current directory. Please remove one of them.")
			os.Exit(1)
		}
		configFile = configFileYAML
	} else if _, existErr = os.Stat(configFileYML); existErr == nil {
		configFile = configFileYML
	} else {
		fmt.Println("Config file 'sqlbn.yaml' or 'sqlbn.yml' not found in current directory.")
		os.Exit(1)
	}

	// Read the configuration file
	config := readConfig(configFile)

	// Generate the SQL file
	if err := generateSQL(config); err != nil {
		fmt.Println("Error generating SQL file:", err)
		os.Exit(1)
	}

	fmt.Println("SQL file generated successfully!")
}

// Function to read the YAML configuration file
func readConfig(configFile string) Config {
	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("Error reading the configuration file:", err)
		os.Exit(1)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing the configuration file:", err)
		os.Exit(1)
	}

	return config
}

// Function to generate the SQL file
func generateSQL(config Config) error {
	out, err := os.Create(config.OutputDir)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Println("Failed to close create file:", err)
		}
	}(out)

	files, err := os.ReadDir(config.QueryDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".sql" {
			filePath := filepath.Join(config.QueryDir, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			if _, err := out.Write(content); err != nil {
				fmt.Println("Error writing to output file:", err)
				continue
			}

			if _, err := out.WriteString("\n\n"); err != nil {
				fmt.Println("Error writing to output file:", err)
				continue
			}
		}
	}

	return nil
}
