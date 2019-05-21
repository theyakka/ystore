<br/>
<img src="https://storage.googleapis.com/product-logos/logo_ystore.png" width="70" height="70">
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

## Creating an empty store

At it's core, ystore is a simple **key-value** store. You can easily
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

## FAQ

### Can I use this as a database?

You could, but it may not be the best solution. Our goal is to provide a super flexible data management solution for a variety of application data and to not be a full scale database. A quick search on Google will find a variety of great solutions if you need to embed a database.

### How is this different to Viper?

First up, Viper is **fantastic** and super useful. ystore isn't trying to compete with viper. 

The overall intent of ystore is to offer a more *generic* data mechanism, to allow files/formats to be used interchangeably, and to make it simple to just parse a bunch of it at the same time. Viper is more suited for application configuration specific task. You can still use ystore for application configuration (it works great!) but you may end up needing to do a few additional tasks yourself.  

## Related

[Viper](https://github.com/spf13/viper) - Go configuration with fangs

## License
ystore is released under a modified MIT license. See LICENSE for details.

Portions of the code are derived from other works. Please see NOTICE for details on usage and their associated licenses.