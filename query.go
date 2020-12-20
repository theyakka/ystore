package ystore

import (
	"log"
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
	// no query conditions so we can't proceed
	if len(q.criteria) == 0 {
		return NewQueryError("you must specify one or more query conditions", 0)
	}
	// run the query
	var results []interface{}
	// TODO - should also error
	q.findPath(q.store.data, q.criteria, &results, options)
	// if there were no results we're done
	if len(results) == 0 {
		return NewQueryError("there were no results", 404)
	}
	// set the value
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Ptr || !rv.IsValid() {
		return NewQueryError("destination must be a valid pointer to something", 0)
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
		if nextSource == nil {
			*results = nil
			break
		}
		if len(criteria) == 1 {
			*results = append(*results, nextSource)
		} else {
			q.findPath(nextSource, criteria[1:], results, options)
		}
	case []interface{}:
		sliceVal := source.([]interface{})
		for _, v := range sliceVal {
			q.findPath(v, criteria, results, options)
		}
		log.Println("....")
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
	code    int
}

func NewQueryError(message string, code int) *QueryError {
	return &QueryError{
		message: message,
		code:    code,
	}
}

func (q QueryError) Error() string {
	return q.message
}

func (q QueryError) Code() int {
	return q.code
}

func (q QueryError) IsNoResults() bool {
	return q.code == 404
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
