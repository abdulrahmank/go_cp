package go_cp

import (
	"golang.org/x/exp/mmap"
)

func Load(fileName string) ([]byte, error) {
	if readerAt, e := mmap.Open(fileName); e != nil {
		return nil, e
	} else {
		byteArr := make([]byte, readerAt.Len())
		readerAt.ReadAt(byteArr, 0)
		return byteArr, nil
	}
	return nil, nil
}
