ystore is an application data management tool. It can be used to parse configuration files or used as a way to import arbitrary data into your application from data files (e.g.: from a CMS or other system).

## Features

- support for JSON, TOML and YAML formats
- parse directories of files into single stores
- support nested keys
- simple key syntax (e.g.: `node.sub.key`)

## FAQ

### How is this different to Viper?

First up, Viper is **awesome** and super useful. ystore isn't trying to compete with viper. The overall intent of ystore is to offer a more *generic* mechanism for data / configuration file parsing, to allow formats to be used interchangeably, and to make it simple to just parse a bunch of it at the same time. 

ystore is less focused on application configuration and more of reading in arbitrary application data / building a "database" for you app data. 

## License
ystore is released under a modified MIT license. See LICENSE for details.

Portions of the code are derived from other works. Please see NOTICE for details on usage and their associated licenses.