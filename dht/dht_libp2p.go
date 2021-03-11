package dht

import (
	"context"
	"fmt"

	libp2pdht "github.com/libp2p/go-libp2p-kad-dht"
)

/**
 * DHT using libp2p kad
 *
 */

type DHTLibp2p struct {
	dht libp2pdht.IpfsDHT

	metric_total_puts int
	metric_total_gets int
}

func NewDHTLibp2p() DHTLibp2p {
	return DHTLibp2p{}
}

// Put puts a key value on the DHT
func (d *DHTLibp2p) Put(key string, val string) {
	err := d.dht.PutValue(context.TODO(), key, []byte(val))
	if err != nil {
		fmt.Printf("PutValue error %v\n", err)
	}
	d.metric_total_puts++
}

// Get gets a key from the dht
func (d *DHTLibp2p) Get(key string) []string {
	MAX_VALS := 1000 // For now. TODO: Stream results
	vals, err := d.dht.GetValues(context.TODO(), key, MAX_VALS)
	if err != nil {
		fmt.Printf("GetValues error %v\n", err)
	}
	d.metric_total_gets++

	results := make([]string, 0)
	for _, val := range vals {
		// Contains the value, and where it came from (peer). For now just use val
		results = append(results, string(val.Val))
	}

	return results
}
