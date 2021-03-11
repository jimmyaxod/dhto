package index_kvp

/**
 * Represents something that has a list of kvps
 *
 */

type Dataitem interface {
	GetKVPs() []DataKVP
}

type RealDataitem struct {
	kvps []DataKVP
}

// NewDataitem creates a new dataitem
func NewDataitem() *RealDataitem {
	return &RealDataitem{
		kvps: make([]DataKVP, 0),
	}
}

// GetKVPs gets current kvps
func (di *RealDataitem) GetKVPs() []DataKVP {
	return di.kvps
}

// AddKVP add KVP to this item
func (di *RealDataitem) AddKVP(kvp DataKVP) {
	di.kvps = append(di.kvps, kvp)
}

// String shows as string
func (di *RealDataitem) String() string {
	s := "Dataitem {"
	for _, kvp := range di.kvps {
		s = s + kvp.String() + " "
	}
	s = s + "}"
	return s
}
