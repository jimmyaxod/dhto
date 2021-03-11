package dht

/**
 * Simple DHT to show how things can work.
 *
 */

import "fmt"

type DHTsimple struct {
	data      map[string][]string
	num_items int
	num_keys  int

	metric_total_gets       int
	metric_total_gets_count int
	metric_total_puts       int
}

func NewDHTSimple() *DHTsimple {
	d := &DHTsimple{}
	d.data = make(map[string][]string)
	return d
}

// Put puts a value into the DHT
func (d *DHTsimple) Put(hash string, id string) {
	d.metric_total_puts++

	dd, ok := d.data[hash]
	if !ok {
		dd = make([]string, 0)
		d.num_keys++
	}
	d.num_items++
	d.data[hash] = append(dd, id)
}

// Get gets a value from the DHT
func (d *DHTsimple) Get(hash string) []string {
	d.metric_total_gets++

	dd, ok := d.data[hash]
	if ok {
		d.metric_total_gets_count += len(dd)
		return dd
	}
	return nil
}

func (d *DHTsimple) String() string {
	return fmt.Sprintf("DHT keys=%d items=%d metrics(puts=%d gets=%d (%d))", d.num_keys, d.num_items,
		d.metric_total_puts, d.metric_total_gets, d.metric_total_gets_count)
}
