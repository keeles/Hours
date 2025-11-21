package lib

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Projects map[string]Project 

type Project struct {
	Name string `json:"Name"`
	Hours int `json:"Hours"`
}

func GetDataPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".hours.json"
	}
	configDir := filepath.Join(home, ".config", "hours")
	os.MkdirAll(configDir, 0755) 	

	return filepath.Join(configDir, "data.json")
}	 

func CheckFileExists() (error) {
	pathname := GetDataPath()
	_, err := os.Stat(pathname)

	if os.IsNotExist(err) {
		_, err := os.Create(pathname)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadFile() (Projects, error) {
	err := CheckFileExists()
	if err != nil {
		return nil, err
	}

	pathname := GetDataPath()
	file, err := os.Open(pathname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	data := make(Projects)
	if len(contents) > 0 {
		if err := json.Unmarshal(contents, &data); err != nil {
			return nil, err
		}
	}

	return data, nil
}

func AddNewProject(data Projects, newProjName string) (error) {
	newProj := Project{Name: newProjName, Hours: 0}
	data[strings.ToLower(newProjName)] = newProj

	err := WriteFile(data)
	if err != nil {
		return err
	}

	return nil
}

func WriteFile(data Projects) error {
	results, err := json.Marshal(data)
	if err != nil {
		return err
	}

	pathname := GetDataPath()
	err = os.WriteFile(pathname, results, 0666)
	if err != nil {
		return err
	}

	return nil
}