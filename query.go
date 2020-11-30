package ystore

import (
	"reflect"
	"strings"
)

type Query struct {
	store        *Store
	criteria     []criteria
	defaultValue interface{}
}

func NewQuery(store *Store) *Query {
	return &Query{
		store: store,
	}
}

func (ds *Store) NewQuery() *Query {
	return &Query{
		store: ds,
	}
}

func (q *Query) Get(path string) *Query {
	// pre-process the path into it's splits now so that if we need to run the query
	// multiple times we don't have to do it over and over again
	splitPath := strings.Split(path, ".")
	for _, s := range splitPath {
		q.criteria = append(q.criteria, criteria{
			segment: s,
		})
	}
	return q
}

func (q *Query) WithDefault(value interface{}) *Query {
	q.defaultValue = value
	return q
}

func (q *Query) Run(dest interface{}, options *QueryOptions) error {
	return q.RunD(dest, nil, options)
}

func (q *Query) RunD(dest interface{}, defaultValue interface{}, options *QueryOptions) error {
	if len(q.criteria) == 0 {
		return NewQueryError("you must specify one or more query conditions")
	}
	// run the query
	var results []interface{}
	q.findPath(q.store.data, q.criteria, &results, options)
	// set the value
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Ptr || !rv.IsValid() {
		return NewQueryError("destination must be a valid pointer to something")
	}

	elemVal := results[0]
	if elemVal == nil && defaultValue != nil {
		elemVal = defaultValue
	}
	elem := reflect.Indirect(rv.Elem())
	elem.Set(reflect.ValueOf(elemVal))
	return nil
}

func (q *Query) findPath(source interface{}, criteria []criteria, results *[]interface{}, options *QueryOptions) {
	switch source.(type) {
	case map[string]interface{}:
		mapVal := source.(map[string]interface{})
		key := criteria[0].segment
		nextSource := mapVal[key]
		if len(criteria) == 1 {
			*results = append(*results, nextSource)
		} else {
			if nextSource != nil {
				q.findPath(nextSource, criteria, results, options)
			}
		}
	case []interface{}:
		sliceVal := source.([]interface{})
		for _, v := range sliceVal {
			q.findPath(v, criteria[1:], results, options)
		}
	default:
		if len(criteria) == 1 {
			*results = append(*results, source)
		}
	}
}

type criteria struct {
	segment string
}

type QueryError struct {
	message string
}

func NewQueryError(message string) *QueryError {
	return &QueryError{
		message: message,
	}
}

func (q QueryError) Error() string {
	return q.message
}

type QueryOption func(options *QueryOptions)

type QueryOptions struct {
	continueOnFind bool
}

func NewQueryOptions(options ...QueryOption) QueryOptions {
	o := QueryOptions{
		continueOnFind: true,
	}
	for _, of := range options {
		of(&o)
	}
	return o
}

func WithContinueOnFind(shouldContinue bool) QueryOption {
	return func(options *QueryOptions) {
		options.continueOnFind = shouldContinue
	}
}
