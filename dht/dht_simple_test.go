package dht

/**
 * Some tests for dht_simple
 *
 *
 */

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDHTSimple(t *testing.T) {
	dht := NewDHTSimple()

	dht.Put("a", "1")
	dht.Put("a", "2")
	dht.Put("b", "1")

	// Now get them and check...

	// NB here we assume they'll come back in the order they were inserted.
	//    that won't usually be the case.
	a := dht.Get("a")
	assert.Equal(t, len(a), 2)
	assert.Equal(t, a[0], "1")
	assert.Equal(t, a[1], "2")

	b := dht.Get("b")
	assert.Equal(t, len(b), 1)
	assert.Equal(t, b[0], "1")

}
