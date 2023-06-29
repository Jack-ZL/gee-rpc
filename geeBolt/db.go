package geeBolt

import (
	"os"
	"syscall"
)

type DB struct {
	data []byte
	file *os.File
}

const maxMapSize = 1 << 31

func (db *DB) mmap(sz int) error {
	b, err := syscall.Mmap()
	return err
}

func Open(path string) {

}
