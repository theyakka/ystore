package ystore

import (
	"errors"
	"os"
	"reflect"
	"strings"
)

func (ds *Store) AddData(data interface{}) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Map {
		for _, kv := range v.MapKeys() {
			key := kv.String()
			if key != "" {
				ds.Set(key, v.MapIndex(kv).Interface())
			}
		}
	} else if v.Kind() == reflect.Struct {
		vType := v.Type()
		fieldCount := v.NumField()
		if fieldCount == 0 {
			return
		}
		for i := 0; i < fieldCount; i++ {
			fi := vType.Field(i)
			if tag := fi.Tag.Get("ystore"); tag != "" {
				splitTag := strings.Split(tag, ",")
				// TODO - tag options
				ds.Set(splitTag[0], v.Field(i).Interface())
			} else {
				ds.Set(fi.Name, v.Field(i).Interface())
			}
		}
	}
}

func (ds *Store) AddFile(filePath string) error {
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

func (ds *Store) AddFiles(filePaths ...string) error {
	allData := map[string]interface{}{}
	for _, filePath := range filePaths {
		// read the data / config files within the directory
		fileData, dataReadErr := ds.readFile(filePath)
		if dataReadErr != nil && ds.options.stopOnFileErr {
			return dataReadErr
		}
		MergeMaps(fileData, allData, nil)
	}
	MergeMaps(ds.data, allData, nil)
	return nil
}

// AddDir will parse all data files within the directory and all sub-directories
func (ds *Store) AddDir(path string) error {
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

func (ds *Store) AddStores(stores ...*Store) {
	finalMap := map[string]interface{}{}
	for _, store := range stores {
		MergeMaps(store.AllValues(), finalMap, nil)
	}
	ds.AddData(finalMap)
}
