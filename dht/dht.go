package dht

// Very simple interface to simulate a dht
type DHT interface {
	// Put puts a new (hash,value) on to the DHT
	Put(hash string, id string)
	// Get gets all values for a given hash on the DHT
	Get(hash string) []string
}
