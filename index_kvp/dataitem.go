package index_kvp

/**
 * Represents something that has a list of kvps
 *
 */

type Dataitem interface {
	GetKVPs() []DataKVP
}
