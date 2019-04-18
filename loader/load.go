package loader

import (
	"golang.org/x/exp/mmap"
)

type Loader interface {
	Load(fileName string) ([]byte, error)
}

type MMapLoaderImpl struct{}

func (l *MMapLoaderImpl) Load(fileName string) ([]byte, error) {
	if readerAt, e := mmap.Open(fileName); e != nil {
		return nil, e
	} else {
		byteArr := make([]byte, readerAt.Len())
		readerAt.ReadAt(byteArr, 0)
		return byteArr, nil
	}
	return nil, nil
}
