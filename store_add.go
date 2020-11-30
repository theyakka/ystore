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
	} else if source.Kind() == reflect.Map {
		ds.mergeMap(destination, source, options)
	}
}

func (ds *Store) mergeMap(destination map[string]interface{}, source reflect.Value, options *MergeOptions) {
	mapRange := source.MapRange()
	for mapRange.Next() {
		key := mapRange.Key().String()
		val := reflect.ValueOf(mapRange.Value().Interface())
		if val.Type().Kind() == reflect.Ptr {
			val = val.Elem()
		}
		if val.Type().Kind() == reflect.Struct {
			structMap := map[string]interface{}{}
			ds.mergeStruct(structMap, val, options)
			destination[key] = structMap
		} else if val.Type().Kind() == reflect.Map {
			outMap := map[string]interface{}{}
			ds.mergeMap(outMap, val, options)
			destination[key] = outMap
		} else if val.Type().Kind() == reflect.Slice {
			var slice []interface{}
			ds.mergeSlice(&slice, val, options)
			destination[key] = slice
		} else {
			// at this point, we have a raw data element or it's a data element that
			// we aren't going to be processing any further so we just return it and
			// allow it to be assigned another way
			destination[key] = val.Interface()
		}

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
		indexedVal := source.Field(i)
		if fi.Type.Kind() == reflect.Ptr {
			indexedVal = indexedVal.Elem()
		}
		kind := fi.Type.Kind()
		if fi.Type.Kind() == reflect.Interface {
			indexedVal = reflect.ValueOf(indexedVal.Interface())
			kind = indexedVal.Type().Kind()
		}
		if kind == reflect.Struct {
			structMap := map[string]interface{}{}
			ds.mergeStruct(structMap, indexedVal, options)
			destination[key] = structMap
		} else if kind == reflect.Map {
			outMap := map[string]interface{}{}
			ds.mergeMap(outMap, indexedVal, options)
			destination[key] = outMap
		} else if kind == reflect.Slice {
			var slice []interface{}
			ds.mergeSlice(&slice, indexedVal, options)
			destination[key] = slice
		} else {
			// at this point, we have a raw data element or it's a data element that
			// we aren't going to be processing any further so we just return it and
			// allow it to be assigned another way
			destination[key] = indexedVal.Interface()
		}
	}
}

func (ds *Store) mergeSlice(destination *[]interface{}, source reflect.Value, options *MergeOptions) {
	for i := 0; i < source.Len(); i++ {
		iv := source.Index(i)
		if iv.Kind() == reflect.Ptr {
			iv = iv.Elem()
		}
		if iv.Kind() == reflect.Struct {
			structMap := map[string]interface{}{}
			ds.mergeStruct(structMap, iv, options)
			*destination = append(*destination, structMap)
		} else if iv.Kind() == reflect.Map {
			outMap := map[string]interface{}{}
			ds.mergeMap(outMap, iv, options)
			*destination = append(*destination, outMap)
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
