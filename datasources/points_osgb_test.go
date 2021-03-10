package datasources

import (
	"fmt"
	"testing"
)

func TestOSGB(t *testing.T) {

	data, err := GetOSNames("hydrography")

	if err != nil {
		t.Errorf("Can't load %v\n", err)
	}

	fmt.Printf("Loaded OSNames %v\n", len(data))
}
