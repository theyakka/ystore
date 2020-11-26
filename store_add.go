package ystore

import (
	"errors"
	"os"
	"reflect"
	"strings"
)

func (ds *Store) AddData(data interface{}, options *MergeOptions) error {
	v := reflect.ValueOf(data)
	// if the data is a pointer, get the actual element and work from that
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Map && v.Kind() != reflect.Struct {
		return errors.New("this operation only supports maps and struct")
	}
	ds.merge(ds.data, v, options)
	return nil
}

func (ds *Store) merge(destination map[string]interface{}, source reflect.Value, options *MergeOptions) {
	if source.Kind() == reflect.Struct {
		ds.mergeStruct(destination, source, options)
	} else {
		// at this point, we have a raw data element or it's a data element that
		// we aren't going to be processing any further so we just return it and
		// allow it to be assigned another way
	}
}

func (ds *Store) mergeStruct(destination map[string]interface{}, source reflect.Value, options *MergeOptions) {
	vType := source.Type()
	fieldCount := source.NumField()
	if fieldCount == 0 {
		return
	}
	for i := 0; i < fieldCount; i++ {
		fi := vType.Field(i)
		key := fi.Name
		if fi.Anonymous || !isExportedField(key) {
			continue
		}
		if tag := fi.Tag.Get(TagPrefix); tag != "" {
			splitTag := strings.Split(tag, ",")
			// TODO - tag options
			key = splitTag[0]
		}
		if fi.Type.Kind() == reflect.Struct {
			structMap := map[string]interface{}{}
			ds.mergeStruct(structMap, source.Field(i), options)
			destination[key] = structMap
		} else if fi.Type.Kind() == reflect.Slice {
			var slice []interface{}
			ds.mergeSlice(&slice, source.Field(i), options)
			destination[key] = slice
		} else {
			destination[key] = source.Field(i).Interface()
		}
	}
}

func (ds *Store) mergeSlice(destination *[]interface{}, source reflect.Value, options *MergeOptions) {
	for i := 0; i < source.Len(); i++ {
		iv := source.Index(i)
		if iv.Kind() == reflect.Struct {
			structMap := map[string]interface{}{}
			ds.mergeStruct(structMap, iv, options)
			*destination = append(*destination, structMap)
		} else if iv.Kind() == reflect.Slice {
			var slice []interface{}
			ds.mergeSlice(&slice, iv, options)
			*destination = append(*destination, slice)
		} else {
			*destination = append(*destination, iv.Interface())
		}
	}
}

func isExportedField(name string) bool {
	first := name[0]
	if 'a' <= first && first <= 'z' || first == '_' {
		return false
	}
	return true
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
	ds.AddData(finalMap, nil)
}
