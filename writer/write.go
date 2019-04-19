package writer

import (
	"log"
	"os"
	"syscall"
)

type CpWriter interface {
	Write(content []byte, fileName string) (int, error)
}

type CpMMapWriterImpl struct{}

func (c *CpMMapWriterImpl) Write(content []byte, fileName string) (int, error) {
	if file, e := os.Create(fileName); e != nil {
		return len(content), e
	} else {
		file.Seek(int64(len(content)-1), 0)
		file.Write([]byte(" "))
		if mmap, err := syscall.Mmap(int(file.Fd()), 0, len(content), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED); err != nil {
			return len(content), err
		} else {
			copy(mmap, content)
			if err := syscall.Munmap(mmap); err != nil {
				return len(content), err
			}
		}
	}
	return len(content), nil
}

type CpIOUtilWriterImpl struct{}

func (c *CpIOUtilWriterImpl) Write(content []byte, fileName string) (int, error) {
	n := 0
	if file, e := os.Create(fileName); e != nil {
		return n, e
	} else {
		if n, e = file.Write(content); e != nil {
			log.Fatal(e)
		}
		file.Close()
	}
	return n, nil
}
