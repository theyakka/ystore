package ystore

import (
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Store is a giant data map that is constructed from one or more data files
// that are stored within a directory and series of sub-directories
type Store struct {
	// data is the primary storage for all the parsed data/config files
	data map[string]interface{}

	// PrefixDirectories adds a prefix to the data map for any directories that are
	// not the top-level directory. For example: given the file
	// <datadir>/categories/somecat.toml, the contents of the toml file will live in
	// the map under the prefix (key) "categories"
	PrefixDirectories bool

	// Exclude contains patterns that we should be excluding when walking the data
	// directory
	Exclude []string
}

//
func NewStore() *Store {
	return &Store{
		data:              map[string]interface{}{},
		PrefixDirectories: true,
	}
}

func (ds *Store) ReadFile(filePath string) error {
	// clear the data map
	ds.data = map[string]interface{}{}
	// check to see if the directory exists
	if _, statErr := os.Stat(filePath); statErr != nil {
		return statErr
	}
	// read the data / config files within the directory
	dataMap, dataReadErr := ds.readFile(filePath)
	if dataReadErr != nil {
		return dataReadErr
	}
	ds.data = dataMap
	return nil
}

func (ds *Store) ReadFiles(filePaths ...string) error {
	// clear the data map
	ds.data = map[string]interface{}{}
	// check to see if the directory exists
	fullData := map[string]interface{}{}
	for _, filePath := range filePaths {
		if _, statErr := os.Stat(filePath); statErr != nil {
			return statErr
		}
		// read the data / config files within the directory
		fileData, dataReadErr := ds.readFile(filePath)
		if dataReadErr != nil {
			return dataReadErr
		}
		MergeMaps(fileData, fullData, nil)
	}
	ds.data = fullData
	return nil
}

// ReadDir will parse all data files within the directory and all sub-directories
func (ds *Store) ReadDir(path string) error {
	// clear the data map
	ds.data = map[string]interface{}{}
	// check to see if the directory exists
	statInfo, statErr := os.Stat(path)
	if statErr != nil {
		return statErr
	}
	// check to see if the defined directory is actually a directory
	if !statInfo.IsDir() {
		return errors.New("you must specify a directory")
	}
	// read the data / config files within the directory
	dataMap, dataReadErr := ds.readAllFiles(path)
	if dataReadErr != nil {
		return dataReadErr
	}
	ds.data = dataMap
	return nil
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
		if fileDir != path && ds.PrefixDirectories {
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

func (ds *Store) readFile(filePath string) (map[string]interface{}, error) {
	data, dataErr := ioutil.ReadFile(filePath)
	if dataErr != nil {
		return nil, dataErr
	}
	var fileMap map[string]interface{}
	switch filepath.Ext(filePath) {
	case ".toml":
		tomlErr := toml.Unmarshal(data, &fileMap)
		if tomlErr != nil {
			break
		}
		return fileMap, nil
	case ".yaml":
		yamlErr := yaml.Unmarshal(data, &fileMap)
		if yamlErr != nil {
			break
		}
		return fileMap, nil
	case ".json":
		jsonErr := json.Unmarshal(data, &fileMap)
		if jsonErr != nil {
			break
		}
		return fileMap, nil
	}
	return nil, nil
}

func (ds Store) AllValues() map[string]interface{} {
	return ds.data
}

func (ds Store) Get(key string) interface{} {
	splitKey := strings.Split(key, DataKeySeparator)
	val := SearchMap(ds.data, splitKey)
	if val == nil {
		return nil
	}
	return val
}

func (ds Store) GetD(key string, defaultValue interface{}) interface{} {
	splitKey := strings.Split(key, DataKeySeparator)
	val := SearchMap(ds.data, splitKey)
	if val == nil {
		return defaultValue
	}
	return val
}


func (ds Store) GetString(key string) string {
	return cast.ToString(ds.Get(key))
}

func (ds Store) GetStringD(key string, defaultValue string) string {
	return cast.ToString(ds.GetD(key, defaultValue))
}

func (ds Store) GetBool(key string) bool {
	return cast.ToBool(ds.Get(key))
}

func (ds Store) GetBoolD(key string, defaultValue bool) bool {
	return cast.ToBool(ds.GetD(key, defaultValue))
}

func (ds Store) GetInt(key string) int {
	return cast.ToInt(ds.Get(key))
}

func (ds Store) GetIntD(key string, defaultValue int) int {
	return cast.ToInt(ds.GetD(key, defaultValue))
}


func (ds Store) GetFloat(key string) float64 {
	return cast.ToFloat64(ds.Get(key))
}

func (ds Store) GetFloatD(key string, defaultValue float64) float64 {
	return cast.ToFloat64(ds.GetD(key, defaultValue))
}

func (ds Store) GetStringSlice(key string) []string {
	return cast.ToStringSlice(ds.Get(key))
}

func (ds Store) GetStringSliceD(key string, defaultValue []string) []string {
	return cast.ToStringSlice(ds.GetD(key, defaultValue))
}


func (ds Store) GetIntSlice(key string) []int {
	return cast.ToIntSlice(ds.Get(key))
}

func (ds Store) GetIntSliceD(key string, defaultValue []int) []int {
	return cast.ToIntSlice(ds.GetD(key, defaultValue))
}
