package json

import (
	"encoding/json"
	"github.com/theyakka/ystore"
	"io/ioutil"
)

type Driver struct {
	uris []string
}

func NewDriver() ystore.Driver {
	return &Driver{}
}

func (jd *Driver) Load(store *ystore.Store, uris ...string) error {
	jd.uris = uris
	for _, uri := range uris {
		fileData, fileErr := ioutil.ReadFile(uri)
		if fileErr != nil {
			return fileErr
		}
		var fileMap map[string]interface{}
		jsonErr := json.Unmarshal(fileData, &fileMap)
		if jsonErr != nil {
			return jsonErr
		}
		// reset the store entries
		store.Clear()
		// add the map values
		ystore.AddMapValues(store, store.Entries(), fileMap)
	}
	return nil
}

func (jd *Driver) Parameters() *ystore.DriverParameters {
	return &ystore.DriverParameters{
		IsReadOnly:  false,
		Name:        "json",
		AutoPersist: false,
	}
}

func (jd *Driver) Persist() error {
	return nil
}
