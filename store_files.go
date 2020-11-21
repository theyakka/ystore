package ystore

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (ds *Store) readFile(filePath string) (map[string]interface{}, error) {
	// check to see if the directory exists
	if _, statErr := os.Stat(filePath); statErr != nil {
		return nil, statErr
	}
	data, dataErr := ioutil.ReadFile(filePath)
	if dataErr != nil {
		return nil, dataErr
	}
	var fileMap map[string]interface{}
	switch filepath.Ext(filePath) {
	case ".toml":
		tomlErr := toml.Unmarshal(data, &fileMap)
		if tomlErr != nil {
			return nil, tomlErr
		}
		return fileMap, nil
	case ".yaml":
		fallthrough
	case ".yml":
		yamlErr := yaml.Unmarshal(data, &fileMap)
		if yamlErr != nil {
			return nil, yamlErr
		}
		return fileMap, nil
	case ".json":
		jsonErr := json.Unmarshal(data, &fileMap)
		if jsonErr != nil {
			return nil, jsonErr
		}
		return fileMap, nil
	}
	return nil, errors.New(fmt.Sprintf("file type (%s) is unsupported", filepath.Ext(filePath)))
}

func (ds *Store) readAllFiles(dirPath string) (map[string]interface{}, error) {
	var fullData = map[string]interface{}{}
	walkErr := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// skip directories. the walk function will flatten any sub-directories
			return nil
		}
		fileData, fileErr := ds.readFile(path)
		if fileErr != nil {
			return fileErr
		}
		fileDir := filepath.Dir(path)
		if fileDir != path && ds.options.prefixDirectories {
			// add the prefix to the file data map because it is in a sub-directory
			mapPrefix := BaseDir(path)
			fileData = map[string]interface{}{
				mapPrefix: fileData,
			}
		}
		// merge the main data map and the file data map
		MergeMaps(fileData, fullData, nil)
		return nil
	})
	if walkErr != nil {
		return nil, walkErr
	}
	return fullData, nil
}
