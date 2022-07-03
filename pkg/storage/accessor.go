package storage

type usersDataAccessor interface {
}

// Accessor is interface which defines all functions used on database/cache/fts
// in the system. Any other packages should use this Accessor instead of using
// eg. database directly.
type Accessor interface {
	usersDataAccessor
}
