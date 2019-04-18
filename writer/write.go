package writer

import (
	"log"
	"os"
	"syscall"
)

type CpWriter interface {
	Write(content []byte, fileName string) error
}

type CpMMapWriterImpl struct{}

func (c *CpMMapWriterImpl) Write(content []byte, fileName string) error {
	if file, e := os.Create(fileName); e != nil {
		return e
	} else {
		file.Seek(int64(len(content)-1), 0)
		file.Write([]byte(" "))
		if mmap, err := syscall.Mmap(int(file.Fd()), 0, 100, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED); err != nil {
			return err
		} else {
			copy(mmap, content)
			if err := syscall.Munmap(mmap); err != nil {
				return err
			}
		}
	}
	return nil
}

type CpIOUtilWriterImpl struct{}

func (c *CpIOUtilWriterImpl) Write(content []byte, fileName string) error {
	if file, e := os.Create(fileName); e != nil {
		return e
	} else {
		if _, e := file.Write(content); e != nil {
			log.Fatal(e)
		}
		file.Close()
	}
	return nil
}
