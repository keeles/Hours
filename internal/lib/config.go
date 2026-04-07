package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	completion "github.com/keeles/hours/internal/completions"
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

func GetClientNameForCurrentDirectory() (string, bool) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		fmt.Println("Error reading current directory")
		return "", false
	}

	config, err := ReadFile()
	if err != nil {
		return "", false
	}

	absPath, err := filepath.Abs(currentDirectory)
	if err != nil {
		return "", false
	}

	if _, ok := config.Directories[absPath]; ok {
		clientName := config.Directories[absPath]
		return clientName, true
	}

	fmt.Printf("No client associated with %s \n", currentDirectory)

	return "", false
}

func PrintCompletionFile(shell string) error {
	var filepath string

	switch shell {
	case "bash":
		filepath = "shells/hours.sh"
	case "zsh":
		filepath = "shells/_hours"
	case "fish":
		filepath = "shells/hours.fish"
	default:
		filepath = ""
	}

	file, err := completion.Files.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		return err
	}

	return nil
}
