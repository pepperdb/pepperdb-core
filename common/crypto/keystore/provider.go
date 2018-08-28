package keystore

// Provider class represents a "provider" for the
// Security API, where a provider implements some or all parts of
// Security. Services that a provider may implement include:
// Algorithms
// Key generation, conversion, and management facilities (such as for
// algorithm-specific keys).
// Each provider has a name and a version number, and is configured
// in each runtime it is installed in.
type Provider interface {

	// Aliases all alias in provider save
	Aliases() []string

	// SetKey assigns the given key (that has already been protected) to the given alias.
	SetKey(a string, key Key, passphrase []byte) error

	// GetKey returns the key associated with the given alias, using the given
	// password to recover it.
	GetKey(a string, passphrase []byte) (Key, error)

	// Delete remove key
	Delete(a string) error

	// ContainsAlias check provider contains key
	ContainsAlias(a string) (bool, error)

	// Clear all entries in provider
	Clear() error
}
