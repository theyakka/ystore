package ystore

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func (ds *Store) AddData(data interface{}) error {
	v := reflect.ValueOf(data)
	// if the data is a pointer, get the actual element and work from that
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Map && v.Kind() != reflect.Struct {
		return errors.New("this operation only supports maps and struct")
	}
	ds.addDataElement(data, "")
	return nil
}

func (ds *Store) addDataElement(data interface{}, keyPath string) {
	v := reflect.ValueOf(data)
	// if the data is a pointer, get the actual element and work from that
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		ds.addStructElement(v, keyPath)
	} else {
		// at this point, we have a raw data element or it's a data element that
		// we aren't going to be processing any further so we just return it and
		// allow it to be assigned another way
		ds.Set(keyPath, data)
	}
}

func (ds *Store) addStructElement(value reflect.Value, keyPath string) {
	vType := value.Type()
	fieldCount := value.NumField()
	if fieldCount == 0 {
		return
	}
	for i := 0; i < fieldCount; i++ {
		fi := vType.Field(i)
		key := fi.Name
		if tag := fi.Tag.Get("ystore"); tag != "" {
			splitTag := strings.Split(tag, ",")
			// TODO - tag options
			key = splitTag[0]
		}
		if keyPath != "" {
			key = fmt.Sprintf("%s.%s", keyPath, key)
		}
		ds.addDataElement(value.Field(i).Interface(), key)
	}
}

//func (ds *Store) addMapElement(value reflect.Value, keyPath string) {
//	for _, kv := range value.MapKeys() {
//		key := kv.String()
//		if key != "" {
//			ds.addDataElement(value.MapIndex(kv).Interface(), key)
//		}
//	}
//}

//func (ds *Store) addSliceElement(value reflect.Value) []interface{} {
//	var sliceToAdd []interface{}
//	for i := 0; i < value.Len(); i++ {
//		sliceToAdd = append(sliceToAdd, ds.addDataElement("", value.Index(i).Interface()))
//	}
//}

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
