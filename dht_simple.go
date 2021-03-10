package grid

/**
 * Simple DHT to show how things can work.
 *
 */

import "fmt"

// Very simple interface to simulate a dht
type dht interface {
	// Put puts a new (hash,value) on to the DHT
	Put(hash string, id string)
	// Get gets all values for a given hash on the DHT
	Get(hash string) []string
}

type dhtsimple struct {
	data      map[string][]string
	num_items int
	num_keys  int
}

func NewDHTSimple() *dhtsimple {
	d := &dhtsimple{}
	d.data = make(map[string][]string)
	return d
}

// Put puts a value into the DHT
func (d *dhtsimple) Put(hash string, id string) {
	dd, ok := d.data[hash]
	if !ok {
		dd = make([]string, 0)
		d.num_keys++
	}
	d.num_items++
	d.data[hash] = append(dd, id)
}

// Get gets a value from the DHT
func (d *dhtsimple) Get(hash string) []string {
	dd, ok := d.data[hash]
	if ok {
		return dd
	}
	return nil
}

func (d *dhtsimple) String() string {
	return fmt.Sprintf("DHT keys=%d items=%d", d.num_keys, d.num_items)
}
