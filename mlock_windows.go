package bbolt

import (
	"fmt"
)

// mlock locks memory of db file
func mlock(_ *DB, _ int) error {
	return fmt.Errorf("mlock is supported only on UNIX systems")
}

// munlock unlocks memory of db file
func munlock(_ *DB, _ int) error {
	return fmt.Errorf("munlock is supported only on UNIX systems")
}
