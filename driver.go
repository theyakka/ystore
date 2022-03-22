package ystore

type DriverParameters struct {
	// IsReadOnly is true if the drivers lacks support for persistence.
	IsReadOnly bool
	// Name returns a name describing the drivers.
	Name string
	// AutoPersist means the driver will attempt to persist whenever a change is made
	// to the underlying data if this value is set to true.
	AutoPersist bool
}

type Driver interface {
	// Parameters returns details about the drivers and its capabilities.
	Parameters() *DriverParameters
	// Load loads the data from the provided URI.
	Load(store *Store, uris ...string) error
	// Persist saves the data if the drivers supports persistence.
	Persist() error
}
