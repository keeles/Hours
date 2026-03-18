package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	Directories map[string]string `json:"directories"`
}

func GetConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".hours.json"
	}
	configDir := filepath.Join(home, ".config", "hours")
	os.MkdirAll(configDir, 0755)

	return filepath.Join(configDir, "config.json")
}

func CheckFileExists() error {
	pathname := GetConfigPath()
	_, err := os.Stat(pathname)

	if os.IsNotExist(err) {
		file, err := os.Create(pathname)
		if err != nil {
			return err
		}
		file.Close()
	}

	return nil
}

func ReadFile() (Config, error) {
	err := CheckFileExists()
	if err != nil {
		return Config{}, err
	}

	pathname := GetConfigPath()
	file, err := os.Open(pathname)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var data Config

	if len(contents) > 0 {
		if err := json.Unmarshal(contents, &data); err != nil {
			return Config{}, err
		}
	}

	if data.Directories == nil {
		data.Directories = make(map[string]string)
	}

	return data, nil
}

func WriteFile(data Config) error {
	results, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	pathname := GetConfigPath()
	return os.WriteFile(pathname, results, 0666)
}

func AppendDirectory(clientName, directory string) error {
	config, err := ReadFile()
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(directory)
	if err != nil {
		return err
	}

	if existing, ok := config.Directories[absPath]; ok {
		return fmt.Errorf("Directory already mapped to client: %s", existing)
	}

	config.Directories[absPath] = clientName

	return WriteFile(config)
}

func RemoveDirectory(directory string) error {
	config, err := ReadFile()
	if err != nil {
		return err
	}

	absPath, err := filepath.Abs(directory)
	if err != nil {
		return err
	}

	delete(config.Directories, absPath)

	return WriteFile(config)
}

func RemoveDirectoryOfClient(clientName string) error {
	config, err := ReadFile()
	if err != nil {
		return err
	}

	var deleted bool
	for directory, client := range config.Directories {
		if client == clientName {
			delete(config.Directories, directory)
			deleted = true
			break
		}
	}

	if !deleted {
		return fmt.Errorf("No directory found associated with client %s \n", clientName)
	}

	return WriteFile(config)
}

