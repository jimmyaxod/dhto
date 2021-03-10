package index_geospatial

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	CacheDIR = "../gridcache"
)

/**
 * For a given x / y we cache the grid makeup
 *
 */

// GetGridData reads from the cache...
func GetGridData(y int, x int) ([]int, int, error) {
	file, err := os.Open(fmt.Sprintf("%s/grid_%d_%d.data", CacheDIR, y, x))
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data []int
	total := 0
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err == nil {
			total += v
			data = append(data, v)
		}
	}
	// Check it's the correct length
	if len(data) != y {
		return nil, 0, errors.New("Cache had wrong number of entries")
	}
	return data, total, nil
}

// PutGridData writes to the cache...
func PutGridData(y int, x int, data []int) error {
	file, err := os.Create(fmt.Sprintf("%s/grid_%d_%d.data", CacheDIR, y, x))
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, v := range data {
		fmt.Fprintf(w, "%d\n", v)
	}
	return w.Flush()
}
