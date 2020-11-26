package ystore

type StoreOption func(options *StoreOptions)

type StoreOptions struct {
	// prefixDirectories adds a prefix to the data map for any directories that are
	// not the top-level directory. For example: given the file
	// <datadir>/categories/somecat.toml, the contents of the toml file will live in
	// the map under the prefix (key) "categories".
	prefixDirectories bool

	// stopOnFileErr, if true, will cause the add functions to fail whenever there is an
	// error processing a file in a multi-file or directory based operation. If false (the
	// default), the file will be skipped and the operation will continue.
	stopOnFileErr bool

	// exclusions contains patterns that we should be excluding when walking the data
	// directory
	exclusions []string
}

func NewStoreOptions(options ...StoreOption) StoreOptions {
	// create a new options struct with the default values preset
	o := StoreOptions{
		prefixDirectories: true,
		stopOnFileErr:     false,
	}
	for _, of := range options {
		of(&o)
	}
	return o
}

func WithStoreOptions(options StoreOptions) StoreOption {
	return func(storeOptions *StoreOptions) {
		copyOptions(&options, storeOptions)
	}
}

func WithPrefixDirs(prefixDirs bool) StoreOption {
	return func(options *StoreOptions) {
		options.prefixDirectories = prefixDirs
	}
}

func WithExclusions(exclusions ...string) StoreOption {
	return func(options *StoreOptions) {
		options.exclusions = exclusions
	}
}

type MergeOption func(options *MergeOptions)

type MergeOptions struct {
	// overwrite will overwrite any existing value
	overwrite bool

	// appendSlices will append the values for any slices vs overwriting them
	appendSlices bool
}

func NewMergeOptions(options ...MergeOption) MergeOptions {
	mergeOptions := MergeOptions{
		overwrite:    true,
		appendSlices: true,
	}
	for _, of := range options {
		of(&mergeOptions)
	}
	return mergeOptions
}

func WithMergeOptions(options MergeOptions) MergeOption {
	return func(mergeOptions *MergeOptions) {
		copyMergeOptions(&options, mergeOptions)
	}
}

func WithOverwriteFlag(flag bool) MergeOption {
	return func(options *MergeOptions) {
		options.overwrite = flag
	}
}

func WithAppendSlicesFlag(flag bool) MergeOption {
	return func(options *MergeOptions) {
		options.appendSlices = flag
	}
}

func copyOptions(from, to *StoreOptions) {
	to.prefixDirectories = from.prefixDirectories
	to.exclusions = from.exclusions
}

func copyMergeOptions(from, to *MergeOptions) {
	to.overwrite = from.overwrite
	to.appendSlices = from.appendSlices
}
