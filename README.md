<br/>
<img src="https://storage.googleapis.com/product-logos/logo_ystore.svg" width="70" height="70" alt="ystore">
<br/>
ystore is a data management tool. It can be used as a generic key-value datastore, for flexible configuration storage or used as a way to import arbitrary data into your application from data files (e.g.: from a CMS or other system).

The goal of ystore's is to provide flexible access to your data.

## Features
- flexible key-value storage
- load directly from JSON, TOML and YAML files
- merge whole directories of files into single stores
- simple nested key queries (e.g.: `GetString("node.sub.key")`)
- match store keys using a pattern
- `Store` splitting and merging
- conversion from `map` values

## Memory-based Stores

At its core, ystore is a simple **key-value** store. You can easily
create a simple `Store` like so:

```go
store := NewStore()
store.Set("color", "red")
store.Set("length", 100)
fmt.Printf("The item is %s and the length is %d.\n", 
	store.GetString("color"), store.GetInt("length"))
// Output: The item is red and the length is 100.
```

You can also implement your store with default values using a the `map` initializer:

```go
store := NewStoreFromMap(map[string]interface{}{
	"color" : green,
	"length" : 80,
})
fmt.Printf("The item is %s and the length is %d.\n", 
	store.GetString("color"), store.GetInt("length"))
// Output: The item is green and the length is 80.
```

## File-based Stores

Let's assume we have a `.toml` file that looks like this:

```toml
[item]
name="First Item"
colors=["red", "green", "blue"]
numbers=[5, 2, 1, 6, 5]
```

We can load the file by creating a `Store` and then using the `ReadFile` function. Note: You will need to pass the absolute file path to `ReadFile`.

```go
filename, fileErr := filepath.Abs("./config.toml")
if fileErr != nil {
	// handle the path error
	return
}
store := NewStore()
if readErr := store.ReadFile(filename); readErr != nil {
	// handle store load error
	return
}
name := store.GetString("item.name")
numbers := store.GetIntSlice("item.numbers")
fmt.Printf("The item name is '%s' and the numbers slice has %d element(s).\n", name, len(numbers))
// Output: The item name is 'First Item' and the numbers slice has 5 element(s).
```

## FAQ

### Can I use this as a database?

You could, but it may not be the best solution. Our goal is to provide a super flexible, in-memory data solution for a variety of use cases. Our goal isn't to be a full scale database. A quick search on Google will find a variety of great solutions if you need to embed a database.

### What kind of use cases are people using ystore for?

We've been using it for things like environment parsing, configuration file loading, parsing network json responses, and a lot of other things. ystore is really useful whenever you're trying to bring together a variety of different data sources for easy retrieval / manipulation. 

### Is it production ready?

It's currently in use on a variety of production servers right now. So, yes it is. If you're using it, drop us a line to tell us!


## License
ystore is released under a modified MIT license. See LICENSE for details.

Portions of the code are derived from other works. Please see NOTICE for details on usage, and their associated licenses.