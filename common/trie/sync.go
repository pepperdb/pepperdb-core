package trie

// SyncTrie data from other servers
// Sync whole trie to build snapshot
func (t *Trie) SyncTrie(rootHash []byte) error {
	return nil
}

// SyncPath from rootHash to key node from other servers
// Useful for verification quickly
func (t *Trie) SyncPath(rootHash []byte, key []byte) error {
	return nil
}
