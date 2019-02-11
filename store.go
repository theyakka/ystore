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
	// dataDir is the location of all the various data/config files
	dataDir string

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
func StoreWithDir(dirPath string) *Store {
	return &Store{
		dataDir:           dirPath,
		data:              map[string]interface{}{},
		PrefixDirectories: true,
	}
}

// ReadAll will parse all data files within the directory and all sub-directories
func (ds *Store) ReadAll() error {
	// clear the data map
	ds.data = map[string]interface{}{}
	// check to see if the directory exists
	statInfo, statErr := os.Stat(ds.dataDir)
	if statErr != nil {
		return statErr
	}
	// check to see if the defined directory is actually a directory
	if !statInfo.IsDir() {
		return errors.New("you must specify a directory")
	}
	// read the data / config files within the directory
	dataMap, dataReadErr := ds.readData(ds.dataDir)
	if dataReadErr != nil {
		return dataReadErr
	}
	ds.data = dataMap
	return nil
}

func (ds *Store) readData(dirPath string) (map[string]interface{}, error) {
	var fullData = map[string]interface{}{}
	walkErr := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// skip directories. the walk function will flatten any sub-directories
			return nil
		}
		fileData, fileErr := ds.dataWalkFunc(dirPath, path, info, err)
		if fileErr != nil {
			return fileErr
		}
		fileDir := filepath.Dir(path)
		if fileDir != ds.dataDir && ds.PrefixDirectories {
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

func (ds *Store) dataWalkFunc(dirPath string, filePath string, info os.FileInfo, err error) (map[string]interface{}, error) {
	data, dataErr := ioutil.ReadFile(filePath)
	if dataErr != nil {
		return nil, dataErr
	}
	var fileMap map[string]interface{}
	switch filepath.Ext(info.Name()) {
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

func (ds Store) GetString(key string) string {
	return cast.ToString(ds.Get(key))
}

func (ds Store) GetBool(key string) bool {
	return cast.ToBool(ds.Get(key))
}

func (ds Store) GetInt(key string) int {
	return cast.ToInt(ds.Get(key))
}

func (ds Store) GetFloat(key string) float64 {
	return cast.ToFloat64(ds.Get(key))
}

func (ds Store) GetStringSlice(key string) []string {
	return cast.ToStringSlice(ds.Get(key))
}

func (ds Store) GetIntSlice(key string) []int {
	return cast.ToIntSlice(ds.Get(key))
}
